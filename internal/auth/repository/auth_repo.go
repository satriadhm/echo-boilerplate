package repository

import (
	"database/sql"
	"errors"
)

type AuthRepository interface {
	ValidateUser(username, password string) (bool, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (repo *authRepository) ValidateUser(username, password string) (bool, error) {
	row := repo.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? AND password = ?", username, password)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, errors.New("error validating user")
	}
	return count > 0, nil
}
