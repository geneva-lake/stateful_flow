package financials

import (
	"github.com/geneva-lake/stateful_flow/general"
	"github.com/shopspring/decimal"
)

type ApplyRequest struct {
	OrderID      int
	ProductID    int
	ProductPrice decimal.Decimal
}

type Status string

const (
	Success          Status = "success"
	BalanceNotEnough Status = "balance_not_enough"
)

type Response struct {
	Status general.ResponseStatus
	Error  Error
	Result *Result
}

type Result struct {
	Status        Status
	TransactionID *string
}

type Error string

const (
	Internalerror Error = "internal_error"
)
