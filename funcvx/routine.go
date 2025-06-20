package funcvx

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

// RunWithConcurrencyLimit runs a function with a concurrency limit
// It's useful for running functions in parallel with a limit on the number of concurrent goroutines
// @param ctx context.Context - The context of the function
// @param inputs []T - The inputs of the function
// @param limit int - The limit of the number of concurrent goroutines
// @param handler func(ctx context.Context, input T) error - The function to run
// @return []error - The errors of the function
func RunWithConcurrencyLimit[T any](
	ctx context.Context,
	inputs []T,
	limit int,
	handler func(ctx context.Context, input T) error,
) []error {
	var (
		mutex  sync.Mutex
		errors []error
	)
	g, gCtx := errgroup.WithContext(ctx)
	g.SetLimit(limit)

	//
	for _, input := range inputs {
		input := input
		g.Go(func() error {
			err := handler(gCtx, input)
			if err != nil {
				mutex.Lock() // Lock để tránh race condition
				errors = append(errors, err)
				mutex.Unlock() // Unlock sau khi thêm vào errors
			}
			return nil
		})
	}
	_ = g.Wait()
	return errors
}
