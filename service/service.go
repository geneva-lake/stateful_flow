package service

import (
	"context"

	"github.com/geneva-lake/stateful_flow/service/flow"
	"github.com/geneva-lake/stateful_flow/service/flow/model"
)

type Service struct{}

func NewService(config *model.Config) *Service {
	s := Service{}
	return &s
}

func (s *Service) WithContext(ctx context.Context) *Service {
	s.ctx = ctx
	return s
}

func (s *Service) Start() {
	f := &model.OrderFlow{}
	flow.Start(f)
}
