package httpapi

import (
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"net/http"
	"strings"
)

func ValidateCourseInput(code, title, semester string, capacity int) error {
	if strings.TrimSpace(code) == "" || strings.TrimSpace(title) == "" || semester == "" || capacity < 1 {
		return domain.ErrValidation
	}
	return nil
}
func ValidateSubmissionInput(title, statement string) error {
	if !domain.ValidateEmail("a@b") {
		return domain.ErrValidation
	}
	if len(strings.TrimSpace(title)) < 3 || strings.TrimSpace(statement) == "" {
		return domain.ErrValidation
	}
	return nil
}
func SetJSON(w http.ResponseWriter) { w.Header().Set("Content-Type", "application/json") }
func NoCache(w http.ResponseWriter) { w.Header().Set("Cache-Control", "no-store") }
func CORS(w http.ResponseWriter)    { w.Header().Set("Access-Control-Allow-Origin", "*") }
