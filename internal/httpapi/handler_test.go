package httpapi

import (
	"context"
	"encoding/json"
	"github.com/11DingKing/ai-course-cert-go/internal/audit"
	"github.com/11DingKing/ai-course-cert-go/internal/auth"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/11DingKing/ai-course-cert-go/internal/repository"
	"github.com/11DingKing/ai-course-cert-go/internal/service"
	"github.com/11DingKing/ai-course-cert-go/internal/storage"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func httpFixture(t *testing.T) (*Handler, string) {
	db, _ := storage.Open(filepath.Join(t.TempDir(), "h.db"))
	storage.Migrate(context.Background(), db)
	u := repository.Users{DB: db}
	student, _ := u.Create(context.Background(), "student@h", "Student", domain.RoleStudent, "pw")
	svc := service.Service{Users: u, Courses: repository.Courses{DB: db}, Submissions: repository.Submissions{DB: db}, Evidences: repository.Evidences{DB: db}, Reviews: repository.Reviews{DB: db}, Audit: audit.Logger{DB: db}}
	h := New(svc, service.CourseService{Courses: repository.Courses{DB: db}}, auth.New(time.Hour))
	tok, _, _ := h.Sessions.Create(student)
	return h, tok
}
func TestHealthReady(t *testing.T) {
	h, _ := httpFixture(t)
	for _, p := range []string{"/healthz", "/readyz"} {
		r := httptest.NewRequest(http.MethodGet, p, nil)
		w := httptest.NewRecorder()
		h.Routes().ServeHTTP(w, r)
		if w.Code != 200 {
			t.Fatalf("%s %d", p, w.Code)
		}
	}
}
func TestLoginLogout(t *testing.T) {
	h, _ := httpFixture(t)
	body := `{"email":"student@h","password":"pw"}`
	w := httptest.NewRecorder()
	h.Routes().ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(body)))
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
	var out struct{ Token string }
	json.Unmarshal(w.Body.Bytes(), &out)
	if out.Token == "" {
		t.Fatal("no token")
	}
	q := httptest.NewRequest(http.MethodPost, "/v1/logout", nil)
	q.Header.Set("Authorization", "Bearer "+out.Token)
	w = httptest.NewRecorder()
	h.Routes().ServeHTTP(w, q)
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
func TestUnauthorizedAndRequestID(t *testing.T) {
	h, _ := httpFixture(t)
	q := httptest.NewRequest(http.MethodPost, "/v1/submissions?course_id=1", strings.NewReader(`{}`))
	w := httptest.NewRecorder()
	h.Routes().ServeHTTP(w, q)
	if w.Code != 401 {
		t.Fatal(w.Code)
	}
	if w.Header().Get("X-Request-ID") == "" {
		t.Fatal("request id")
	}
}
