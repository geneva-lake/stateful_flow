package main

import (
	"net/http"

	"github.com/geneva-lake/stateful_flow/logger"
	"github.com/geneva-lake/stateful_flow/service/flow/model"
)

func main() {
	cfg, err := model.NewConfig().FromFile("config.yaml").Yaml()
	if err != nil {
		logger.LogErrorMain(err)
		return
	}
	handlers := CreateHandlers(cfg)
	if err := http.ListenAndServe(":"+cfg.Port, handlers); err != nil {
		logger.LogErrorMain(err)
	}
}
