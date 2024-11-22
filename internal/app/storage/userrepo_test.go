package storage_test

import (
	model "hlservice-db/internal/app/models"
	"hlservice-db/internal/app/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepo_Create(t *testing.T) {
	s, teardown := storage.TestStorage(t, testDatabaseURL)
	defer teardown("Users")

	u, err := s.User().Create(&model.User{
		Name:        "Nikita Tutov",
		Email:       "nikita@tutov.com",
		PhoneNumber: "78120005252",
		Region:      "Курская область",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepo_FindByPhoneNumber(t *testing.T) {

}

func TestUserRepo_FindByName(t *testing.T) {

}

func TestUserRepo_FindByEmail(t *testing.T) {

}
