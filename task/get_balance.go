package task

import (
	"context"
	"time"

	"github.com/dheeraj-sn/go-resilient-runner/resilience"
	"github.com/dheeraj-sn/go-resilient-runner/service"
)

type GetBalance struct {
	UID      int64
	Executor *resilience.Executor
}

func NewGetBalance(uid int64) GetBalance {
	return GetBalance{
		UID: uid,
		Executor: resilience.NewExecutor(resilience.ExecutorConfig{
			Timeout:    30 * time.Millisecond,
			MaxRetries: 1,
		}),
	}
}

func (w GetBalance) Name() string {
	return "get_balance"
}

func (w GetBalance) DependencyType() DependencyType {
	return Critical
}

func (w GetBalance) Timeout() time.Duration {
	return 30 * time.Millisecond
}

func (w GetBalance) Execute(ctx context.Context) (any, error) {
	var balance int
	err := w.Executor.Run(ctx, func(ctx context.Context) error {
		val, err := service.GetBalance(ctx, w.UID)
		if err != nil {
			return err
		}
		balance = val
		return nil
	})
	return balance, err
}
