package storage

import model "hlservice-db/internal/app/models"

type UserRepoImpl interface {
	Create(u *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type UserRepo struct {
	storage *Storage
}

func (r *UserRepo) Create(u *model.User) (*model.User, error) {
	err := r.storage.db.QueryRow(
		"insert into users (name, email, phone_number, region) values ($1, $2, $3, (select id from region where name = $4)) returning id",
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
