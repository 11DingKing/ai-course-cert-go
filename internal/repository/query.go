package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"strings"
)

type Query struct{ DB *sql.DB }

func (q Query) SearchSubmissions(ctx context.Context, course int64, status, term string, limit, offset int) ([]domain.Submission, error) {
	where := []string{"course_id=?"}
	args := []any{course}
	if status != "" {
		where = append(where, "status=?")
		args = append(args, status)
	}
	if term != "" {
		where = append(where, "(title LIKE ? OR ai_statement LIKE ?)")
		args = append(args, "%"+term+"%", "%"+term+"%")
	}
	args = append(args, limit, offset)
	sqlq := fmt.Sprintf(`SELECT id,course_id,student_id,status,title,ai_statement,due_at,version,submitted_at,created_at,updated_at FROM submissions WHERE %s ORDER BY updated_at DESC LIMIT ? OFFSET ?`, strings.Join(where, " AND "))
	rows, e := q.DB.QueryContext(ctx, sqlq, args...)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.Submission{}
	for rows.Next() {
		var x domain.Submission
		var st string
		if e := rows.Scan(&x.ID, &x.CourseID, &x.StudentID, &st, &x.Title, &x.AIStatement, &x.DueAt, &x.Version, &x.SubmittedAt, &x.CreatedAt, &x.UpdatedAt); e != nil {
			return nil, e
		}
		x.Status = domain.SubmissionStatus(st)
		out = append(out, x)
	}
	return out, rows.Err()
}
func (q Query) Count(ctx context.Context, course int64, status string) (int, error) {
	var n int
	e := q.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM submissions WHERE course_id=? AND (?='' OR status=?)`, course, status, status).Scan(&n)
	return n, e
}
