package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type FlowStatus int

const (
	Proceed FlowStatus = 1
	Cancel  FlowStatus = 2
)

type StatusStream struct {
	Forward chan FlowStatus
	Back    chan FlowStatus
}

func NewStatusStream() *StatusStream {
	stream := StatusStream{
		Forward: make(chan FlowStatus),
		Back:    make(chan FlowStatus),
	}
	return &stream
}

type OrderStatus string

const (
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
}
