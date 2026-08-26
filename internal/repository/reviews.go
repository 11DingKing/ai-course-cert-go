package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

type Reviews struct{ DB *sql.DB }

func (r Reviews) Add(ctx context.Context, x domain.Review) (domain.Review, error) {
	x.CreatedAt = time.Now().UTC()
	res, e := r.DB.ExecContext(ctx, `INSERT INTO reviews(submission_id,reviewer_id,decision,comment,version,created_at) VALUES(?,?,?,?,?,?)`, x.SubmissionID, x.ReviewerID, x.Decision, x.Comment, x.Version, x.CreatedAt)
	if e == nil {
		x.ID, _ = res.LastInsertId()
	}
	return x, e
}
func (r Reviews) List(ctx context.Context, id int64) ([]domain.Review, error) {
	rows, e := r.DB.QueryContext(ctx, `SELECT id,submission_id,reviewer_id,decision,comment,version,created_at FROM reviews WHERE submission_id=? ORDER BY created_at`, id)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.Review{}
	for rows.Next() {
		var x domain.Review
		var cr string
		if e := rows.Scan(&x.ID, &x.SubmissionID, &x.ReviewerID, &x.Decision, &x.Comment, &x.Version, &cr); e != nil {
			return nil, e
		}
		x.CreatedAt, _ = time.Parse(time.RFC3339Nano, cr)
		out = append(out, x)
	}
	return out, rows.Err()
}
