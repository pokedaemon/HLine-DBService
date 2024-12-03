package postgres_test

import (
	model "hlservice-db/internal/app/models"
	"hlservice-db/internal/app/storage/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepo_Create(t *testing.T) {
	db, teardown := postgres.TestStorage(t, testDatabaseURL)
	defer teardown("Users")

	s := postgres.New(db)

	err := s.User().Create(&model.User{
		Name:        "Nikita Tutov",
		Email:       "nikita@tutov.com",
		PhoneNumber: "78120005252",
		Region:      "Курская область",
	})

	assert.NoError(t, err)
}

func TestUserRepo_FindByPhoneNumber(t *testing.T) {

}

func TestUserRepo_FindByName(t *testing.T) {

}

func TestUserRepo_FindByEmail(t *testing.T) {
	db, teardown := postgres.TestStorage(t, testDatabaseURL)
	defer teardown("Users")

	s := postgres.New(db)

	err := s.User().Create(&model.User{
		Name:        "Nikita Tutov",
		Email:       "nikita@tutov.com",
		PhoneNumber: "78120005252",
		Region:      "Курская область",
	})

	assert.NoError(t, err)

	result, err := s.User().FindByEmail("nikita@tutov.com")

	assert.NoError(t, err)

	assert.Equal(t, "Nikita Tutov", result.Name)
	assert.Equal(t, "78120005252", result.PhoneNumber)
}
