package service

import ("context"; "time"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func SchedulePublishedReview(ctx context.Context, jobs *repository.DurableReviewJobs) <-chan struct{} {
 done:=make(chan struct{})
 go func(){ defer close(done); time.Sleep(10*time.Millisecond); jobs.Save(ctx, "publish-review") }()
 return done
}
