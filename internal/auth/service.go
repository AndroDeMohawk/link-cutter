package auth

import (
	"errors"

	"github.com/AndroDeMohawk/link-cutter/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepository *user.Repository
}

func NewService(userRepository *user.Repository) *Service {
	return &Service{UserRepository: userRepository}
}

func (s *Service) Register(email, password, name string) (string, error) {
	existedUser, _ := s.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Username: name,
	}
	_, err = s.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil

}

func (s *Service) Login(email, password string) (string, error) {
	existedUser, _ := s.UserRepository.FindByEmail(email)
	if existedUser == nil {
		return "", errors.New(ErrUserExists)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}
	return existedUser.Email, nil
}
