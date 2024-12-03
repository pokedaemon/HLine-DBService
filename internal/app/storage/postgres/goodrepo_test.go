package postgres_test

import (
	model "hlservice-db/internal/app/models"
	"hlservice-db/internal/app/storage/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoodRepo_Create(t *testing.T) {
	db, teardown := postgres.TestStorage(t, testDatabaseURL)
	defer teardown("Goods")

	s := postgres.New(db)

	err := s.Good().Create(&model.Good{
		Name:  "Лопата",
		Count: 1000,
	})

	assert.NoError(t, err)
}

func TestGoodRepo_FindByName(t *testing.T) {
	db, teardown := postgres.TestStorage(t, testDatabaseURL)
	defer teardown("Goods")

	s := postgres.New(db)

	err := s.Good().Create(&model.Good{
		Name:  "Лопата",
		Count: 1000,
	})

	assert.NoError(t, err)

	result, err := s.Good().FindByName("Лопата")

	assert.NoError(t, err)

	assert.Equal(t, 1000, result.Count)
	assert.Equal(t, "Лопата", result.Name)
}
