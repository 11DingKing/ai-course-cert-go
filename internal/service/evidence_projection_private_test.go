package service

import ("reflect"; "testing"; "github.com/11DingKing/ai-course-cert-go/internal/repository")

func TestEvidenceProjectionDoesNotMutateRepositorySnapshot(t *testing.T) {
 store := repository.NewEvidenceSnapshot([]string{"oral.wav", "essay.pdf"})
 got := MaskEvidenceForStudent(store)
 if !reflect.DeepEqual(got, []string{"masked:oral.wav", "masked:essay.pdf"}) { t.Fatalf("投影结果错误: %v", got) }
 if !reflect.DeepEqual(store.Stored(), []string{"oral.wav", "essay.pdf"}) { t.Fatalf("仓库存储被投影污染: %v", store.Stored()) }
}
