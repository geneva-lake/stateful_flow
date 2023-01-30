package bookkeeping

import (
	"github.com/geneva-lake/stateful_flow/general"
	"github.com/geneva-lake/stateful_flow/logger"
	"github.com/geneva-lake/stateful_flow/service/flow/model"
	"github.com/google/uuid"
)

type BookkepingUnit model.OrderFlow

func (f *BookkepingUnit) Process(previous *model.StatusStream, next *model.StatusStream) {
	status := <-previous.Forward
	if status == model.Cancel {
		next.Forward <- model.Cancel
		return
	}
	unit := "bookkeping"
	additional := make(map[string]interface{})
	breq := ApplyRequest{
		UserID:       f.UserID,
		ProductID:    f.ProductID,
		OrderID:      f.OrderID,
		ProductPrice: f.ProductPrice,
	}
	resp, err := general.MakeHTTPRequest[ApplyRequest, ApplyResponse]("POST", f.Config.BookkepingApplyURL, &breq)
	if err != nil {
		f.Err = err
		f.OrderStatus = model.OrderInternalError
		next.Forward <- model.Cancel
		previous.Back <- model.Cancel
		additional["order_status"] = model.OrderInternalError
		go logger.Log(logger.Error, unit, err, nil, additional)
		return
	}
	if resp.Status == general.StatusError {
		switch resp.Error {
		case ProductNotFound:
			f.OrderStatus = model.OrderInternalError
			next.Forward <- model.Cancel
			previous.Back <- model.Cancel
		case Internalerror:
			f.OrderStatus = model.OrderInternalError
			next.Forward <- model.Cancel
			previous.Back <- model.Cancel
		}
		return
	}
	var recordID *uuid.UUID
	switch resp.Result.Status {
	case Success:
		f.OrderStatus = model.OrderBookkepingSuccess
		recordID = resp.Result.RecordID
		next.Forward <- model.Proceed
	case ProductNotAvailable:
		f.OrderStatus = model.OrderProductNotAvailable
		next.Forward <- model.Cancel
		previous.Back <- model.Cancel
		return
	}
	status = <-next.Back
	if status == model.Cancel {
		previous.Back <- model.Cancel
		return
	}

	updresp, err := general.MakeHTTPRequest[interface{}, UpdateResponse]("GET", f.Config.BookkepingUpdateURL+recordID.String(), nil)
	if err != nil {
		f.Err = err
		f.OrderStatus = model.OrderInternalError
		next.Forward <- model.Cancel
		previous.Back <- model.Cancel
		return
	}
	if updresp.Status == general.StatusError {
		next.Forward <- model.Cancel
		previous.Back <- model.Cancel
	}
}
