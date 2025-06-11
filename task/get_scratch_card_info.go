package task

import (
	"context"
	"time"

	"github.com/dheeraj-sn/go-resilient-runner/resilience"
	"github.com/dheeraj-sn/go-resilient-runner/service"
)

type GetScratchCardInfo struct {
	UID      int64
	Executor *resilience.Executor
}

func NewGetScratchCardInfo(uid int64) GetScratchCardInfo {
	return GetScratchCardInfo{
		UID: uid,
		Executor: resilience.NewExecutor(resilience.ExecutorConfig{
			Timeout:    20 * time.Millisecond,
			MaxRetries: 1,
		}),
	}
}

func (v GetScratchCardInfo) Name() string {
	return "get_scratch_card_info"
}

func (v GetScratchCardInfo) DependencyType() DependencyType {
	return BestEffort
}

func (v GetScratchCardInfo) Timeout() time.Duration {
	return 20 * time.Millisecond
}

func (v GetScratchCardInfo) Execute(ctx context.Context) (any, error) {
	var scratchCardID int64
	err := v.Executor.Run(ctx, func(ctx context.Context) error {
		val, err := service.GetScratchCardInfo(ctx, v.UID)
		if err != nil {
			return err
		}
		scratchCardID = val
		return nil
	})
	if err != nil {
		scratchCardID = 0 // fallback default
	}
	return scratchCardID, nil
}
