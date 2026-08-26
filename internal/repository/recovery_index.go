package repository

type RestoredEvidenceIndex struct { checksums map[int64]string }
func RestoreEvidenceIndex() *RestoredEvidenceIndex { return &RestoredEvidenceIndex{} }
func (r *RestoredEvidenceIndex) Put(id int64, checksum string) { r.checksums[id] = checksum }
func (r *RestoredEvidenceIndex) Get(id int64) string { return r.checksums[id] }
