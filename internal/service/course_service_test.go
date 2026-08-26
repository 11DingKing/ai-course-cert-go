package service

import (
	"context"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/11DingKing/ai-course-cert-go/internal/repository"
	"github.com/11DingKing/ai-course-cert-go/internal/storage"
	"path/filepath"
	"testing"
	"time"
)

func TestCourseRoleAndSchedule(t *testing.T) {
	db, _ := storage.Open(filepath.Join(t.TempDir(), "c.db"))
	storage.Migrate(context.Background(), db)
	s := CourseService{Courses: repository.Courses{DB: db}}
	x := domain.Course{Code: "X", Title: "T", Semester: "S", OpensAt: time.Now(), ClosesAt: time.Now().Add(time.Hour), Capacity: 2}
	if _, e := s.Create(context.Background(), domain.User{Role: domain.RoleStudent}, x); e == nil {
		t.Fatal("student created")
	}
	if _, e := s.Create(context.Background(), domain.User{Role: domain.RoleTeacher}, x); e != nil {
		t.Fatal(e)
	}
	x.Code = "Y"
	x.ClosesAt = x.OpensAt.Add(-time.Hour)
	if _, e := s.Create(context.Background(), domain.User{Role: domain.RoleAdmin}, x); e == nil {
		t.Fatal("bad schedule")
	}
}
