package user

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	StoreUser(u User) (User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) StoreUser(u User) (User, error) {
	query := `
	   INSERT INTO users (first_name, last_name, email, password)
	   VALUES ($1, $2, $3, $4)
	   RETURNING id, first_name, last_name, email, password, is_admin
	`

	var createdUser User

	err := r.db.Get(&createdUser, query, u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		return User{}, err
	}

	return createdUser, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	query := `
		SELECT id, first_name, last_name, email, password, is_admin
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	var user User
	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Password != pass {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
