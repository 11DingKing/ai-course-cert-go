package repository

import (
	"context"
	"database/sql"
	"time"
)

type Cleanup struct{ DB *sql.DB }

func (c Cleanup) ArchiveOld(ctx context.Context, before time.Time) (int64, error) {
	r, e := c.DB.ExecContext(ctx, `UPDATE submissions SET status='archived',updated_at=? WHERE status='approved' AND updated_at<?`, time.Now().UTC(), before)
	if e != nil {
		return 0, e
	}
	return r.RowsAffected()
}
func (c Cleanup) DeleteAuditBefore(ctx context.Context, before time.Time) (int64, error) {
	r, e := c.DB.ExecContext(ctx, `DELETE FROM audit_events WHERE created_at<?`, before)
	if e != nil {
		return 0, e
	}
	return r.RowsAffected()
}
func (c Cleanup) Vacuum(ctx context.Context) error { _, e := c.DB.ExecContext(ctx, `VACUUM`); return e }

var _ = sql.ErrNoRows
