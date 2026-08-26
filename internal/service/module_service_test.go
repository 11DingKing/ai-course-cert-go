package service

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/11DingKing/ai-course-cert-go/internal/storage"
	"path/filepath"
	"testing"
)

func TestModuleService(t *testing.T) {
	db, _ := storage.Open(filepath.Join(t.TempDir(), "m.db"))
	storage.Migrate(context.Background(), db)
	m := ModuleService{db}
	teacher := domain.User{ID: 1, Role: domain.RoleTeacher}
	if _, e := m.Add(context.Background(), domain.User{Role: domain.RoleStudent}, 1, "x", "writing", 10); e == nil {
		t.Fatal("forbidden")
	}
	if _, e := m.Add(context.Background(), teacher, 1, "x", "writing", 10); e != nil {
		t.Fatal(e)
	}
	if _, e := m.Add(context.Background(), teacher, 1, "x", "writing", 10); e == nil {
		t.Fatal("duplicate")
	}
	var n int
	db.QueryRow(`SELECT COUNT(*) FROM modules`).Scan(&n)
	if n != 1 {
		t.Fatal(n)
	}
	_ = sql.ErrNoRows
}
