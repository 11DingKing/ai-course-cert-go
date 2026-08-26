package worker

import (
	"context"
	"sync"
	"time"
)

type Scheduler struct {
	mu     sync.Mutex
	jobs   map[string]func(context.Context)
	cancel context.CancelFunc
}

func NewScheduler() *Scheduler { return &Scheduler{jobs: map[string]func(context.Context){}} }
func (s *Scheduler) Register(name string, fn func(context.Context)) {
	s.mu.Lock()
	s.jobs[name] = fn
	s.mu.Unlock()
}
func (s *Scheduler) Run(ctx context.Context) {
	ctx, s.cancel = context.WithCancel(ctx)
	s.mu.Lock()
	cp := map[string]func(context.Context){}
	for k, v := range s.jobs {
		cp[k] = v
	}
	s.mu.Unlock()
	for _, fn := range cp {
		go fn(ctx)
	}
}
func (s *Scheduler) Stop() {
	s.mu.Lock()
	if s.cancel != nil {
		s.cancel()
	}
	s.mu.Unlock()
}
func Every(ctx context.Context, d time.Duration, fn func()) {
	if d <= 0 {
		d = time.Second
	}
	go func() {
		t := time.NewTicker(d)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				fn()
			}
		}
	}()
}
