package repository

import "errors"

// RetryEvidenceStore simulates an evidence store whose first write fails with
// a transient outbox fault and succeeds on retry. A failed write must not leave
// a persisted version behind, so retries stay idempotent: one business
// operation produces at most one version.
type RetryEvidenceStore struct {
	versions, calls int
	last            error
}

func (s *RetryEvidenceStore) CreateVersion() error {
	s.calls++
	if s.calls == 1 {
		// Transient outbox failure: nothing committed, no version persisted.
		s.last = errors.New("transient outbox failure")
		return s.last
	}
	s.versions++
	s.last = nil
	return nil
}

func (s *RetryEvidenceStore) Versions() int    { return s.versions }
func (s *RetryEvidenceStore) LastError() error { return s.last }
func (s *RetryEvidenceStore) Attempts() int    { return s.calls }
