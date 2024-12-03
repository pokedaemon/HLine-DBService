package postgres

import model "hlservice-db/internal/app/models"

type UserRepo struct {
	storage *Storage
}

func (r *UserRepo) Create(u *model.User) error {
	query := `INSERT INTO Users 
	(name, email, phone_number, region_id) VALUES 
	($1, $2, $3, (SELECT id FROM regions WHERE name = $4))
	RETURNING id`
	err := r.storage.db.QueryRow(
		query,
		u.Name, u.Email, u.PhoneNumber, u.Region,
	).Scan(&u.ID)

	return err
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	query := ``
	err := r.storage.db.QueryRow(
		query,
		email).Scan(
		&user.ID, &user.Name, &user.PhoneNumber, &user.Region)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *UserRepo) FindByName(name string) (*model.User, error) {
	return nil, nil
}

func (r *UserRepo) FindByPhoneNumber(phoneNumber string) (*model.User, error) {
	return nil, nil
}
