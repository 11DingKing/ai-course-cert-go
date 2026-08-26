package service

import ("context"; "errors"; "testing"; "time")

func TestCancelledReviewResumeStopsImmediately(t *testing.T) {
 ctx, cancel := context.WithCancel(context.Background())
 cancel()
 started := time.Now()
 err := ResumePendingReview(ctx, 120*time.Millisecond)
 if !errors.Is(err, context.Canceled) { t.Fatalf("取消后返回 %v，期望 context.Canceled", err) }
 if time.Since(started) > 50*time.Millisecond { t.Fatalf("取消未及时传播") }
}
