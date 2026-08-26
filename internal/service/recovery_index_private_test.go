package service

import "testing"

func TestRestartRecoveryInitializesEvidenceIndex(t *testing.T) {
 defer func() { if recovered := recover(); recovered != nil { t.Fatalf("重启恢复发生 panic: %v", recovered) } }()
 index := RecoverEvidenceIndex(map[int64]string{7:"abc12345"})
 if got := index.Get(7); got != "abc12345" { t.Fatalf("恢复后索引为 %q", got) }
}
