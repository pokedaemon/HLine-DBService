package storage

import model "hlservice-db/internal/app/models"

type GoodRepoImpl interface {
	Create(g *model.Good) (*model.Good, error)
	FindByName(name string) (*model.Good, error)
}

type GoodRepo struct {
	storage *Storage
}

func (r *GoodRepo) Create(g *model.Good) (*model.Good, error) {
	return nil, nil
}

func (r *GoodRepo) FindByName(name string) (*model.Good, error) {
	return nil, nil
}