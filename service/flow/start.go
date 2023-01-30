package flow

import (
	"github.com/geneva-lake/stateful_flow/general"
	"github.com/geneva-lake/stateful_flow/service/flow/bookkeeping"
	"github.com/geneva-lake/stateful_flow/service/flow/financials"
	"github.com/geneva-lake/stateful_flow/service/flow/model"
	"github.com/geneva-lake/stateful_flow/service/flow/storage"
)

func Start(flow *model.OrderFlow) {
	repo := storage.NewRepository(general.NewPgsql(flow.Config.DBConnectionString))
	bu := (*bookkeeping.BookkepingUnit)(flow)
	fu := (*financials.FinancialsUnit)(flow)
	su := (*storage.StorageUnit)(flow)
	start := model.NewStatusStream()
	storage2bookkeeping := model.NewStatusStream()
	bookkeeping2financials := model.NewStatusStream()
	go su.Process(repo, start, storage2bookkeeping)
	go bu.Process(storage2bookkeeping, bookkeeping2financials)
	go fu.Process(bookkeeping2financials)
	start.Forward <- model.Proceed
	<-start.Back
}
