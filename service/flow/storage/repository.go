package storage

import (
	"encoding/json"

	"github.com/geneva-lake/stateful_flow/general"
	"github.com/geneva-lake/stateful_flow/service/flow/model"
)

type Repository struct {
	pgsql *general.Pgsql
}

func NewRepository(connection *general.Pgsql) *Repository {
	return &Repository{
		pgsql: connection,
	}
}

func (r *Repository) OrderApply(order *StoredOrder) error {
	j, err := json.Marshal(order)
	if err != nil {
		return err
	}
	_, err = r.pgsql.Query("select orders.order_apply($1::json)", j)
	return err
}

func (r *Repository) OrderUpdate(orderID int, status model.OrderStatus) error {
	_, err := r.pgsql.Query("select orders.order_update($1::int, $2::text)", orderID, status)
	return err
}
