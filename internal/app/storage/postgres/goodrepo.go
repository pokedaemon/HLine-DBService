package postgres

import model "hlservice-db/internal/app/models"

type GoodRepo struct {
	storage *Storage
}

func (r *GoodRepo) Create(g *model.Good) error {
	query := `INSERT INTO Goods (name, count) VALUES ($1, $2) RETURNING id`

	err := r.storage.db.QueryRow(query, g.Name, g.Count).Scan(&g.ID)
	return err
}

func (r *GoodRepo) FindByName(name string) (*model.Good, error) {
	return nil, nil
}
