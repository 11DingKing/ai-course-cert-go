package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/apperr"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

type Courses struct{ DB *sql.DB }

func (c Courses) Create(ctx context.Context, x domain.Course) (domain.Course, error) {
	r, e := c.DB.ExecContext(ctx, `INSERT INTO courses(code,title,semester,opens_at,closes_at,capacity,created_at) VALUES(?,?,?,?,?,?,?)`, x.Code, x.Title, x.Semester, x.OpensAt.UTC(), x.ClosesAt.UTC(), x.Capacity, time.Now().UTC())
	if e != nil {
		return x, apperr.E(apperr.Conflict, "course exists", e)
	}
	x.ID, _ = r.LastInsertId()
	return x, nil
}
func (c Courses) Get(ctx context.Context, id int64) (domain.Course, error) {
	var x domain.Course
	var o, cl, cr string
	e := c.DB.QueryRowContext(ctx, `SELECT id,code,title,semester,opens_at,closes_at,capacity,created_at FROM courses WHERE id=?`, id).Scan(&x.ID, &x.Code, &x.Title, &x.Semester, &o, &cl, &x.Capacity, &cr)
	x.OpensAt, _ = time.Parse(time.RFC3339Nano, o)
	x.ClosesAt, _ = time.Parse(time.RFC3339Nano, cl)
	x.CreatedAt, _ = time.Parse(time.RFC3339Nano, cr)
	if e == sql.ErrNoRows {
		return x, apperr.E(apperr.NotFound, "course", e)
	}
	return x, e
}
func (c Courses) List(ctx context.Context, limit, offset int) ([]domain.Course, error) {
	rows, e := c.DB.QueryContext(ctx, `SELECT id,code,title,semester,opens_at,closes_at,capacity,created_at FROM courses ORDER BY opens_at LIMIT ? OFFSET ?`, limit, offset)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.Course{}
	for rows.Next() {
		var x domain.Course
		var o, cl, cr string
		if e := rows.Scan(&x.ID, &x.Code, &x.Title, &x.Semester, &o, &cl, &x.Capacity, &cr); e != nil {
			return nil, e
		}
		x.OpensAt, _ = time.Parse(time.RFC3339Nano, o)
		x.ClosesAt, _ = time.Parse(time.RFC3339Nano, cl)
		x.CreatedAt, _ = time.Parse(time.RFC3339Nano, cr)
		out = append(out, x)
	}
	return out, rows.Err()
}
