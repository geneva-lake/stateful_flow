package main

import (
	"net/http"

	"github.com/geneva-lake/stateful_flow/service"
	"github.com/geneva-lake/stateful_flow/service/flow/model"
	"github.com/gorilla/mux"
)

//   - -------------------------------------------------------------------------------------------------------------------
//     Create handlers for http request
//   - -------------------------------------------------------------------------------------------------------------------
func CreateHandlers(cfg *model.Config) http.Handler {
	e := service.MakeOrderEndpoint(cfg)
	r := mux.NewRouter()
	r.Methods("POST").Path("/orders/apply").HandlerFunc(e)
	return r
}
