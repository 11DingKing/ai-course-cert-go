package repository

import "errors"

type ReviewSlotPool struct { used, limit int }
func NewReviewSlotPool(limit int) *ReviewSlotPool { return &ReviewSlotPool{limit:limit} }
func (p *ReviewSlotPool) Acquire() error { if p.used >= p.limit { return errors.New("review slots exhausted") }; p.used++; return nil }
func (p *ReviewSlotPool) Release() { p.used-- }
func (p *ReviewSlotPool) Used() int { return p.used }
