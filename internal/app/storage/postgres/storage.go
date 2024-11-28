package postgres

import (
	"database/sql"
	"hlservice-db/internal/app/storage"

	_ "github.com/lib/pq" // driver
)

type Storage struct {
	user  *UserRepo
	good  *GoodRepo
	order *OrderRepo
	db    *sql.DB // Maybe use pool?
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// User return implementation of user repository
func (s *Storage) User() storage.UserRepository {
	if s.user != nil {
		return s.user
	}
	s.user = &UserRepo{
		storage: s,
	}
	return s.user
}

// Good return implementation repository of good
func (s *Storage) Good() storage.GoodRepository {
	if s.good != nil {
		return s.good
	}

	s.good = &GoodRepo{
		storage: s,
	}
	return s.good
}

// Order return implementation of order repository
func (s *Storage) Order() storage.OrderRepository {
	if s.order != nil {
		return s.order
	}

	s.order = &OrderRepo{
		storage: s,
	}
	return s.order
}
