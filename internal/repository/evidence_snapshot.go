package repository

type EvidenceSnapshot struct{ values []string }

func NewEvidenceSnapshot(values []string) *EvidenceSnapshot {
	return &EvidenceSnapshot{values: append([]string(nil), values...)}
}

// Values returns a defensive copy so callers cannot mutate the stored evidence.
func (s *EvidenceSnapshot) Values() []string { return append([]string(nil), s.values...) }

// Stored returns a defensive copy of the persisted evidence URIs.
func (s *EvidenceSnapshot) Stored() []string { return append([]string(nil), s.values...) }
