package bookkeeping

import (
	"github.com/geneva-lake/stateful_flow/general"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type RequestStatus string

const (
	OrderCreated  RequestStatus = "order_created"
	OrderCanceled RequestStatus = "order_canceled"
	OrderPaid     RequestStatus = "order_paid"
)

type ApplyRequest struct {
	UserID       uuid.UUID
	OrderID      int
	ProductID    int
	ProductPrice decimal.Decimal
}

type Status string

const (
	Success             Status = "success"
	ProductNotAvailable Status = "product_not_available"
)

type ApplyResponse struct {
	Status general.ResponseStatus
	Error  Error
	Result *ApplyResult
}

type ApplyResult struct {
	Status   Status
	RecordID *uuid.UUID
}

type Error string

const (
	ProductNotFound Error = "product_not_found"
	Internalerror   Error = "internal_error"
)

type UpdateResponse struct {
	Status general.ResponseStatus
}
