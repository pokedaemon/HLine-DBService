package storage

import model "hlservice-db/internal/app/models"

// TODO: Realize this methods and UNCOMMENT storage.User()!
type UserRepoImpl interface {
	Create(u *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByName(name string) (*model.User, error)
	FindByPhoneNumber(number string) (*model.User, error)
}

type UserRepo struct {
	storage *Storage
}

func (r *UserRepo) Create(u *model.User) (*model.User, error) {
	query := `INSERT INTO Users 
	(name, email, phone_number, region_id) VALUES 
	($1, $2, $3, (SELECT id FROM regions WHERE name = $4))
	RETURNING id`
	err := r.storage.db.QueryRow(
		query,
		u.Name, u.Email, u.PhoneNumber, u.Region,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}

func (r *UserRepo) FindByName(name string) (*model.User, error) {
	return nil, nil
}

func (r *UserRepo) FindByPhoneNumber(number string) (*model.User, error) {
	return nil, nil
}
