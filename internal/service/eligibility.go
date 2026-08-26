package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
)

type Eligibility struct{ DB *sql.DB }

func (e Eligibility) Check(ctx context.Context, student, course int64) (bool, error) {
	var role string
	if err := e.DB.QueryRowContext(ctx, `SELECT role FROM users WHERE id=?`, student).Scan(&role); err != nil {
		return false, err
	}
	if role != "student" {
		return false, fmt.Errorf("not student")
	}
	var cap int
	var used int
	if err := e.DB.QueryRowContext(ctx, `SELECT capacity FROM courses WHERE id=?`, course).Scan(&cap); err != nil {
		return false, err
	}
	if err := e.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM submissions WHERE course_id=?`, course).Scan(&used); err != nil {
		return false, err
	}
	return domain.CapacityAvailable(used, cap), nil
}
func (e Eligibility) HasApproved(ctx context.Context, student int64) bool {
	var n int
	e.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM submissions WHERE student_id=? AND status='approved'`, student).Scan(&n)
	return n > 0
}
func (e Eligibility) ModuleCoverage(ctx context.Context, submission int64) int {
	var n int
	e.DB.QueryRowContext(ctx, `SELECT COUNT(*) FROM evidences WHERE submission_id=?`, submission).Scan(&n)
	return n
}
