package service

import (
	"context"
	"errors"
	"github.com/11DingKing/ai-course-cert-go/internal/apperr"
	"github.com/11DingKing/ai-course-cert-go/internal/audit"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/11DingKing/ai-course-cert-go/internal/repository"
	"time"
)

type Service struct {
	Users       repository.Users
	Courses     repository.Courses
	Submissions repository.Submissions
	Evidences   repository.Evidences
	Reviews     repository.Reviews
	Audit       audit.Logger
}

func (s Service) Login(ctx context.Context, email, pw string) (domain.User, error) {
	u, e := s.Users.ByEmail(ctx, email)
	if e != nil {
		return u, e
	}
	if u.PasswordHash != pw {
		return u, apperr.E(apperr.Unauthorized, "invalid credentials", errors.New("password mismatch"))
	}
	return u, nil
}
func (s Service) Submit(ctx context.Context, actor int64, id int64, version int, req string) (domain.Submission, error) {
	x, e := s.Submissions.Get(ctx, id)
	if e != nil {
		return x, e
	}
	if x.StudentID != actor {
		return x, apperr.E(apperr.Forbidden, "not owner", nil)
	}
	if time.Now().UTC().After(x.DueAt) {
		return x, apperr.E(apperr.Deadline, "submission overdue", nil)
	}
	y, e := s.Submissions.Transition(ctx, id, x.Status, domain.StatusSubmitted, version)
	if e == nil {
		_ = s.Audit.Record(ctx, actor, "submission", id, "submit", "ok", req, map[string]any{"version": version})
	}
	return y, e
}
func (s Service) Review(ctx context.Context, actor int64, id int64, version int, decision, comment, req string) (domain.Review, error) {
	if decision != "approve" && decision != "return" {
		return domain.Review{}, apperr.E(apperr.Invalid, "unsupported decision", nil)
	}
	x, e := s.Submissions.Get(ctx, id)
	if e != nil {
		return domain.Review{}, e
	}
	to := domain.StatusReturned
	if decision == "approve" {
		to = domain.StatusApproved
	}
	if _, e = s.Submissions.Transition(ctx, id, x.Status, to, version); e != nil {
		return domain.Review{}, e
	}
	r, e := s.Reviews.Add(ctx, domain.Review{SubmissionID: id, ReviewerID: actor, Decision: decision, Comment: comment, Version: version})
	if e == nil {
		_ = s.Audit.Record(ctx, actor, "submission", id, "review", "ok", req, map[string]any{"decision": decision})
	}
	return r, e
}
func (s Service) AddEvidence(ctx context.Context, actor int64, id int64, uri, checksum, description string) (domain.Evidence, error) {
	x, e := s.Submissions.Get(ctx, id)
	if e != nil {
		return domain.Evidence{}, e
	}
	if x.StudentID != actor {
		return domain.Evidence{}, apperr.E(apperr.Forbidden, "not owner", nil)
	}
	if uri == "" || checksum == "" {
		return domain.Evidence{}, apperr.E(apperr.Invalid, "evidence metadata required", nil)
	}
	return s.Evidences.Add(ctx, domain.Evidence{SubmissionID: id, URI: uri, Checksum: checksum, Description: description})
}
