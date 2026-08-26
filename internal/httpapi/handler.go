package httpapi

import (
	"encoding/json"
	"github.com/11DingKing/ai-course-cert-go/internal/apperr"
	"github.com/11DingKing/ai-course-cert-go/internal/auth"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/11DingKing/ai-course-cert-go/internal/middleware"
	"github.com/11DingKing/ai-course-cert-go/internal/service"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	Svc      service.Service
	Courses  service.CourseService
	Sessions *auth.Manager
}

func New(s service.Service, c service.CourseService, m *auth.Manager) *Handler {
	return &Handler{Svc: s, Courses: c, Sessions: m}
}
func write(w http.ResponseWriter, v any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var in struct{ Email, Password string }
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		write(w, map[string]string{"code": "invalid"}, 400)
		return
	}
	u, e := h.Svc.Login(r.Context(), in.Email, in.Password)
	if e != nil {
		write(w, map[string]string{"code": "unauthorized", "message": "invalid credentials"}, 401)
		return
	}
	t, exp, e := h.Sessions.Create(u)
	if e != nil {
		write(w, map[string]string{"code": "internal"}, 500)
		return
	}
	write(w, map[string]any{"token": t, "expires_at": exp}, 200)
}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get("Authorization")
	if len(t) > 7 {
		h.Sessions.Revoke(t[7:])
	}
	write(w, map[string]string{"status": "revoked"}, 200)
}
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	write(w, map[string]string{"status": "ok"}, 200)
}
func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	if h.Svc.Users.DB.PingContext(r.Context()) != nil {
		write(w, map[string]string{"status": "not_ready"}, 503)
		return
	}
	write(w, map[string]string{"status": "ready"}, 200)
}
func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.User(r)
	var in struct {
		Code, Title, Semester string
		OpensAt, ClosesAt     string
		Capacity              int
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		write(w, map[string]string{"code": "invalid"}, 400)
		return
	}
	o, _ := time.Parse(time.RFC3339, in.OpensAt)
	c, _ := time.Parse(time.RFC3339, in.ClosesAt)
	x, e := h.Courses.Create(r.Context(), u, domain.Course{Code: in.Code, Title: in.Title, Semester: in.Semester, OpensAt: o, ClosesAt: c, Capacity: in.Capacity})
	if e != nil {
		write(w, map[string]string{"code": string(codeOf(e)), "message": e.Error()}, 400)
		return
	}
	write(w, x, 201)
}
func codeOf(e error) apperr.Code {
	if a, ok := e.(*apperr.Error); ok {
		return a.Code
	}
	return apperr.Internal
}
func (h *Handler) CreateSubmission(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.User(r)
	cid, _ := strconv.ParseInt(r.URL.Query().Get("course_id"), 10, 64)
	var in struct{ Title, AIStatement string }
	json.NewDecoder(r.Body).Decode(&in)
	x, e := h.Svc.Submissions.Create(r.Context(), domain.Submission{CourseID: cid, StudentID: u.ID, Title: in.Title, AIStatement: in.AIStatement})
	if e != nil {
		write(w, map[string]string{"code": string(codeOf(e)), "message": e.Error()}, 400)
		return
	}
	write(w, x, 201)
}
func (h *Handler) Submit(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.User(r)
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	v, _ := strconv.Atoi(r.URL.Query().Get("version"))
	x, e := h.Svc.Submit(r.Context(), u.ID, id, v, middleware.Request(r))
	if e != nil {
		write(w, map[string]string{"code": string(codeOf(e)), "message": e.Error()}, 409)
		return
	}
	write(w, x, 200)
}
