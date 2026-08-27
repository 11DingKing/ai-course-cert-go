package httpapi

import (
	"encoding/json"
	"github.com/11DingKing/ai-course-cert-go/internal/apperr"
	"net/http"
)

type ErrorBody struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func writeError(w http.ResponseWriter, r *http.Request, e error) {
	status := apperr.Status(e)
	write(w, ErrorBody{Code: string(codeOf(e)), Message: apperr.Message(e), RequestID: r.Header.Get("X-Request-ID")}, status)
}
func decode(r *http.Request, v any) error {
	if r.Body == nil {
		return &json.InvalidUnmarshalError{Type: nil}
	}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	return d.Decode(v)
}
