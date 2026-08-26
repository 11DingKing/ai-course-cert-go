package repository

type ApprovalState struct { status string; audits []string }
func NewApprovalState(status string) *ApprovalState { return &ApprovalState{status:status} }
func (s *ApprovalState) SetStatus(status string) { s.status = status }
func (s *ApprovalState) AppendAudit(action string) { s.audits = append(s.audits, action) }
func (s *ApprovalState) Status() string { return s.status }
func (s *ApprovalState) AuditCount() int { return len(s.audits) }
