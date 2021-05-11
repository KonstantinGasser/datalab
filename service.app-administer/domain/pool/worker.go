package pool

import (
	"context"
	"sync"
)

type Job func() (interface{}, error)
type JobResult struct {
	Value interface{}
	Err   error
}

func New(ctx context.Context, workerCount int) (chan<- Job, <-chan JobResult) {
	var wg sync.WaitGroup

	jobs := make(chan Job, workerCount)
	results := make(chan JobResult)

	go watch(ctx, &wg, results)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go work(ctx, &wg, jobs, results)
	}
	return jobs, results
}

func work(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, result chan<- JobResult) {
	for job := range jobs {
		value, err := job()
		result <- JobResult{Value: value, Err: err}
		wg.Done()
	}
}

func watch(ctx context.Context, wg *sync.WaitGroup, results chan JobResult) {
	wg.Wait()
	close(results)
}
