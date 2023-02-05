package bookkeeping

import (
	"fmt"

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
	breq := ApplyRequest{
		UserID:       f.UserID,
		ProductID:    f.ProductID,
		OrderID:      f.OrderID,
		ProductPrice: f.ProductPrice,
	}
	resp, err := general.MakeHTTPRequest[ApplyRequest, ApplyResponse]("POST", f.Config.BookkepingApplyURL, &breq)
	if err != nil {
		f.OrderStatus = model.OrderInternalError
		go logger.LogUnit(logger.Error, f.Config.Name, err,
			f.OrderID, unit, string(model.OrderInternalError))
		next.Forward <- model.Cancel
		previous.Back <- model.Cancel
		return
	}
	if resp.Status == general.StatusError {
		switch resp.Error {
		case ProductNotFound:
			f.OrderStatus = model.OrderInternalError
			go logger.LogUnit(logger.Info, f.Config.Name, nil,
				f.OrderID, unit, string(ProductNotFound))
			next.Forward <- model.Cancel
			previous.Back <- model.Cancel
		case InternalError:
			f.OrderStatus = model.OrderInternalError
			go logger.LogUnit(logger.Info, f.Config.Name, nil,
				f.OrderID, unit, string(InternalError))
			next.Forward <- model.Cancel
			previous.Back <- model.Cancel
		}
		return
	}
	var recordID *uuid.UUID
	switch resp.Result.Status {
	case Success:
		next.Forward <- model.Proceed
	case ProductNotAvailable:
		f.OrderStatus = model.OrderProductNotAvailable
		next.Forward <- model.Cancel
		previous.Back <- model.Cancel
		go logger.LogUnit(logger.Info, f.Config.Name, nil,
			f.OrderID, unit, string(ProductNotAvailable))
		return
	}
	status = <-next.Back
	if status == model.Cancel {
		previous.Back <- model.Cancel
		return
	}
	url := fmt.Sprintf(f.Config.BookkepingUpdateURL, recordID.String())
	updresp, err := general.MakeHTTPRequest[interface{}, UpdateResponse]("PUT", url, nil)
	if err != nil {
		f.OrderStatus = model.OrderInternalError
		go logger.LogUnit(logger.Error, f.Config.Name, err,
			f.OrderID, unit, string(model.OrderInternalError))
		next.Forward <- model.Cancel
		previous.Back <- model.Cancel
		return
	}
	if updresp.Status == general.StatusError {
		go logger.LogUnit(logger.Info, f.Config.Name, nil,
			f.OrderID, unit, string(model.OrderInternalError))
		next.Forward <- model.Cancel
		previous.Back <- model.Cancel
	}
}
