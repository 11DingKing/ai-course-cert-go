package service

import (
	"context"
	"time"

	"github.com/11DingKing/ai-course-cert-go/internal/repository"
)

// SchedulePublishedReview asynchronously enqueues a confirmed review-publish
// task into the outbox. The enqueue runs on a context detached from the
// request lifecycle: the teacher may close the page immediately after
// submitting, which cancels the request context, but the already-confirmed
// publish task must still be durably written so the student receives the
// notification. Propagating the request context would let cancellation
// silently drop the task, so the background work does not depend on it.
func SchedulePublishedReview(ctx context.Context, jobs *repository.DurableReviewJobs) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		time.Sleep(10 * time.Millisecond)
		// Detach from the request: the outbox write must survive request
		// cancellation so the student is still notified.
		bg, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = jobs.Save(bg, "publish-review")
	}()
	return done
}
