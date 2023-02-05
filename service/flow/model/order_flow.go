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
	OrderCreated             OrderStatus = "created"
	OrderSuccess             OrderStatus = "success"
	OrderInternalError       OrderStatus = "internal_error"
	OrderProductNotAvailable OrderStatus = "product_not_available"
	OrderBalanceNotEnough    OrderStatus = "balance_not_enough"
)

type OrderFlow struct {
	Config       *Config
	OrderStatus  OrderStatus
	OrderID      int
	UserID       uuid.UUID
	ProductID    int
	ProductPrice decimal.Decimal
}
