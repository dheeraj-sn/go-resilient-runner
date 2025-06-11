package service

import (
	"context"
	"errors"
	"time"
)

func GetBalance(ctx context.Context, uid int64) (int, error) {
	return 10, nil
}

func GetScratchCardInfo(ctx context.Context, uid int64) (int64, error) {
	select {
	case <-time.After(25 * time.Millisecond):
		return 1000, nil
	case <-ctx.Done():
		return 0, errors.New("timeout")
	}
}
