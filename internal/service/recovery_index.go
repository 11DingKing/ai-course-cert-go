package service

import "github.com/11DingKing/ai-course-cert-go/internal/repository"

func RecoverEvidenceIndex(rows map[int64]string) *repository.RestoredEvidenceIndex {
 index := repository.RestoreEvidenceIndex()
 for id, checksum := range rows { index.Put(id, checksum) }
 return index
}
