package middleware

import (
	"github.com/11DingKing/ai-course-cert-go/internal/auth"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequireRole(t *testing.T) {
	m := auth.New(time.Hour)
	tok, _, _ := m.Create(domain.User{ID: 1, Role: domain.RoleStudent})
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(204) })
	h := Require(m, RequireRole(domain.RoleTeacher)(next))
	q := httptest.NewRequest(http.MethodGet, "/", nil)
	q.Header.Set("Authorization", "Bearer "+tok)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, q)
	if w.Code != 403 {
		t.Fatal(w.Code)
	}
}
func TestRequestIDPreserved(t *testing.T) {
	h := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Request(r) != "fixed" {
			t.Fatal(Request(r))
		}
	}))
	q := httptest.NewRequest(http.MethodGet, "/", nil)
	q.Header.Set("X-Request-ID", "fixed")
	h.ServeHTTP(httptest.NewRecorder(), q)
}
