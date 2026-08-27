package repository

import("context"; "sync")
type DurableReviewJobs struct { mu sync.Mutex; jobs []string }
func(j *DurableReviewJobs) Save(ctx context.Context,name string) error { if err:=ctx.Err();err!=nil{return err};j.mu.Lock();defer j.mu.Unlock();j.jobs=append(j.jobs,name);return nil }
func(j *DurableReviewJobs) Count()int{j.mu.Lock();defer j.mu.Unlock();return len(j.jobs)}
