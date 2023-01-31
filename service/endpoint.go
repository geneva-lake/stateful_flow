package service

import (
	"encoding/json"
	"net/http"

	"github.com/geneva-lake/stateful_flow/general"
	"github.com/geneva-lake/stateful_flow/logger"
	"github.com/geneva-lake/stateful_flow/service/flow"
	"github.com/geneva-lake/stateful_flow/service/flow/model"
)

//   - -------------------------------------------------------------------------------------------------------------------
//     Make user information endpoint where the chain of collecting user information
//     is started
//   - -------------------------------------------------------------------------------------------------------------------
func MakeOrderEndpoint(cfg *model.Config) general.Endpoint {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		defer func() {
			r := recover()
			if r != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp := model.OrderResponse{
					Status: general.StatusError,
					Error:  model.InternalError,
				}
				go logger.Log(logger.Panic, cfg.Name, http.StatusInternalServerError, nil, r, nil, resp)
				json.NewEncoder(w).Encode(resp)
				return
			}
		}()
		req, err := general.RequestDecode[model.OrderRequest](r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := model.OrderResponse{
				Status: general.StatusError,
				Error:  model.WrongRequest,
			}
			go logger.Log(logger.Panic, cfg.Name, http.StatusInternalServerError, nil, r, nil, resp)
			json.NewEncoder(w).Encode(resp)
			return
		}
		f := &model.OrderFlow{
			Config:       cfg,
			UserID:       req.UserID,
			ProductID:    req.ProductID,
			ProductPrice: req.ProductPrice,
		}
		flow.Process(f)
		res := model.OrderResult{
			OrderStatus: f.OrderStatus,
			OrderID:     f.OrderID,
		}
		resp := model.OrderResponse{
			Status: general.StatusOK,
			Result: &res,
		}
		json.NewEncoder(w).Encode(resp)
	}
}