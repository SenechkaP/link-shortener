package jwt_test

import (
	"advpractice/pkg/jwt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const testEnvPath = "/Users/senechka/Desktop/golang/advpractice/pkg/jwt/.env"

func getSecret() string {
	err := godotenv.Load(testEnvPath)
	if err != nil {
		panic(err)
	}

	return os.Getenv("SECRET")
}

func TestJWTCreate(t *testing.T) {
	const email = "test@test.test"
	jwtService := jwt.NewJWT(getSecret())
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Invalid token")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal to %s", data.Email, email)
	}
}
