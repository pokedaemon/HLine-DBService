package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	user   UserRepoImpl
	config *Config
	db     *sql.DB
}

func New(config *Config) *Storage {
	return &Storage{
		config: NewConfig(),
		db:     nil,
	}
}

func (s *Storage) Open() error {
	// lazy conn. need ping.
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	// because we want check a real conn to db
	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) User() UserRepoImpl {
	if s.user != nil {
		return s.user
	}

	s.user = &UserRepo{
		storage: s,
	}
	return s.user
}
