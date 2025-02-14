package repository

import (
	"database/sql"
	"ecommerce-inventory/model"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	row := repo.db.QueryRow(`SELECT id, username, password FROM users WHERE username = ?`, username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) RegisterUser(user *model.User) error {
	_, err := repo.db.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, user.Username, user.Password)
	return err
}
