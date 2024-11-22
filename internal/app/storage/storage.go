package storage

import (
	"database/sql"

	_ "github.com/lib/pq" // driver
)

type Storage struct {
	user   UserRepoImpl
	good   GoodRepoImpl
	order  OrderRepoImpl
	config *Config
	db     *sql.DB // Maybe use pool?
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

// Close just close connection
func (s *Storage) Close() error {
	return s.db.Close()
}

// User return implementation of user repository
func (s *Storage) User() UserRepoImpl {
	if s.user != nil {
		return s.user
	}
	// TODO: Uncomment this after realization UserRepo
	s.user = &UserRepo{
		storage: s,
	}
	return s.user
}

// Good return implementation repository of good
func (s *Storage) Good() GoodRepoImpl {
	if s.good != nil {
		return s.good
	}

	s.good = &GoodRepo{
		storage: s,
	}
	return s.good
}

// Order return implementation of order repository
func (s *Storage) Order() OrderRepoImpl {
	if s.order != nil {
		return s.order
	}

	s.order = &OrderRepo{
		storage: s,
	}
	return s.order
}
