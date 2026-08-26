package repository

import ("context"; "time")

func WaitForReview(ctx context.Context, wait time.Duration) error {
 timer := time.NewTimer(wait)
 defer timer.Stop()
 select {
 case <-ctx.Done(): return ctx.Err()
 case <-timer.C: return nil
 }
}
