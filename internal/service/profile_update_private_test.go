package service

import ("testing"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func TestConcurrentProfileUpdatesPreserveVersion(t *testing.T) {
 profile := &repository.CompetencyProfile{}
 UpdateProfileConcurrently(profile, 32)
 if profile.Version() != 32 { t.Fatalf("并发更新后版本为 %d", profile.Version()) }
}
