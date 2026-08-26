package service

import ("testing"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func TestCapacityConflictKeepsStableHTTPMapping(t *testing.T) {
 err := repository.CapacityFailure(42)
 if got := CapacityHTTPStatus(err); got != 409 {
  t.Fatalf("评审名额冲突映射为 %d，期望 409", got)
 }
}
