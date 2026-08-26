package audit

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

func List(ctx context.Context, db *sql.DB, entity string, id int64, limit int) ([]domain.AuditEvent, error) {
	rows, e := db.QueryContext(ctx, `SELECT id,actor_id,entity_type,entity_id,action,result,request_id,metadata,created_at FROM audit_events WHERE entity_type=? AND entity_id=? ORDER BY id DESC LIMIT ?`, entity, id, limit)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.AuditEvent{}
	for rows.Next() {
		var x domain.AuditEvent
		if e := rows.Scan(&x.ID, &x.ActorID, &x.EntityType, &x.EntityID, &x.Action, &x.Result, &x.RequestID, &x.Metadata, &x.CreatedAt); e != nil {
			return nil, e
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
func Since(ctx context.Context, db *sql.DB, since time.Time) (int, error) {
	var n int
	e := db.QueryRowContext(ctx, `SELECT COUNT(*) FROM audit_events WHERE created_at>=?`, since).Scan(&n)
	return n, e
}
