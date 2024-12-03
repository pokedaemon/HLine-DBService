package postgres_test

import (
	model "hlservice-db/internal/app/models"
	"hlservice-db/internal/app/storage"
	"hlservice-db/internal/app/storage/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderRepo_Create(t *testing.T) {
	db, teardown := postgres.TestStorage(t, testDatabaseURL)
	defer teardown("Orders", "Users", "Goods")

	s := postgres.New(db)

	err := s.Good().Create(&model.Good{
		Name:  "Куртки зимние(мужские)",
		Count: 1000,
	})

	assert.NoError(t, err)

	err = s.User().Create(&model.User{
		Name:        "Nikita Tutov",
		Email:       "nikita@tutov.com",
		PhoneNumber: "78120005252",
		Region:      "Курская область",
	})

	assert.NoError(t, err)

	err = s.Order().Create(&model.Order{
		Email:  "nikita@tutov.com",
		Good:   "Куртки зимние(мужские)",
		Count:  25,
		Status: 1,
	})

	assert.NoError(t, err)

}

func TestOrderRepo_UpdateStatus(t *testing.T) {
	db, teardown := postgres.TestStorage(t, testDatabaseURL)
	defer teardown("Orders", "Users", "Goods")

	s := postgres.New(db)

	err := s.Good().Create(&model.Good{
		Name:  "Куртки зимние(женские)",
		Count: 1000,
	})

	assert.NoError(t, err)

	err = s.User().Create(&model.User{
		Name:        "Nikita Tutov",
		Email:       "nikita@tutov.com",
		PhoneNumber: "78120005252",
		Region:      "Курская область",
	})

	assert.NoError(t, err)

	err = s.Order().Create(&model.Order{
		ID:     "432543",
		Email:  "nikita@tutov.com",
		Good:   "Куртки зимние(женские)",
		Count:  52,
		Status: 1,
	})

	assert.NoError(t, err)

	err = s.Order().UpdateStatus(storage.OrderStatusSecond, "432543")

	assert.NoError(t, err)
}
