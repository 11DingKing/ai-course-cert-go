package service

import (
	"context"
	"fmt"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"github.com/11DingKing/ai-course-cert-go/internal/repository"
)

type BatchService struct{ Submissions repository.Submissions }
type BatchResult struct {
	ID  int64
	OK  bool
	Err error
}

func (b BatchService) SubmitMany(ctx context.Context, actor int64, ids []int64) []BatchResult {
	out := make([]BatchResult, 0, len(ids))
	for _, id := range ids {
		x, e := b.Submissions.Get(ctx, id)
		if e != nil {
			out = append(out, BatchResult{ID: id, Err: e})
			continue
		}
		if x.StudentID != actor {
			out = append(out, BatchResult{ID: id, Err: fmt.Errorf("owner mismatch")})
			continue
		}
		_, e = b.Submissions.Transition(ctx, id, x.Status, domain.StatusSubmitted, x.Version)
		out = append(out, BatchResult{ID: id, OK: e == nil, Err: e})
	}
	return out
}
func (b BatchService) ArchiveApproved(ctx context.Context, ids []int64) int {
	n := 0
	for _, id := range ids {
		x, e := b.Submissions.Get(ctx, id)
		if e == nil && x.Status == domain.StatusApproved {
			if _, e = b.Submissions.Transition(ctx, id, x.Status, domain.StatusArchived, x.Version); e == nil {
				n++
			}
		}
	}
	return n
}
