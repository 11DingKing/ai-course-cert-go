package service

import("context"; "testing"; "github.com/11DingKing/ai-course-cert-go/internal/repository")
func TestPublishedReviewJobSurvivesRequestCancellation(t *testing.T){
 ctx,cancel:=context.WithCancel(context.Background());jobs:=&repository.DurableReviewJobs{}
 done:=SchedulePublishedReview(ctx,jobs);cancel();<-done
 if jobs.Count()!=1{t.Fatalf("请求取消后发布任务数量为 %d",jobs.Count())}
}
