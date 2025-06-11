package resilience

import (
	"context"
	"time"

	"github.com/slok/goresilience"
	"github.com/slok/goresilience/retry"
	"github.com/slok/goresilience/timeout"
)

type Executor struct {
	runner goresilience.Runner
}

type ExecutorConfig struct {
	Timeout    time.Duration
	MaxRetries int
}

func NewExecutor(cfg ExecutorConfig) *Executor {
	middlewares := []goresilience.Middleware{
		timeout.NewMiddleware(timeout.Config{Timeout: cfg.Timeout}),
		retry.NewMiddleware(retry.Config{Times: cfg.MaxRetries}),
	}
	return &Executor{
		runner: goresilience.RunnerChain(middlewares...),
	}
}

func (e *Executor) Run(ctx context.Context, fn func(ctx context.Context) error) error {
	return e.runner.Run(ctx, fn)
}
