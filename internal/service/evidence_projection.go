package service

import "github.com/11DingKing/ai-course-cert-go/internal/repository"

func MaskEvidenceForStudent(store *repository.EvidenceSnapshot) []string {
 values := store.Values()
 for i := range values { values[i] = "masked:" + values[i] }
 return values
}
