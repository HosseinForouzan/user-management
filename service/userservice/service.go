package userservice

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/HosseinForouzan/user-management/entity"
	"github.com/HosseinForouzan/user-management/pkg/bcrypt"
)

type Repository interface {
	Register(u entity.User) (entity.User, error)
	IsEmailUnique(email string) (bool, error)
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
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
		Password: bcrypt.HashPassword(req.Password),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return RegisterResponse{createdUser}, nil
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string 	`json:"email"`
	PhoneNumber string `json:"phone_number"`
	Token string `json:"token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	if req.Email != ""{
		userByEmail, err := s.repo.GetUserByEmail(req.Email)
		if err != nil {
			return LoginResponse{}, fmt.Errorf(err.Error())
			}

		if !bcrypt.CheckPasswordHash(req.Password, userByEmail.Password) {
				return LoginResponse{}, fmt.Errorf("your credential is not correct")
	}
}

	if req.PhoneNumber != "" {
		userByPhoneNumber, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
		if err != nil {
			return LoginResponse{}, fmt.Errorf(err.Error())
		}

		if !bcrypt.CheckPasswordHash(req.Password, userByPhoneNumber.Password) {
			return LoginResponse{}, fmt.Errorf("your credential is not correct")
		}
	}

	token, err := CreateToken(req.Email)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("token is not valid %w", err)
	}
	


	return LoginResponse{Email: req.Email, PhoneNumber: req.PhoneNumber, Token: token}, nil
}

var secretKey = []byte("secret-key")

func CreateToken(email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "email": email, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

// func VerifyToken(tokenString string) error {
//    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//       return secretKey, nil
//    })
  
//    if err != nil {
//       return err
//    }
  
//    if !token.Valid {
//       return fmt.Errorf("invalid token")
//    }
  
//    return nil
// }