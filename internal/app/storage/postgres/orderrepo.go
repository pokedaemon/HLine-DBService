package postgres

import model "hlservice-db/internal/app/models"

type OrderRepo struct {
	storage *Storage
}

func (r *OrderRepo) Create(o *model.Order) error {
	query := `INSERT INTO Orders (id, user_id, good, count, status)
	VALUES ($1, (SELECT id FROM Users WHERE email = $2), (SELECT id FROM Goods WHERE name = $3),
	$4, $5)`
	_, err := r.storage.db.Exec(query, o.ID, o.Email, o.Good, o.Count, o.Status)
	return err
}

func (r *OrderRepo) UpdateStatus(status string, order string) error {
	query := `UPDATE Orders SET status = (SELECT id FROM Status WHERE name = $1) 
	WHERE id = $2`
	_, err := r.storage.db.Exec(query, status, order)

	return err
}
