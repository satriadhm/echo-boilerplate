package usecase

import (
	"errors"

	"github.com/satriadhm/echo-boilerplate/internal/auth/repository"
)

type AuthUsecase interface {
	Login(username, password string) (string, error)
}

type authUsecase struct {
	repo repository.AuthRepository
}

func NewAuthUsecase(repo repository.AuthRepository) AuthUsecase {
	return &authUsecase{repo: repo}
}

func (uc *authUsecase) Login(username, password string) (string, error) {
	valid, err := uc.repo.ValidateUser(username, password)
	if err != nil || !valid {
		return "", errors.New("login failed")
	}
	// Generate JWT token (mock implementation)
	token := "mock-jwt-token"
	return token, nil
}
