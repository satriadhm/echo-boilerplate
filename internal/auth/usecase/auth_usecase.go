package usecase

import "errors"

type AuthUsecase interface {
	Login(username, password string) (string, error)
}

type authUsecase struct {
	repo AuthRepository
}

func NewAuthUsecase(repo AuthRepository) AuthUsecase {
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
