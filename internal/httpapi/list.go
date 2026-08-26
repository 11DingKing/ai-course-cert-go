package httpapi

import (
	"github.com/11DingKing/ai-course-cert-go/internal/pagination"
	"github.com/11DingKing/ai-course-cert-go/internal/repository"
	"net/http"
	"strconv"
)

func ListCourses(w http.ResponseWriter, r *http.Request, c repository.Courses) {
	p := pagination.Normalize(parseInt(r, "limit"), parseInt(r, "offset"))
	xs, e := c.List(r.Context(), p.Limit, p.Offset)
	if e != nil {
		writeError(w, r, e)
		return
	}
	write(w, map[string]any{"items": xs, "limit": p.Limit, "offset": p.Offset}, 200)
}
func parseInt(r *http.Request, k string) int   { v, _ := strconv.Atoi(r.URL.Query().Get(k)); return v }
func ParseBool(r *http.Request, k string) bool { return r.URL.Query().Get(k) == "true" }
func Header(r *http.Request, k string) string  { return r.Header.Get(k) }
func MethodAllowed(r *http.Request, methods ...string) bool {
	for _, m := range methods {
		if r.Method == m {
			return true
		}
	}
	return false
}
