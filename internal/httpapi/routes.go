package httpapi

import (
	"github.com/11DingKing/ai-course-cert-go/internal/middleware"
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/healthz", h.Health)
	m.HandleFunc("/readyz", h.Ready)
	m.HandleFunc("/v1/login", h.Login)
	m.Handle("/v1/logout", middleware.Require(h.Sessions, http.HandlerFunc(h.Logout)))
	m.Handle("/v1/courses", middleware.Require(h.Sessions, middleware.RequireRole("admin", "teacher")(http.HandlerFunc(h.CreateCourse))))
	m.Handle("/v1/submissions", middleware.Require(h.Sessions, middleware.RequireRole("student")(http.HandlerFunc(h.CreateSubmission))))
	m.Handle("/v1/submissions/submit", middleware.Require(h.Sessions, middleware.RequireRole("student")(http.HandlerFunc(h.Submit))))
	return middleware.RequestID(m)
}
