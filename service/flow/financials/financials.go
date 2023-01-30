package financials

import (
	"fmt"

	"github.com/geneva-lake/stateful_flow/general"
	"github.com/geneva-lake/stateful_flow/logger"
	"github.com/geneva-lake/stateful_flow/service/flow/model"
)

type FinancialsUnit model.OrderFlow

func (f *FinancialsUnit) Process(previous *model.StatusStream) {
	status := <-previous.Forward
	if status == model.Cancel {
		return
	}
	unit := "financials"
	breq := ApplyRequest{
		ProductID:    f.ProductID,
		OrderID:      f.OrderID,
		ProductPrice: f.ProductPrice,
	}
	url := fmt.Sprintf(f.Config.FinancialsApplyURL, f.UserID.String())
	resp, err := general.MakeHTTPRequest[ApplyRequest, Response]("POST", url, &breq)
	if err != nil {
		go logger.Log(logger.Error, unit, 0, err, nil, nil, nil)
		f.OrderStatus = model.OrderInternalError
		previous.Back <- model.Cancel
		return
	}
	if resp.Status == general.StatusError {
		go logger.Log(logger.Info, unit, 0, err, nil, nil, resp)
		f.OrderStatus = model.OrderInternalError
		previous.Back <- model.Cancel
		return
	}
	transactionID := ""
	switch resp.Result.Status {
	case Success:
		f.OrderStatus = model.OrderFinancialsSuccess
		previous.Back <- model.Proceed
		transactionID = *resp.Result.TransactionID
	case BalanceNotEnough:
		f.OrderStatus = model.OrderBalanceNotEnough
		previous.Back <- model.Cancel
	}
	status = <-previous.Forward
	if status == model.Cancel {
		url := fmt.Sprintf(f.Config.FinancialsRollbackURL, transactionID)
		resp, err := general.MakeHTTPRequest[interface{}, Response]("PUT", url, nil)
		if err != nil {
			go logger.Log(logger.Error, unit, 0, err, nil, nil, nil)
			return
		}
		if resp.Status == general.StatusError {
			go logger.Log(logger.Info, unit, 0, err, nil, nil, resp)
		}
	}
}
