package service

import "github.com/11DingKing/ai-course-cert-go/internal/repository"

func ProcessReviewBatch(pool *repository.ReviewSlotPool, ids []int64) error {
	for range ids {
		if err := pool.Acquire(); err != nil {
			return err
		}
		func() {
			defer pool.Release()
			// review work for this material runs here
		}()
	}
	return nil
}
