package engine

import (
	"gitlab.com/radix-ads/rtb/internal/budget"
	"gitlab.com/radix-ads/rtb/internal/domain"
	"gitlab.com/radix-ads/rtb/internal/engine/pipeline"
	"gitlab.com/radix-ads/rtb/internal/pacing"
	"gitlab.com/radix-ads/rtb/internal/store"
)

type Engine struct {
	Pipeline *pipeline.Pipeline

	Snapshot *store.Snapshot
	Checker  *budget.Checker
	Pacer    *pacing.Pacer
}

func (e *Engine) Process(req *domain.BidRequest) *domain.BidResponse {
	return e.Pipeline.Process(req,
		e.Snapshot,
		e.Checker,
		e.Pacer,
	)
}
