package repository

import (
	"context"
	"database/sql"
	"time"
)

type Stats struct{ DB *sql.DB }

func (s Stats) StatusCounts(ctx context.Context, course int64) (map[string]int, error) {
	rows, e := s.DB.QueryContext(ctx, `SELECT status,COUNT(*) FROM submissions WHERE course_id=? GROUP BY status`, course)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := map[string]int{}
	for rows.Next() {
		var k string
		var n int
		if e := rows.Scan(&k, &n); e != nil {
			return nil, e
		}
		out[k] = n
	}
	return out, rows.Err()
}
func (s Stats) ActiveCourses(ctx context.Context, now time.Time) (int, error) {
	var n int
	e := s.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM courses WHERE opens_at<=? AND closes_at>?`, now, now).Scan(&n)
	return n, e
}
func (s Stats) EvidenceCount(ctx context.Context, student int64) (int, error) {
	var n int
	e := s.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM evidences e JOIN submissions s ON s.id=e.submission_id WHERE s.student_id=?`, student).Scan(&n)
	return n, e
}
