package storage

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"
)

func TestMigrateCreatesTables(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db.sqlite")
	db, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer db.Close()
	if e = Migrate(context.Background(), db); e != nil {
		t.Fatal(e)
	}
	var n int
	if e = db.QueryRow(`SELECT COUNT(*) FROM sqlite_master WHERE type='table'`).Scan(&n); e != nil || n < 9 {
		t.Fatalf("tables=%d err=%v", n, e)
	}
	db.Close()
	db, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer db.Close()
	if e = Migrate(context.Background(), db); e != nil {
		t.Fatal(e)
	}
	if _, e = db.Exec(`SELECT 1`); e != nil {
		t.Fatal(e)
	}
}
func TestMigrateContextCancel(t *testing.T) {
	db, _ := sql.Open("sqlite", ":memory:")
	defer db.Close()
	ctx, c := context.WithCancel(context.Background())
	c()
	if e := Migrate(ctx, db); e == nil {
		t.Fatal("expected cancellation")
	}
	_ = os.ErrClosed
}
