package storage

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	_ "modernc.org/sqlite"
	"strings"
)

//go:embed migrations/*.sql
var migrationFS embed.FS

func Open(path string) (*sql.DB, error) {
	db, e := sql.Open("sqlite", path)
	if e != nil {
		return nil, e
	}
	db.SetMaxOpenConns(8)
	if e = db.Ping(); e != nil {
		return nil, e
	}
	return db, nil
}
func Migrate(ctx context.Context, db *sql.DB) error {
	rows, e := migrationFS.ReadFile("migrations/001_init.sql")
	if e != nil {
		return e
	}
	for _, stmt := range strings.Split(string(rows), ";") {
		s := strings.TrimSpace(stmt)
		if s == "" {
			continue
		}
		if _, e = db.ExecContext(ctx, s); e != nil {
			return fmt.Errorf("migration: %w", e)
		}
	}
	return nil
}
