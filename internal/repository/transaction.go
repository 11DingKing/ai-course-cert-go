package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

type Transaction struct{ DB *sql.DB }

func (t Transaction) CreateSubmissionWithEvidence(ctx context.Context, s domain.Submission, e domain.Evidence) (domain.Submission, domain.Evidence, error) {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return s, e, err
	}
	defer tx.Rollback()
	now := time.Now().UTC()
	s.Status = domain.StatusDraft
	s.Version = 1
	s.CreatedAt = now
	s.UpdatedAt = now
	s.DueAt = now.Add(24 * time.Hour)
	r, err := tx.ExecContext(ctx, `INSERT INTO submissions(course_id,student_id,status,title,ai_statement,due_at,version,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?)`, s.CourseID, s.StudentID, s.Status, s.Title, s.AIStatement, s.DueAt, s.Version, s.CreatedAt, s.UpdatedAt)
	if err != nil {
		return s, e, err
	}
	s.ID, _ = r.LastInsertId()
	e.SubmissionID = s.ID
	e.Version = 1
	e.CreatedAt = now
	r, err = tx.ExecContext(ctx, `INSERT INTO evidences(submission_id,version,uri,checksum,description,created_at) VALUES(?,?,?,?,?,?)`, e.SubmissionID, e.Version, e.URI, e.Checksum, e.Description, e.CreatedAt)
	if err != nil {
		return s, e, err
	}
	e.ID, _ = r.LastInsertId()
	if err = tx.Commit(); err != nil {
		return s, e, err
	}
	return s, e, nil
}
func (t Transaction) InTx(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, e := t.DB.BeginTx(ctx, nil)
	if e != nil {
		return e
	}
	if e = fn(tx); e != nil {
		tx.Rollback()
		return e
	}
	return tx.Commit()
}
func (t Transaction) LockSubmission(ctx context.Context, id int64) (*sql.Tx, error) {
	tx, e := t.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if e != nil {
		return nil, e
	}
	var n int
	if e = tx.QueryRowContext(ctx, `SELECT id FROM submissions WHERE id=?`, id).Scan(&n); e != nil {
		tx.Rollback()
		return nil, e
	}
	return tx, nil
}
