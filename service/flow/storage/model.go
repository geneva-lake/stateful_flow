package storage

import (
	"github.com/geneva-lake/stateful_flow/service/flow/model"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StoredOrder struct {
	Status       model.OrderStatus
	OrderID      int
	UserID       uuid.UUID
	ProductID    int
	ProductPrice decimal.Decimal
}
