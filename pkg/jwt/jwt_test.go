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
	const userId uint = 1
	jwtService := jwt.NewJWT(getSecret())
	token, err := jwtService.Create(jwt.JWTData{
		UserId: userId,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Invalid token")
	}
	if data.UserId != userId {
		t.Fatalf("UserId %d not equal to %d", data.UserId, userId)
	}
}
