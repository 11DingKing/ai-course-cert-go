package auth

import (
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"testing"
	"time"
)

func TestSessionLifecycle(t *testing.T) {
	m := New(time.Hour)
	u := domain.User{ID: 1, Role: domain.RoleStudent}
	tok, exp, e := m.Create(u)
	if e != nil || tok == "" || exp.Before(time.Now()) {
		t.Fatal(e)
	}
	got, ok := m.Get(tok)
	if !ok || got.ID != 1 {
		t.Fatal("missing")
	}
	m.Revoke(tok)
	if _, ok = m.Get(tok); ok {
		t.Fatal("revoked")
	}
}
func TestSessionExpires(t *testing.T) {
	m := New(time.Millisecond)
	tok, _, _ := m.Create(domain.User{ID: 2})
	time.Sleep(3 * time.Millisecond)
	if _, ok := m.Get(tok); ok {
		t.Fatal("expired token accepted")
	}
}
