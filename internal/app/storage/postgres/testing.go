package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// TestStorage open and ping test db and return teardown function that truncate tables
func TestStorage(t *testing.T, dbURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatal(err)
	}

	if db.Ping(); err != nil {
		t.Fatal(err)
	}

	// truncate just dynamic tables!!!
	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("truncate %s cascade", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		db.Close()
	}
}
