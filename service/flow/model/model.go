package model

import (
	"github.com/geneva-lake/stateful_flow/general"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderRequest struct {
	UserID       uuid.UUID
	ProductID    int
	ProductPrice decimal.Decimal
}

type OrderError string

const (
	WrongRequest  OrderError = "wrong_request"
	InternalError OrderError = "internal_error"
)

type OrderResponse struct {
	Status general.ResponseStatus
	Error  OrderError
	Result *OrderResult
}

type OrderResult struct {
	OrderStatus OrderStatus
	OrderID     int
}
