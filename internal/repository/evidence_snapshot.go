package repository

type EvidenceSnapshot struct { values []string }
func NewEvidenceSnapshot(values []string) *EvidenceSnapshot {
 return &EvidenceSnapshot{values: append([]string(nil), values...)}
}
func (s *EvidenceSnapshot) Values() []string { return s.values }
func (s *EvidenceSnapshot) Stored() []string { return append([]string(nil), s.values...) }
