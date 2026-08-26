package worker

import (
	"context"
	"database/sql"
	"sync"
	"time"
)

type Worker struct {
	DB       *sql.DB
	Interval time.Duration
	stop     chan struct{}
	once     sync.Once
}

func New(db *sql.DB) *Worker {
	return &Worker{DB: db, Interval: time.Second, stop: make(chan struct{})}
}
func (w *Worker) Start(ctx context.Context) {
	go func() {
		t := time.NewTicker(w.Interval)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-w.stop:
				return
			case <-t.C:
				w.process(ctx)
			}
		}
	}()
}
func (w *Worker) Stop() { w.once.Do(func() { close(w.stop) }) }
func (w *Worker) process(ctx context.Context) {
	rows, e := w.DB.QueryContext(ctx, `SELECT id FROM outbox_jobs WHERE status='pending' AND run_after<=? LIMIT 10`, time.Now().UTC())
	if e != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		rows.Scan(&id)
		w.DB.ExecContext(ctx, `UPDATE outbox_jobs SET status='done',updated_at=? WHERE id=?`, time.Now().UTC(), id)
	}
}
