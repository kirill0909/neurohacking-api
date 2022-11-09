package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirill0909/neurohacking-api/models"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (name, email, password_hash, date_creation, last_update)
	 VALUES ($1, $2, $3, now(), now()) RETURNING id`, usersTable)
	row := u.db.QueryRow(query, user.Name, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	// Get only the id from DB, beacause only user id needed to generate the token
	query := fmt.Sprintf(`SELECT id FROM %s WHERE email=$1 AND password_hash=$2`, usersTable)

	err := u.db.Get(&user, query, email, password)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}
