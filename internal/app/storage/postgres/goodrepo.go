package postgres

import (
	"database/sql"
	model "hlservice-db/internal/app/models"
	"hlservice-db/internal/app/storage"
)

type GoodRepo struct {
	storage *Storage
}

func (r *GoodRepo) Create(g *model.Good) error {
	query := `INSERT INTO Goods (name, count) VALUES ($1, $2) RETURNING id`

	err := r.storage.db.QueryRow(query, g.Name, g.Count).Scan(&g.ID)
	return err
}

func (r *GoodRepo) FindByName(name string) (*model.Good, error) {
	query := `SELECT id, name, count FROM Goods WHERE name = $1`

	result := &model.Good{}

	err := r.storage.db.QueryRow(query, name).Scan(&result.ID, &result.Name, &result.Count)
	if err != nil {
		if err == sql.ErrNoRows {
			err = storage.ErrNotFound
		}
		return nil, err
	}
	return result, nil
}
