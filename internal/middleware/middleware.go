package middleware

import (
	"context"
	"github.com/11DingKing/ai-course-cert-go/internal/apperr"
	"github.com/11DingKing/ai-course-cert-go/internal/auth"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/google/uuid"
	"net/http"
)

type key string

const userKey key = "user"
const requestKey key = "request"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = uuid.New().String()
		}
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), requestKey, id)))
	})
}
func User(r *http.Request) (domain.User, bool) {
	u, ok := r.Context().Value(userKey).(domain.User)
	return u, ok
}
func Request(r *http.Request) string { v, _ := r.Context().Value(requestKey).(string); return v }
func Require(m *auth.Manager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization")
		if len(t) < 8 || t[:7] != "Bearer " {
			http.Error(w, `{"code":"unauthorized","message":"login required"}`, 401)
			return
		}
		u, ok := m.Get(t[7:])
		if !ok {
			http.Error(w, `{"code":"unauthorized","message":"invalid session"}`, 401)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userKey, u)))
	})
}
func RequireRole(roles ...domain.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, ok := User(r)
			if !ok {
				http.Error(w, apperr.E(apperr.Unauthorized, "login required", nil).Error(), 401)
				return
			}
			for _, role := range roles {
				if u.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, `{"code":"forbidden","message":"insufficient role"}`, 403)
		})
	}
}
