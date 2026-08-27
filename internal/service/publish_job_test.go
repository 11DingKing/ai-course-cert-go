package service

import (
	"context"
	"testing"
	"time"

	"github.com/11DingKing/ai-course-cert-go/internal/repository"
)

// SchedulePublishedReview must enqueue the confirmed publish task even when
// the teacher closes the page immediately, which cancels the request context.
// If the background save were tied to the request context, the task would be
// silently dropped and the student would never receive the publish notification.
func TestSchedulePublishedReviewSurvivesRequestCancellation(t *testing.T) {
	jobs := &repository.DurableReviewJobs{}

	// Cancel the request context before scheduling, simulating the teacher
	// closing the page the instant after the publish request is accepted.
	req, cancel := context.WithCancel(context.Background())
	cancel()

	done := SchedulePublishedReview(req, jobs)
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("schedule did not complete")
	}

	if n := jobs.Count(); n != 1 {
		t.Fatalf("expected 1 durable publish job after request cancellation, got %d", n)
	}
}
