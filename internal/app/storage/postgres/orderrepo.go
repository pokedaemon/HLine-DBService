package postgres

import model "hlservice-db/internal/app/models"

type OrderRepo struct {
	storage *Storage
}

func (r *OrderRepo) Create(o *model.Order) error {
	query := `INSERT INTO Orders (user_id, good, count, status)
	VALUES ((SELECT id FROM Users WHERE email = $1), (SELECT id FROM Goods WHERE name = $2),
	$3, $4) RETURNING user_id`
	var id int
	err := r.storage.db.QueryRow(query, o.Email, o.Good, o.Count, o.Status).Scan(&id)
	return err
}

func (r *OrderRepo) UpdateStatus(status string) (*model.Order, error) {
	return nil, nil
}
