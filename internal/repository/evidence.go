package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

type Evidences struct{ DB *sql.DB }

func (e Evidences) Add(ctx context.Context, x domain.Evidence) (domain.Evidence, error) {
	var v int
	e.DB.QueryRowContext(ctx, `SELECT COALESCE(MAX(version),0)+1 FROM evidences WHERE submission_id=?`, x.SubmissionID).Scan(&v)
	x.Version = v
	x.CreatedAt = time.Now().UTC()
	r, err := e.DB.ExecContext(ctx, `INSERT INTO evidences(submission_id,version,uri,checksum,description,created_at) VALUES(?,?,?,?,?,?)`, x.SubmissionID, x.Version, x.URI, x.Checksum, x.Description, x.CreatedAt)
	if err == nil {
		x.ID, _ = r.LastInsertId()
	}
	return x, err
}
func (e Evidences) List(ctx context.Context, id int64) ([]domain.Evidence, error) {
	rows, err := e.DB.QueryContext(ctx, `SELECT id,submission_id,version,uri,checksum,description,created_at FROM evidences WHERE submission_id=? ORDER BY version`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []domain.Evidence{}
	for rows.Next() {
		var x domain.Evidence
		var cr string
		if err := rows.Scan(&x.ID, &x.SubmissionID, &x.Version, &x.URI, &x.Checksum, &x.Description, &cr); err != nil {
			return nil, err
		}
		x.CreatedAt, _ = time.Parse(time.RFC3339Nano, cr)
		out = append(out, x)
	}
	return out, rows.Err()
}
