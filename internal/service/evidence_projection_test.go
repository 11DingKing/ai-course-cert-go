package service

import (
	"testing"

	"github.com/11DingKing/ai-course-cert-go/internal/repository"
)

func TestMaskEvidenceForStudentDoesNotPolluteStored(t *testing.T) {
	original := []string{"https://repo/audio/a1.wav", "https://repo/docs/d1.pdf"}
	store := repository.NewEvidenceSnapshot(original)

	masked := MaskEvidenceForStudent(store)
	for i, v := range original {
		if masked[i] != "masked:"+v {
			t.Fatalf("masked[%d] = %q, want %q", i, masked[i], "masked:"+v)
		}
	}

	stored := store.Stored()
	for i, v := range original {
		if stored[i] != v {
			t.Fatalf("stored[%d] = %q, want original %q (masking must not pollute repository)", i, stored[i], v)
		}
	}
}
