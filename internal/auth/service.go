package auth

import (
	"advpractice/internal/user"
	"advpractice/pkg/di"
	"advpractice/pkg/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{userRepository}
}

func (service *AuthService) login(data *LoginRequest) error {
	user, err := service.UserRepository.GetByEmail(data.Email)
	if err != nil {
		return errors.New(ErrNonExistentEmail)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return errors.New(ErrWrongPassword)
	}

	return nil
}

func (service *AuthService) register(data *RegistrateRequest) error {
	existingUser, _ := service.UserRepository.GetByEmail(data.Email)
	if existingUser != nil {
		return errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &user.User{
		Name:     data.Name,
		Password: string(hashedPassword),
		Email:    data.Email,
	}

	_, err = service.UserRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthService) getJWT(data jwt.JWTData, secret string) (string, error) {
	j := jwt.NewJWT(secret)
	token, err := j.Create(data)
	if err != nil {
		return "", errors.New(ErrJWT)
	}

	return token, nil
}
