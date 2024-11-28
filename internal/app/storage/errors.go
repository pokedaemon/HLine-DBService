package storage

import "errors"

var (
	ErrNotFound = errors.New("record in db not found")
)
