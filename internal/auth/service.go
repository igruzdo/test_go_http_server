package auth

import (
	"errors"
	"http_server/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepo}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.GetByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	user := &user.User{
		Email:    email,
		Password: "",
		Name:     name,
	}

	user, err := service.UserRepository.Create(user)

	if err != nil {
		return "", err
	}

	return user.Email, nil
}
