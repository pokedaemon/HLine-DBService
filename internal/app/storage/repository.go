package storage

import model "hlservice-db/internal/app/models"

type UserRepository interface {
	Create(u *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByName(name string) (*model.User, error)
	FindByPhoneNumber(phoneNumber string) (*model.User, error)
}

type GoodRepository interface {
	Create(g *model.Good) error
	FindByName(name string) (*model.Good, error)
}

type OrderRepository interface {
	Create(o *model.Order) error // return a user_id in db of just created user
	UpdateStatus(status string) (*model.Order, error)
}
