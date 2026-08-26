package worker

import (
	"context"
	"time"
)

type RetryPolicy struct {
	Max  int
	Base time.Duration
}

func (p RetryPolicy) Delay(attempt int) time.Duration {
	if attempt < 0 {
		attempt = 0
	}
	if attempt > 10 {
		attempt = 10
	}
	d := p.Base
	for i := 0; i < attempt; i++ {
		d *= 2
	}
	return d
}
func Run(ctx context.Context, p RetryPolicy, fn func(context.Context) error) error {
	var e error
	max := p.Max
	if max <= 0 {
		max = 1
	}
	for i := 0; i < max; i++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		e = fn(ctx)
		if e == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(p.Delay(i)):
		}
	}
	return e
}
