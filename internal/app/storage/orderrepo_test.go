package storage_test

import (
	model "hlservice-db/internal/app/models"
	"hlservice-db/internal/app/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderRepo_Create(t *testing.T) {
	s, teardown := storage.TestStorage(t, testDatabaseURL)
	defer s.Close()
	defer teardown("Orders", "Users")

	_, err := s.User().Create(&model.User{
		Name:        "Nikita Tutov",
		Email:       "nikita@tutov.com",
		PhoneNumber: "78120005252",
		Region:      "Курская область",
	})

	o, err := s.Order().Create(&model.Order{
		Email:  "nikita@tutov.com",
		Good:   "Куртки зимние(мужские)",
		Count:  25,
		Status: 1,
	})

	t.Log(*o)

	assert.NoError(t, err)
	assert.NotNil(t, o)
}

func TestOrderRepo_UpdateStatus(t *testing.T) {
	s, teardown := storage.TestStorage(t, testDatabaseURL)
	defer s.Close()
	defer teardown("Orders", "Users")

	_, err := s.User().Create(&model.User{
		Name:        "Nikita Tutov",
		Email:       "nikita@tutov.com",
		PhoneNumber: "78120005252",
		Region:      "Курская область",
	})

	o, err := s.Order().Create(&model.Order{
		Email:  "nikita@tutov.com",
		Good:   "Куртки зимние(женские)",
		Count:  52,
		Status: 1,
	})

	assert.NoError(t, err)
	assert.NotNil(t, o)

	status, err := s.Order().UpdateStatus(&model.Order{
		Email:  "nikita@tutov.com",
		Status: 2,
	})

	assert.NoError(t, err)
	assert.NotNil(t, status)
}
