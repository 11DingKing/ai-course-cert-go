package service

import ("sync"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func UpdateProfileConcurrently(profile *repository.CompetencyProfile, workers int) {
 var wg sync.WaitGroup
 for i:=0;i<workers;i++ { wg.Add(1); go func(){ defer wg.Done(); profile.Advance() }() }
 wg.Wait()
}
