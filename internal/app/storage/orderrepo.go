package storage

import model "hlservice-db/internal/app/models"

type OrderRepoImpl interface {
	Create(o *model.Order) (*model.Order, error)
	UpdateStatus(status string) (*model.Order, error)
}

type OrderRepo struct {
	storage *Storage
}

func (r *OrderRepo) Create(o *model.Order) (*model.Order, error) {
	return nil, nil
}

func (r *OrderRepo) UpdateStatus(status string) (*model.Order, error) {
	return nil, nil
}
