package userservice

import (
	"fmt"

	"github.com/HosseinForouzan/user-management/entity"
)

type Repository interface {
	Register(u entity.User) (entity.User, error)
	IsEmailUnique(email string) (bool, error)
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	User entity.User `json:"user"`
}
func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {

	isEmailUnique, err := s.repo.IsEmailUnique(req.Email)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	isPhoneNumberUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	if !isEmailUnique {
		return RegisterResponse{}, fmt.Errorf("email is not unique")
	}

	if !isPhoneNumberUnique {
		return RegisterResponse{}, fmt.Errorf("phone number is not unique")
	}

	user := entity.User {
		ID: 0,
		Name: req.Name,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		Password: req.Password,
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return RegisterResponse{createdUser}, nil
}

