package postgres_test

import (
	"os"
	"testing"
)

var (
	testDatabaseURL string
)

func TestMain(m *testing.M) {
	testDatabaseURL = os.Getenv("DATABASE_URL")
	if testDatabaseURL == "" {
		testDatabaseURL = "host=localhost dbname=reverseshoptest sslmode=disable"
	}

	os.Exit(m.Run())
}
