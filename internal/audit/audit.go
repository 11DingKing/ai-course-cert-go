package audit

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

type Logger struct{ DB *sql.DB }

func (l Logger) Record(ctx context.Context, actor int64, typ string, id int64, action, result, request string, meta any) error {
	b, _ := json.Marshal(meta)
	_, e := l.DB.ExecContext(ctx, `INSERT INTO audit_events(actor_id,entity_type,entity_id,action,result,request_id,metadata,created_at) VALUES(?,?,?,?,?,?,?,?)`, actor, typ, id, action, result, request, string(b), time.Now().UTC())
	return e
}
