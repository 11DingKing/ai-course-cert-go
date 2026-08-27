package service

import ("testing"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func TestEvidenceRetryDoesNotCreateDuplicateVersion(t *testing.T) {
 store := &repository.RetryEvidenceStore{}
 if err := PersistEvidenceWithRetry(store); err != nil { t.Fatalf("重试未恢复: %v", err) }
 if store.Versions() != 1 { t.Fatalf("一次提交产生 %d 个证据版本", store.Versions()) }
}
