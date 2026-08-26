package domain

import "time"

type Role string

const (
	RoleStudent  Role = "student"
	RoleTeacher  Role = "teacher"
	RoleReviewer Role = "reviewer"
	RoleAdmin    Role = "admin"
)

type SubmissionStatus string

const (
	StatusDraft     SubmissionStatus = "draft"
	StatusSubmitted SubmissionStatus = "submitted"
	StatusReturned  SubmissionStatus = "returned"
	StatusApproved  SubmissionStatus = "approved"
	StatusArchived  SubmissionStatus = "archived"
)

type User struct {
	ID           int64
	Email, Name  string
	Role         Role
	PasswordHash string
	CreatedAt    time.Time
	RevokedAt    *time.Time
}
type Course struct {
	ID                int64
	Code, Title       string
	Semester          string
	OpensAt, ClosesAt time.Time
	Capacity          int
	CreatedAt         time.Time
}
type Module struct {
	ID, CourseID int64
	Name         string
	Kind         string
	Weight       int
	CreatedAt    time.Time
}
type Submission struct {
	ID, CourseID, StudentID int64
	Status                  SubmissionStatus
	Title, AIStatement      string
	DueAt                   time.Time
	Version                 int
	SubmittedAt             *time.Time
	CreatedAt, UpdatedAt    time.Time
}
type Evidence struct {
	ID, SubmissionID           int64
	Version                    int
	URI, Checksum, Description string
	CreatedAt                  time.Time
}
type Review struct {
	ID, SubmissionID, ReviewerID int64
	Decision                     string
	Comment                      string
	Version                      int
	CreatedAt                    time.Time
}
type AuditEvent struct {
	ID                        int64
	ActorID                   int64
	EntityType                string
	EntityID                  int64
	Action, Result, RequestID string
	Metadata                  string
	CreatedAt                 time.Time
}
type IdempotencyRecord struct {
	Key, Operation string
	UserID         int64
	ResultID       int64
	CreatedAt      time.Time
}
type OutboxJob struct {
	ID                    int64
	Kind, Payload, Status string
	Attempts              int
	RunAfter              time.Time
	LastError             string
	CreatedAt, UpdatedAt  time.Time
}

func (s SubmissionStatus) CanTransition(to SubmissionStatus) bool {
	switch s {
	case StatusDraft:
		return to == StatusSubmitted
	case StatusSubmitted:
		return to == StatusReturned || to == StatusApproved
	case StatusReturned:
		return to == StatusSubmitted
	case StatusApproved:
		return to == StatusArchived
	default:
		return false
	}
}
