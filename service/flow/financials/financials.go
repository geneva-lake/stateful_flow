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
		f.OrderStatus = model.OrderInternalError
		go logger.LogUnit(logger.Error, f.Config.Name, err,
			f.OrderID, unit, string(model.OrderInternalError))
		previous.Back <- model.Cancel
		return
	}
	if resp.Status == general.StatusError {
		go logger.LogUnit(logger.Info, f.Config.Name, nil,
			f.OrderID, unit, string(model.OrderInternalError))
		f.OrderStatus = model.OrderInternalError
		previous.Back <- model.Cancel
		return
	}
	transactionID := ""
	switch resp.Result.Status {
	case Success:
		transactionID = *resp.Result.TransactionID
		previous.Back <- model.Proceed
	case BalanceNotEnough:
		f.OrderStatus = model.OrderBalanceNotEnough
		go logger.LogUnit(logger.Info, f.Config.Name, nil,
			f.OrderID, unit, string(model.OrderBalanceNotEnough))
		previous.Back <- model.Cancel
	}
	status = <-previous.Forward
	if status == model.Cancel {
		url := fmt.Sprintf(f.Config.FinancialsRollbackURL, transactionID)
		resp, err := general.MakeHTTPRequest[interface{}, Response]("PUT", url, nil)
		if err != nil {
			go logger.LogUnit(logger.Error, f.Config.Name, err,
				f.OrderID, unit, string(f.OrderStatus))
			return
		}
		if resp.Status == general.StatusError {
			go logger.LogUnit(logger.Info, f.Config.Name, nil,
				f.OrderID, unit, string(f.OrderStatus))
		}
	}
}
