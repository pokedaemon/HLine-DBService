package storage

import (
	"fmt"
	"strings"
	"testing"
)

// TestStorage open and ping test db and return teardown function that truncate tables
func TestStorage(t *testing.T, dbURL string) (*Storage, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = dbURL

	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("truncate %s cascade", strings.Join(tables, ""))); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}
}
