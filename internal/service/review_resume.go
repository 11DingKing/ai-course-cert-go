package service

import (
 "context"
 "time"
 "github.com/11DingKing/ai-course-cert-go/internal/repository"
)

func ResumePendingReview(ctx context.Context, wait time.Duration) error {
 return repository.WaitForReview(context.Background(), wait)
}
