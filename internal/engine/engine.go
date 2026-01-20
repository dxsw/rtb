package engine

import (
	"context"

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

func (e *Engine) Process(ctx context.Context, req *domain.BidRequest) *domain.BidResponse {
	return e.Pipeline.Process(ctx, req,
		e.Snapshot,
		e.Checker,
		e.Pacer,
	)
}
