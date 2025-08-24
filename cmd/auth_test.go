package main

import (
	"advpractice/internal/auth"
	"advpractice/internal/user"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const testEnvPath = "/Users/senechka/Desktop/golang/advpractice/cmd/.env"

func initDb() *gorm.DB {
	err := godotenv.Load(testEnvPath)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Name:     "test",
		Email:    "test@test.test",
		Password: "$2a$10$MoNKNS7SiRjfpurKwcJQsOE1E9qjK.rIQ24UzPCfm4y3cn2QsUsLq",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "test@test.test").
		Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	db := initDb()
	initData(db)
	defer removeData(db)

	ctx := t.Context()

	ts := httptest.NewServer(App(ctx, testEnvPath))
	defer ts.Close()

	data, _ := json.Marshal(auth.LoginRequest{
		Email:    "test@test.test",
		Password: "testpassword",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}
	var result auth.TokenResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	if result.Token == "" {
		t.Fatal("Token empty")
	}
}

func TestLoginFail(t *testing.T) {
	db := initDb()
	initData(db)
	defer removeData(db)

	ctx := t.Context()

	ts := httptest.NewServer(App(ctx, testEnvPath))
	defer ts.Close()

	data, _ := json.Marshal(auth.LoginRequest{
		Email:    "test@test.test",
		Password: "wrongtestpassword",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, res.StatusCode)
	}
}

func TestRegisterSuccess(t *testing.T) {
	db := initDb()
	defer removeData(db)

	ctx := t.Context()

	ts := httptest.NewServer(App(ctx, testEnvPath))
	defer ts.Close()

	data, _ := json.Marshal(auth.RegistrateRequest{
		Name:     "test",
		Email:    "test@test.test",
		Password: "testpassword",
	})

	res, err := http.Post(ts.URL+"/auth/register", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 201 {
		t.Fatalf("Expected %d got %d", 201, res.StatusCode)
	}
	var result auth.TokenResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	if result.Token == "" {
		t.Fatal("Token empty")
	}
}
