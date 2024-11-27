package storage

import model "hlservice-db/internal/app/models"

type OrderRepoImpl interface {
	Create(o *model.Order) (*int, error) // return a user_id in db
	UpdateStatus(o *model.Order) (*model.Order, error)
}

type OrderRepo struct {
	storage *Storage
}

func (r *OrderRepo) Create(o *model.Order) (*int, error) {
	query := `INSERT INTO Orders (user_id, good, count, status)
	VALUES ((SELECT id FROM Users WHERE email = $1), (SELECT id FROM Goods WHERE name = $2),
	$3, $4) RETURNING user_id`
	var id int
	err := r.storage.db.QueryRow(query, o.Email, o.Good, o.Count, o.Status).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *OrderRepo) UpdateStatus(o *model.Order) (*model.Order, error) {
	query := `UPDATE Orders 
	SET status = $1
	WHERE user_id = (SELECT id FROM Users WHERE email = $2)`
	_, err := r.storage.db.Exec(query, o.Status, o.Email)
	if err != nil {
		return nil, err
	}
	return o, nil
}
