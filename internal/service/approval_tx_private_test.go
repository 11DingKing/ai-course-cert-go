package service

import ("testing"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func TestApprovalAuditFailureRollsBackStatus(t *testing.T) {
 state := repository.NewApprovalState("submitted")
 if err := ApproveWithAudit(state, true); err == nil { t.Fatal("审计失败未返回错误") }
 if state.Status() != "submitted" { t.Fatalf("失败后状态残留为 %s", state.Status()) }
 if state.AuditCount() != 0 { t.Fatalf("失败后产生了审计记录") }
}
