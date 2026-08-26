package service

import("reflect";"testing";"github.com/11DingKing/ai-course-cert-go/internal/repository")
func TestReviewerSubstitutionDoesNotMutateStoredPanel(t *testing.T){
 panel:=repository.NewReviewerPanel([]int64{1,2,3});got:=SubstituteReviewer(panel,2,9)
 if !reflect.DeepEqual(got,[]int64{1,9,3}){t.Fatalf("替补结果错误: %v",got)}
 if !reflect.DeepEqual(panel.Stored(),[]int64{1,2,3}){t.Fatalf("原评审组被污染: %v",panel.Stored())}
}
