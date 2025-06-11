package orchestrator

import (
	"context"
	"fmt"
	"sync"

	"github.com/dheeraj-sn/go-resilient-runner/task"
)

type TaskResult struct {
	Name  string
	Value any
	Error error
}

type TaskRunner struct {
	Tasks []task.Task
}

func (tr *TaskRunner) AddTask(t task.Task) {
	tr.Tasks = append(tr.Tasks, t)
}

func (tr *TaskRunner) RunAll(ctx context.Context) ([]TaskResult, error) {
	var wg sync.WaitGroup
	resultCh := make(chan TaskResult, len(tr.Tasks))
	errCh := make(chan error, 1)

	for _, t := range tr.Tasks {
		wg.Add(1)
		go func(t task.Task) {
			defer wg.Done()
			taskCtx, cancel := context.WithTimeout(ctx, t.Timeout())
			defer cancel()

			val, err := t.Execute(taskCtx)
			if err != nil && t.DependencyType() == task.Critical {
				select {
				case errCh <- fmt.Errorf("critical task %s failed: %w", t.Name(), err):
				default:
				}
			}
			resultCh <- TaskResult{Name: t.Name(), Value: val, Error: err}
		}(t)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []TaskResult
	for {
		select {
		case res, ok := <-resultCh:
			if !ok {
				return results, nil
			}
			results = append(results, res)
		case err := <-errCh:
			return nil, err
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
