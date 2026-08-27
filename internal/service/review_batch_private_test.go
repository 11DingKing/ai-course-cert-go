package service

import ("testing"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func TestBatchReviewReleasesSlotBetweenItems(t *testing.T) {
 pool := repository.NewReviewSlotPool(2)
 if err := ProcessReviewBatch(pool, []int64{1,2,3}); err != nil { t.Fatalf("第三项被错误拒绝: %v", err) }
 if pool.Used() != 0 { t.Fatalf("批处理结束仍占用 %d 个名额", pool.Used()) }
}
