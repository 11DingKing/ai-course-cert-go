package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

type Outbox struct{ DB *sql.DB }

func (o Outbox) Enqueue(ctx context.Context, kind, payload string) (domain.OutboxJob, error) {
	now := time.Now().UTC()
	r, e := o.DB.ExecContext(ctx, `INSERT INTO outbox_jobs(kind,payload,status,attempts,run_after,last_error,created_at,updated_at) VALUES(?,?,?,0,?,?,?,?)`, kind, payload, "pending", now, "", now, now)
	if e != nil {
		return domain.OutboxJob{}, e
	}
	id, _ := r.LastInsertId()
	return domain.OutboxJob{ID: id, Kind: kind, Payload: payload, Status: "pending", RunAfter: now, CreatedAt: now, UpdatedAt: now}, nil
}
func (o Outbox) Claim(ctx context.Context, limit int) ([]domain.OutboxJob, error) {
	rows, e := o.DB.QueryContext(ctx, `SELECT id,kind,payload,status,attempts,run_after,last_error,created_at,updated_at FROM outbox_jobs WHERE status='pending' AND run_after<=? ORDER BY id LIMIT ?`, time.Now().UTC(), limit)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.OutboxJob{}
	for rows.Next() {
		var x domain.OutboxJob
		if e := rows.Scan(&x.ID, &x.Kind, &x.Payload, &x.Status, &x.Attempts, &x.RunAfter, &x.LastError, &x.CreatedAt, &x.UpdatedAt); e != nil {
			return nil, e
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
func (o Outbox) Complete(ctx context.Context, id int64) error {
	_, e := o.DB.ExecContext(ctx, `UPDATE outbox_jobs SET status='done',updated_at=? WHERE id=?`, time.Now().UTC(), id)
	return e
}
func (o Outbox) Fail(ctx context.Context, id int64, msg string) error {
	_, e := o.DB.ExecContext(ctx, `UPDATE outbox_jobs SET status=CASE WHEN attempts+1>=5 THEN 'dead' ELSE 'pending' END,attempts=attempts+1,last_error=?,updated_at=? WHERE id=?`, msg, time.Now().UTC(), id)
	return e
}
