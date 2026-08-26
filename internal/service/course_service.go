package service

import (
	"context"
	"github.com/11DingKing/ai-course-cert-go/internal/apperr"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/11DingKing/ai-course-cert-go/internal/repository"
	"time"
)

type CourseService struct{ Courses repository.Courses }

func (c CourseService) Create(ctx context.Context, actor domain.User, x domain.Course) (domain.Course, error) {
	if actor.Role != domain.RoleAdmin && actor.Role != domain.RoleTeacher {
		return x, apperr.E(apperr.Forbidden, "role cannot create course", nil)
	}
	if x.ClosesAt.Before(x.OpensAt) || x.Capacity <= 0 {
		return x, apperr.E(apperr.Invalid, "invalid schedule", nil)
	}
	x.CreatedAt = time.Now().UTC()
	return c.Courses.Create(ctx, x)
}
