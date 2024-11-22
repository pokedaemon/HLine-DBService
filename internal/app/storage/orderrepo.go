package storage

import model "hlservice-db/internal/app/models"

type OrderRepoImpl interface {
	Create(o *model.Order) (*model.Order, error)
	UpdateStatus(status string) (*model.Order, error)
}

type OrderRepo struct {
	storage *Storage
}
