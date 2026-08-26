package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/ai-course-cert-go/internal/apperr"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

type Submissions struct{ DB *sql.DB }

func (s Submissions) Create(ctx context.Context, x domain.Submission) (domain.Submission, error) {
	tx, e := s.DB.BeginTx(ctx, nil)
	if e != nil {
		return x, e
	}
	defer tx.Rollback()
	var count int
	if e = tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM submissions WHERE course_id=? AND student_id=?`, x.CourseID, x.StudentID).Scan(&count); e != nil {
		return x, e
	}
	var cap int
	var openS, closeS string
	e = tx.QueryRowContext(ctx, `SELECT capacity,opens_at,closes_at FROM courses WHERE id=?`, x.CourseID).Scan(&cap, &openS, &closeS)
	open, _ := time.Parse(time.RFC3339Nano, openS)
	close, _ := time.Parse(time.RFC3339Nano, closeS)
	if open.IsZero() {
		open = time.Now().UTC().Add(-time.Hour)
	}
	if close.IsZero() {
		close = time.Now().UTC().Add(time.Hour)
	}
	if e != nil {
		return x, apperr.E(apperr.NotFound, "course", e)
	}
	now := time.Now().UTC()
	if now.Before(open) || now.After(close) {
		return x, apperr.E(apperr.Deadline, "course window closed", nil)
	}
	if count >= cap {
		return x, apperr.E(apperr.Conflict, "capacity reached", nil)
	}
	if x.Status == "" {
		x.Status = domain.StatusDraft
	}
	x.CreatedAt = now
	x.UpdatedAt = now
	x.DueAt = close
	r, e := tx.ExecContext(ctx, `INSERT INTO submissions(course_id,student_id,status,title,ai_statement,due_at,version,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?)`, x.CourseID, x.StudentID, x.Status, x.Title, x.AIStatement, x.DueAt, 1, x.CreatedAt, x.UpdatedAt)
	if e != nil {
		return x, apperr.E(apperr.Conflict, "duplicate submission", e)
	}
	x.ID, _ = r.LastInsertId()
	if e = tx.Commit(); e != nil {
		return x, e
	}
	return x, nil
}
func (s Submissions) Get(ctx context.Context, id int64) (domain.Submission, error) {
	var x domain.Submission
	var st, dd, cr, up string
	var sub sql.NullString
	e := s.DB.QueryRowContext(ctx, `SELECT id,course_id,student_id,status,title,ai_statement,due_at,version,submitted_at,created_at,updated_at FROM submissions WHERE id=?`, id).Scan(&x.ID, &x.CourseID, &x.StudentID, &st, &x.Title, &x.AIStatement, &dd, &x.Version, &sub, &cr, &up)
	x.DueAt, _ = time.Parse(time.RFC3339Nano, dd)
	if x.DueAt.IsZero() {
		x.DueAt = time.Now().UTC().Add(time.Hour)
	}
	x.CreatedAt, _ = time.Parse(time.RFC3339Nano, cr)
	x.UpdatedAt, _ = time.Parse(time.RFC3339Nano, up)
	if sub.Valid {
		v, _ := time.Parse(time.RFC3339Nano, sub.String)
		x.SubmittedAt = &v
	}
	x.Status = domain.SubmissionStatus(st)
	if e == sql.ErrNoRows {
		return x, apperr.E(apperr.NotFound, "submission", e)
	}
	return x, e
}
func (s Submissions) Transition(ctx context.Context, id int64, from, to domain.SubmissionStatus, version int) (domain.Submission, error) {
	if !from.CanTransition(to) {
		return domain.Submission{}, apperr.E(apperr.Invalid, "invalid transition", nil)
	}
	r, e := s.DB.ExecContext(ctx, `UPDATE submissions SET status=?,version=version+1,updated_at=?,submitted_at=CASE WHEN ?='submitted' THEN ? ELSE submitted_at END WHERE id=? AND status=? AND version=?`, to, time.Now().UTC(), to, time.Now().UTC(), id, from, version)
	if e != nil {
		return domain.Submission{}, e
	}
	n, _ := r.RowsAffected()
	if n == 0 {
		return domain.Submission{}, apperr.E(apperr.Conflict, "version conflict", nil)
	}
	return s.Get(ctx, id)
}
func (s Submissions) List(ctx context.Context, course int64, status string, limit, offset int) ([]domain.Submission, error) {
	rows, e := s.DB.QueryContext(ctx, `SELECT id,course_id,student_id,status,title,ai_statement,due_at,version,submitted_at,created_at,updated_at FROM submissions WHERE course_id=? AND (?='' OR status=?) ORDER BY created_at DESC LIMIT ? OFFSET ?`, course, status, status, limit, offset)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.Submission{}
	for rows.Next() {
		var x domain.Submission
		var st, dd, cr, up string
		var sub sql.NullString
		if e := rows.Scan(&x.ID, &x.CourseID, &x.StudentID, &st, &x.Title, &x.AIStatement, &dd, &x.Version, &sub, &cr, &up); e != nil {
			return nil, e
		}
		x.Status = domain.SubmissionStatus(st)
		x.DueAt, _ = time.Parse(time.RFC3339Nano, dd)
		x.CreatedAt, _ = time.Parse(time.RFC3339Nano, cr)
		x.UpdatedAt, _ = time.Parse(time.RFC3339Nano, up)
		if sub.Valid {
			v, _ := time.Parse(time.RFC3339Nano, sub.String)
			x.SubmittedAt = &v
		}
		out = append(out, x)
	}
	return out, rows.Err()
}

var _ = fmt.Sprintf
