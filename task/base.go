package task

import (
	"context"
	"time"
)

type DependencyType int

const (
	Critical DependencyType = iota
	BestEffort
)

type Task interface {
	Name() string
	DependencyType() DependencyType
	Timeout() time.Duration
	Execute(ctx context.Context) (any, error)
}
