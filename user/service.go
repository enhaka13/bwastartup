package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{
		Name:       input.Name,
		Email:      input.Email,
		Occupation: input.Occupation,
		Role:       "user",
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	user.PasswordHash = string(PasswordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return User{}, err
	}

	if user.ID == 0 {
		return User{}, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return User{}, err
	}

	return user, nil
}
