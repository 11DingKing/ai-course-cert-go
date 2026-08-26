package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/ai-course-cert-go/internal/apperr"
	"github.com/11DingKing/ai-course-cert-go/internal/domain"
	"time"
)

type Users struct{ DB *sql.DB }

func (u Users) Create(ctx context.Context, email, name string, role domain.Role, pw string) (domain.User, error) {
	r, e := u.DB.ExecContext(ctx, `INSERT INTO users(email,name,role,password_hash,created_at) VALUES(?,?,?,?,?)`, email, name, role, pw, time.Now().UTC().Format(time.RFC3339Nano))
	if e != nil {
		return domain.User{}, apperr.E(apperr.Conflict, "email exists", e)
	}
	id, _ := r.LastInsertId()
	return domain.User{ID: id, Email: email, Name: name, Role: role, PasswordHash: pw}, nil
}
func (u Users) ByEmail(ctx context.Context, email string) (domain.User, error) {
	var x domain.User
	var role string
	var created string
	e := u.DB.QueryRowContext(ctx, `SELECT id,email,name,role,password_hash,created_at FROM users WHERE email=?`, email).Scan(&x.ID, &x.Email, &x.Name, &role, &x.PasswordHash, &created)
	if e == sql.ErrNoRows {
		return x, apperr.E(apperr.NotFound, "user", e)
	}
	if e != nil {
		return x, e
	}
	x.Role = domain.Role(role)
	x.CreatedAt, _ = time.Parse(time.RFC3339Nano, created)
	return x, nil
}
