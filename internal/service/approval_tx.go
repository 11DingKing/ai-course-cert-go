package service

import ("errors"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func ApproveWithAudit(state *repository.ApprovalState, failAudit bool) error {
 state.SetStatus("approved")
 if failAudit { return errors.New("audit storage unavailable") }
 state.AppendAudit("approved")
 return nil
}
