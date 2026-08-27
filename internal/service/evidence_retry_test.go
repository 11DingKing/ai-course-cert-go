package service

import (
	"testing"

	"github.com/11DingKing/ai-course-cert-go/internal/repository"
)

// TestPersistEvidenceWithRetryOneVersion guards against the regression where a
// transient outbox failure on the first attempt persisted a version anyway, so
// the retry produced a second version and reviewers saw one upload twice.
func TestPersistEvidenceWithRetryOneVersion(t *testing.T) {
	var s repository.RetryEvidenceStore
	if e := PersistEvidenceWithRetry(&s); e != nil {
		t.Fatalf("retry should succeed after transient outbox failure: %v", e)
	}
	if v := s.Versions(); v != 1 {
		t.Fatalf("one upload must produce exactly one version, got %d", v)
	}
}
