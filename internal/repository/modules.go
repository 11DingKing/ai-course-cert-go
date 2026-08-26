package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

type Modules struct{ DB *sql.DB }

func (m Modules) List(ctx context.Context, course int64) ([]domain.Module, error) {
	rows, e := m.DB.QueryContext(ctx, `SELECT id,course_id,name,kind,weight,created_at FROM modules WHERE course_id=? ORDER BY weight,name`, course)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.Module{}
	for rows.Next() {
		var x domain.Module
		if e := rows.Scan(&x.ID, &x.CourseID, &x.Name, &x.Kind, &x.Weight, &x.CreatedAt); e != nil {
			return nil, e
		}
		out = append(out, x)
	}
	return out, rows.Err()
}
func (m Modules) UpdateWeight(ctx context.Context, id int64, weight int) error {
	_, e := m.DB.ExecContext(ctx, `UPDATE modules SET weight=? WHERE id=?`, weight, id)
	return e
}
func (m Modules) Delete(ctx context.Context, id int64) error {
	_, e := m.DB.ExecContext(ctx, `DELETE FROM modules WHERE id=?`, id)
	return e
}
func (m Modules) Touch(ctx context.Context, id int64) error {
	_, e := m.DB.ExecContext(ctx, `UPDATE modules SET created_at=? WHERE id=?`, time.Now().UTC(), id)
	return e
}
