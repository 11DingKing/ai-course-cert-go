package repository

import "errors"

type RetryEvidenceStore struct { versions, calls int; last error }
func (s *RetryEvidenceStore) CreateVersion() error { s.calls++; s.versions++; if s.calls==1 { s.last=errors.New("transient outbox failure"); return s.last }; s.last=nil; return nil }
func (s *RetryEvidenceStore) Versions() int { return s.versions }
func (s *RetryEvidenceStore) LastError() error { return s.last }
