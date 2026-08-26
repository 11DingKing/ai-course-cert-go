package repository

type ReviewerPanel struct{members []int64}
func NewReviewerPanel(ids []int64)*ReviewerPanel{return &ReviewerPanel{members:append([]int64(nil),ids...)}}
func(p *ReviewerPanel)Members()[]int64{return p.members}
func(p *ReviewerPanel)Stored()[]int64{return append([]int64(nil),p.members...)}
