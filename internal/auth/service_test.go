package auth_test

import (
	"testing"

	"github.com/AndroDeMohawk/link-cutter/internal/auth"
	"github.com/AndroDeMohawk/link-cutter/internal/user"
)

type MockUserRepository struct {
}

func (r *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func (r *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{
		Email: "vasya@gmail.com",
	}, nil
}
func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "vasya@gmail.com"
	authService := auth.NewService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1234", "Vasya")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("email does not match")
	}
}
