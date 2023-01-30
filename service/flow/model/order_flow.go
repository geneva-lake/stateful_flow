package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderStatus string

const (
	NotPossible              OrderStatus = "not_possible"
	Created                  OrderStatus = "created"
	OrderInternalError       OrderStatus = "internal_error"
	OrderBookkepingSuccess   OrderStatus = "bookkeping_success"
	OrderProductNotAvailable OrderStatus = "product_not_available"
	OrderFinancialsSuccess   OrderStatus = "financials_success"
	OrderBalanceNotEnough    OrderStatus = "balance_not_enough"
	Paid                     OrderStatus = "paid"
	Canceled                 OrderStatus = "canceled"
)

type OrderFlow struct {
	Config       *Config
	OrderStatus  OrderStatus
	OrderID      int
	UserID       uuid.UUID
	ProductID    int
	ProductPrice decimal.Decimal
	Err          error
	Ctx          context.Context
}
