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

	g, err := s.Good().Create(&model.Good{
		Name:  "Лопата",
		Count: 1000,
	})

	assert.NoError(t, err)
	assert.NotNil(t, g)

	t.Log(g.ID)
}

func TestGoodRepo_FindByName(t *testing.T) {

}
