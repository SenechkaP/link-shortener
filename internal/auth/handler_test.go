package auth

import (
	"advpractice/configs"
	"advpractice/internal/user"
	"advpractice/pkg/db"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, errors.New("Failed init mock db")
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, errors.New("Failed init gorm")
	}
	userRepo := user.NewUserRepository(&db.Db{DB: gormDb})
	handler := AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestHandlerLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"email", "password"})
	rows.AddRow("test@test.test", "$2a$10$MoNKNS7SiRjfpurKwcJQsOE1E9qjK.rIQ24UzPCfm4y3cn2QsUsLq")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	data, _ := json.Marshal(LoginRequest{
		Email:    "test@test.test",
		Password: "testpassword",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.login()(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandlerRegisterSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"name", "email", "password"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}))
	mock.ExpectCommit()

	data, _ := json.Marshal(RegistrateRequest{
		Name:     "test",
		Email:    "test@test.test",
		Password: "testpassword",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.register()(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("Expected %d, got %d", http.StatusCreated, w.Code)
	}
}
