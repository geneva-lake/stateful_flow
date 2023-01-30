package storage

import "github.com/geneva-lake/stateful_flow/service/flow/model"

type StorageUnit model.OrderFlow

func (f *StorageUnit) Process(repo *Repository, previous *model.StatusStream, next *model.StatusStream) {
	status := <-previous.Forward
	if status == model.Cancel {
		next.Forward <- model.Cancel
		return
	}
	stored := StoredOrder{
		UserID:       f.UserID,
		ProductID:    f.ProductID,
		OrderID:      f.OrderID,
		ProductPrice: f.ProductPrice,
		Status:       model.Created,
	}
	err := repo.OrderApply(&stored)
	if err != nil {
		f.OrderStatus = model.OrderInternalError
		next.Forward <- model.Cancel
		return
	}
	next.Forward <- model.Proceed
	status = <-next.Back
	switch status {
	case model.Cancel:
		repo.OrderUpdate(f.OrderID, f.OrderStatus)
	}
	previous.Back <- model.Proceed
}
