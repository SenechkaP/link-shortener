package auth

import (
	"advpractice/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (*MockUserRepository) Create(*user.User) (*user.User, error) {
	return &user.User{}, nil
}

func (*MockUserRepository) GetByEmail(string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	authService := NewAuthService(&MockUserRepository{})
	_, err := authService.register(&RegistrateRequest{
		Name:     "test",
		Email:    "test@test.test",
		Password: "testpassword",
	})
	if err != nil {
		t.Fatal(err)
	}
}
