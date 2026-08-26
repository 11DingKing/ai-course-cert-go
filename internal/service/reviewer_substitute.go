package service

import "github.com/11DingKing/ai-course-cert-go/internal/repository"

func SubstituteReviewer(panel *repository.ReviewerPanel, oldID,newID int64) []int64 {
 members:=panel.Members()
 for i,id:=range members{if id==oldID{members[i]=newID}}
 return members
}
