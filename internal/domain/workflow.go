package domain

import "time"

type WorkflowEvent struct {
	At       time.Time
	From, To SubmissionStatus
	ActorID  int64
	Reason   string
}

func NewWorkflowEvent(from, to SubmissionStatus, actor int64, reason string) WorkflowEvent {
	return WorkflowEvent{At: time.Now().UTC(), From: from, To: to, ActorID: actor, Reason: reason}
}
func ValidateWorkflow(events []WorkflowEvent) bool {
	if len(events) == 0 {
		return true
	}
	for i, e := range events {
		if !e.From.CanTransition(e.To) {
			return false
		}
		if i > 0 && events[i-1].To != e.From {
			return false
		}
	}
	return true
}
func TransitionPath(from SubmissionStatus) []SubmissionStatus {
	switch from {
	case StatusDraft:
		return []SubmissionStatus{StatusSubmitted}
	case StatusSubmitted:
		return []SubmissionStatus{StatusReturned, StatusApproved}
	case StatusReturned:
		return []SubmissionStatus{StatusSubmitted}
	case StatusApproved:
		return []SubmissionStatus{StatusArchived}
	default:
		return nil
	}
}
func StatusLabel(s SubmissionStatus) string {
	switch s {
	case StatusDraft:
		return "草稿"
	case StatusSubmitted:
		return "已提交"
	case StatusReturned:
		return "退回"
	case StatusApproved:
		return "通过"
	case StatusArchived:
		return "归档"
	}
	return "未知"
}
func IsReviewable(s SubmissionStatus) bool    { return s == StatusSubmitted }
func IsResubmittable(s SubmissionStatus) bool { return s == StatusDraft || s == StatusReturned }
func NeedsEvidence(s SubmissionStatus) bool   { return s == StatusSubmitted || s == StatusApproved }
func WindowRemaining(c Course, now time.Time) time.Duration {
	if now.Before(c.OpensAt) {
		return c.OpensAt.Sub(now)
	}
	if now.After(c.ClosesAt) {
		return 0
	}
	return c.ClosesAt.Sub(now)
}
func CapacityAvailable(used, capacity int) bool { return used < capacity && capacity > 0 }
func NormalizeSemester(s string) string {
	if len(s) > 20 {
		return s[:20]
	}
	return s
}
func NormalizeTitle(s string) string {
	for len(s) > 0 && (s[0] == ' ' || s[0] == '\n') {
		s = s[1:]
	}
	return s
}
func EvidenceVersions(xs []Evidence) bool {
	for i, x := range xs {
		if x.Version != i+1 {
			return false
		}
	}
	return true
}
func ReviewMatchesVersion(r Review, s Submission) bool {
	return r.SubmissionID == s.ID && r.Version <= s.Version
}
func CanArchive(s Submission, rs []Review) bool {
	if s.Status != StatusApproved {
		return false
	}
	for _, r := range rs {
		if r.SubmissionID == s.ID && r.Decision == "approve" {
			return true
		}
	}
	return false
}
