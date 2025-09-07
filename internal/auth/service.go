package auth

import (
	"errors"
	"http_server/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepo}
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, _ := service.UserRepository.GetByEmail(email)

	if existedUser == nil {
		return "", errors.New(ErrWrongCreds)
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))

	if err != nil {
		return "", err
	}

	return existedUser.Email, nil
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.GetByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	user := &user.User{
		Email:    email,
		Password: string(hashedPass),
		Name:     name,
	}

	createdUser, err := service.UserRepository.Create(user)

	if err != nil {
		return "", err
	}

	return createdUser.Email, nil
}
