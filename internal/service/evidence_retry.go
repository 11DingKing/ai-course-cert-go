package service

import "github.com/11DingKing/ai-course-cert-go/internal/repository"

func PersistEvidenceWithRetry(store *repository.RetryEvidenceStore) error {
 for attempt:=0; attempt<2; attempt++ {
  if err:=store.CreateVersion(); err==nil { return nil }
 }
 return store.LastError()
}
