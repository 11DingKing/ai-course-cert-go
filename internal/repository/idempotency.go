package repository

import (
	"context"
	"database/sql"
	"time"
)

type Idempotency struct{ DB *sql.DB }

func (i Idempotency) Find(ctx context.Context, key, op string, user int64) (int64, bool, error) {
	var id int64
	e := i.DB.QueryRowContext(ctx, `SELECT result_id FROM idempotency_records WHERE key=? AND operation=? AND user_id=?`, key, op, user).Scan(&id)
	if e == sql.ErrNoRows {
		return 0, false, nil
	}
	return id, e == nil, e
}
func (i Idempotency) Save(ctx context.Context, key, op string, user, id int64) error {
	_, e := i.DB.ExecContext(ctx, `INSERT INTO idempotency_records(key,operation,user_id,result_id,created_at) VALUES(?,?,?,?,?)`, key, op, user, id, time.Now().UTC())
	return e
}
func (i Idempotency) Purge(ctx context.Context, before time.Time) (int64, error) {
	r, e := i.DB.ExecContext(ctx, `DELETE FROM idempotency_records WHERE created_at<?`, before)
	if e != nil {
		return 0, e
	}
	n, _ := r.RowsAffected()
	return n, nil
}
