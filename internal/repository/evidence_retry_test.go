package repository

import "testing"

func TestRetryEvidenceStoreIdempotentAfterTransientFailure(t *testing.T) {
	var s RetryEvidenceStore
	// First attempt: transient outbox fault. No version must be persisted.
	if e := s.CreateVersion(); e == nil {
		t.Fatal("first attempt should fail with transient outbox error")
	}
	if v := s.Versions(); v != 0 {
		t.Fatalf("transient failure persisted version %d, want 0", v)
	}
	// Retry succeeds; exactly one version is produced for the whole operation.
	if e := s.CreateVersion(); e != nil {
		t.Fatalf("retry should succeed: %v", e)
	}
	if v := s.Versions(); v != 1 {
		t.Fatalf("after retry want 1 version, got %d", v)
	}
}

func TestRetryEvidenceStoreOneVersionAfterRetry(t *testing.T) {
	var s RetryEvidenceStore
	// One upload: first write fails transiently, retry succeeds. The whole
	// operation must produce exactly one persisted version.
	for attempt := 0; attempt < 2; attempt++ {
		if err := s.CreateVersion(); err == nil {
			break
		}
	}
	if v := s.Versions(); v != 1 {
		t.Fatalf("one upload must produce exactly one version, got %d", v)
	}
}
