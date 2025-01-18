package repository

import (
	"errors"
)

type AuthRepository interface {
	ValidateUser(username, password string) (bool, error)
}

type authRepository struct {
	db *config.Database
}

func NewAuthRepository(db *config.Database) AuthRepository {
	return &authRepository{db: db}
}

func (repo *authRepository) ValidateUser(username, password string) (bool, error) {
	// This is a mock implementation. Replace with actual DB query.
	if username == "jon" && password == "shhh!" {
		return true, nil
	}
	return false, errors.New("invalid credentials")
}
