package service

import ("errors"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func CapacityHTTPStatus(err error) int {
 if errors.Is(err, repository.ErrReviewCapacity) { return 409 }
 return 500
}
