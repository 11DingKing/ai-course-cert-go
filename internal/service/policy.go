package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"strings"
	"time"
)

type Policy struct {
	DB    *sql.DB
	Clock func() time.Time
}

func (p Policy) now() time.Time {
	if p.Clock != nil {
		return p.Clock()
	}
	return time.Now().UTC()
}
func (p Policy) CanSubmit(ctx context.Context, s domain.Submission) error {
	if s.Status != domain.StatusDraft && s.Status != domain.StatusReturned {
		return fmt.Errorf("status %s cannot submit", s.Status)
	}
	if p.now().After(s.DueAt) {
		return fmt.Errorf("deadline exceeded")
	}
	return nil
}
func (p Policy) CanReview(ctx context.Context, u domain.User, s domain.Submission) error {
	if !domain.AllowedRole(u.Role, domain.RoleReviewer, domain.RoleTeacher, domain.RoleAdmin) {
		return fmt.Errorf("role denied")
	}
	if !domain.IsReviewable(s.Status) {
		return fmt.Errorf("not reviewable")
	}
	return nil
}
func (p Policy) CanArchive(ctx context.Context, u domain.User, s domain.Submission) error {
	if u.Role != domain.RoleAdmin {
		return fmt.Errorf("admin required")
	}
	if s.Status != domain.StatusApproved {
		return fmt.Errorf("approval required")
	}
	return nil
}
func (p Policy) CheckAIStatement(v string) bool {
	return strings.Contains(strings.ToLower(v), "ai") || strings.Contains(v, "人工智能")
}
func (p Policy) CheckTitle(v string) bool { return len(strings.TrimSpace(v)) >= 3 && len(v) <= 200 }
func (p Policy) CheckURI(v string) bool {
	return strings.HasPrefix(v, "https://") || strings.HasPrefix(v, "s3://")
}
func (p Policy) CheckChecksum(v string) bool { return len(v) >= 8 }
func (p Policy) NormalizeComment(v string) string {
	return strings.TrimSpace(strings.ReplaceAll(v, "\x00", ""))
}
func (p Policy) DueDate(c domain.Course) time.Time   { return c.ClosesAt }
func (p Policy) IsLate(due, now time.Time) bool      { return now.After(due) }
func (p Policy) GracePeriod(due time.Time) time.Time { return due.Add(72 * time.Hour) }
func (p Policy) InGrace(due, now time.Time) bool {
	return now.After(due) && now.Before(p.GracePeriod(due))
}
func (p Policy) MaxEvidence() int     { return 20 }
func (p Policy) MaxTitleBytes() int   { return 200 }
func (p Policy) MaxCommentBytes() int { return 4000 }
func (p Policy) ValidateEvidence(e domain.Evidence) error {
	if !p.CheckURI(e.URI) {
		return fmt.Errorf("uri must be https")
	}
	if !p.CheckChecksum(e.Checksum) {
		return fmt.Errorf("checksum missing")
	}
	if !e.Complete() {
		return fmt.Errorf("incomplete evidence")
	}
	return nil
}
func (p Policy) ValidateReview(r domain.Review) error {
	if !domain.ReviewDecisionValid(r.Decision) {
		return fmt.Errorf("decision invalid")
	}
	if p.NormalizeComment(r.Comment) == "" {
		return fmt.Errorf("comment required")
	}
	return nil
}
func (p Policy) ValidateCourse(c domain.Course) error         { return domain.ValidateCourse(c) }
func (p Policy) ValidateSubmission(s domain.Submission) error { return domain.ValidateSubmission(s) }
func (p Policy) CanResubmit(s domain.Submission) bool {
	return domain.IsResubmittable(s.Status) && !p.IsLate(s.DueAt, p.now())
}
func (p Policy) ShouldNotify(s domain.Submission) bool {
	return s.Status == domain.StatusSubmitted || s.Status == domain.StatusReturned
}
func (p Policy) NotificationKind(s domain.Submission) string {
	switch s.Status {
	case domain.StatusSubmitted:
		return "review_requested"
	case domain.StatusReturned:
		return "resubmission_requested"
	case domain.StatusApproved:
		return "certificate_ready"
	}
	return ""
}
func (p Policy) RetryLimit() int { return 5 }
func (p Policy) RetryDelay(attempt int) time.Duration {
	if attempt < 0 {
		attempt = 0
	}
	if attempt > 6 {
		attempt = 6
	}
	return time.Duration(1<<attempt) * time.Second
}
func (p Policy) ShouldRetry(attempt int, err error) bool {
	return err != nil && attempt < p.RetryLimit()
}
func (p Policy) IsPermanent(err error) bool {
	return err != nil && strings.Contains(err.Error(), "permanent")
}
func (p Policy) AuditAction(from, to domain.SubmissionStatus) string {
	return string(from) + "_to_" + string(to)
}
func (p Policy) QueryLimit(v int) int {
	if v <= 0 {
		return 20
	}
	if v > 100 {
		return 100
	}
	return v
}
func (p Policy) QueryOffset(v int) int {
	if v < 0 {
		return 0
	}
	return v
}
func (p Policy) SortKey(v string) string {
	switch v {
	case "created", "updated", "title":
		return v
	}
	return "created"
}
func (p Policy) SortDirection(v string) string {
	if strings.EqualFold(v, "asc") {
		return "asc"
	}
	return "desc"
}
func (p Policy) IsSafeSort(v string) bool { return p.SortKey(v) == v || v == "" }
func (p Policy) RoleCanView(r domain.Role) bool {
	return domain.AllowedRole(r, domain.RoleStudent, domain.RoleTeacher, domain.RoleReviewer, domain.RoleAdmin)
}
func (p Policy) RoleCanExport(r domain.Role) bool {
	return domain.AllowedRole(r, domain.RoleTeacher, domain.RoleReviewer, domain.RoleAdmin)
}
func (p Policy) RoleCanManageUsers(r domain.Role) bool { return r == domain.RoleAdmin }
func (p Policy) RoleCanManageCourse(r domain.Role) bool {
	return domain.AllowedRole(r, domain.RoleTeacher, domain.RoleAdmin)
}
func (p Policy) RoleCanReview(r domain.Role) bool {
	return domain.AllowedRole(r, domain.RoleReviewer, domain.RoleTeacher, domain.RoleAdmin)
}
func (p Policy) RoleCanSubmit(r domain.Role) bool { return r == domain.RoleStudent }
func (p Policy) ContextDeadline(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
func (p Policy) EnsureContext(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return nil
}
func (p Policy) TxOptions() sql.TxOptions { return sql.TxOptions{Isolation: sql.LevelSerializable} }
func (p Policy) Metadata(actor int64, request string) map[string]any {
	return map[string]any{"actor_id": actor, "request_id": request, "at": p.now().Format(time.RFC3339Nano)}
}
