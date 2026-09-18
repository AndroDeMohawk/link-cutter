package jwt_test

import (
	"os"
	"testing"

	JWT "github.com/AndroDeMohawk/link-cutter/pkg/jwt"
)

func TestJWTCreate(t *testing.T) {
	const email = "vasya@gmail.com"
	jwtService := JWT.NewJWT(os.Getenv("TOKEN"))
	token, err := jwtService.Create(JWT.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is not valid")
	}
	if data.Email != email {
		t.Fatal("Data email does not match")
	}

}
