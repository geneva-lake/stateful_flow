package storage

import (
	"math/rand"

	"github.com/geneva-lake/stateful_flow/logger"
	"github.com/geneva-lake/stateful_flow/service/flow/model"
)

type StorageUnit model.OrderFlow

func (f *StorageUnit) Process(repo *Repository, previous *model.StatusStream, next *model.StatusStream) {
	status := <-previous.Forward
	if status == model.Cancel {
		next.Forward <- model.Cancel
		return
	}
	unit := "storage"
	orderID := rand.Intn(100)
	stored := StoredOrder{
		UserID:       f.UserID,
		ProductID:    f.ProductID,
		OrderID:      orderID,
		ProductPrice: f.ProductPrice,
		Status:       model.Created,
	}
	err := repo.OrderApply(&stored)
	if err != nil {
		f.OrderStatus = model.OrderInternalError
		go logger.LogUnit(logger.Error, f.Config.Name, err,
			f.OrderID, unit, string(model.OrderInternalError))
		next.Forward <- model.Cancel
		previous.Back <- model.Cancel
		return
	}
	f.OrderID = orderID
	next.Forward <- model.Proceed
	<-next.Back
	err = repo.OrderUpdate(f.OrderID, f.OrderStatus)
	if err != nil {
		go logger.LogUnit(logger.Error, f.Config.Name, err,
			f.OrderID, unit, string(f.OrderStatus))
	}
	previous.Back <- model.Proceed
}
