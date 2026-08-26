package service

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/apperr"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

type ModuleService struct{ DB *sql.DB }

func (m ModuleService) Add(ctx context.Context, actor domain.User, course int64, name, kind string, weight int) (domain.Module, error) {
	if actor.Role != domain.RoleTeacher && actor.Role != domain.RoleAdmin {
		return domain.Module{}, apperr.E(apperr.Forbidden, "role cannot edit modules", nil)
	}
	if name == "" || weight <= 0 {
		return domain.Module{}, apperr.E(apperr.Invalid, "module fields", nil)
	}
	r, e := m.DB.ExecContext(ctx, `INSERT INTO modules(course_id,name,kind,weight,created_at) VALUES(?,?,?,?,?)`, course, name, kind, weight, time.Now().UTC())
	if e != nil {
		return domain.Module{}, apperr.E(apperr.Conflict, "module exists", e)
	}
	id, _ := r.LastInsertId()
	return domain.Module{ID: id, CourseID: course, Name: name, Kind: kind, Weight: weight}, nil
}
