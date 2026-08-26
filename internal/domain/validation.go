package domain

import (
	"strings"
	"time"
)

func ValidateEmail(v string) bool { return strings.Contains(v, "@") && len(v) >= 5 }
func ValidateCourse(c Course) error {
	if c.Code == "" || c.Title == "" || c.Semester == "" {
		return ErrValidation
	}
	if c.Capacity < 1 {
		return ErrValidation
	}
	if c.ClosesAt.Before(c.OpensAt) {
		return ErrValidation
	}
	return nil
}
func ValidateSubmission(s Submission) error {
	if s.CourseID <= 0 || s.StudentID <= 0 || strings.TrimSpace(s.Title) == "" {
		return ErrValidation
	}
	if strings.TrimSpace(s.AIStatement) == "" {
		return ErrValidation
	}
	return nil
}
func IsOpen(c Course, now time.Time) bool { return !now.Before(c.OpensAt) && now.Before(c.ClosesAt) }

var ErrValidation = &ValidationError{}

type ValidationError struct{}

func (*ValidationError) Error() string { return "validation failed" }
func AllowedRole(r Role, allowed ...Role) bool {
	for _, x := range allowed {
		if r == x {
			return true
		}
	}
	return false
}
func (s Submission) IsTerminal() bool { return s.Status == StatusArchived }
func (s Submission) CanEdit(u User) bool {
	return u.Role == RoleAdmin || (u.Role == RoleStudent && u.ID == s.StudentID)
}
func (e Evidence) Complete() bool       { return e.URI != "" && e.Checksum != "" && e.Description != "" }
func ReviewDecisionValid(v string) bool { return v == "approve" || v == "return" }
