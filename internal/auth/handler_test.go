package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndroDeMohawk/link-cutter/configs"
	"github.com/AndroDeMohawk/link-cutter/internal/auth"
	"github.com/AndroDeMohawk/link-cutter/internal/user"
	"github.com/AndroDeMohawk/link-cutter/pkg/db"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.Handler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database}))
	if err != nil {
		return nil, nil, err
	}

	userRepo := user.NewRepository(&db.Db{
		DB: gormDb,
	})
	handler := auth.Handler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		Service: auth.NewService(userRepo),
	}
	return &handler, mock, nil
}

func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("vasya@gmail.com", "$2a$10$kpesPmh7axOtLIfz8WGXxeJnJI2bJw4ZXUW0TAakMrpcpx2teEQgW")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	if err != nil {
		t.Fatalf("Unexpected error while bootstrap %v", err)
		return
	}
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "vasya@gmail.com",
		Password: "1234",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/login", reader)
	handler.Login()(w, req)
	if w.Result().StatusCode != 200 {
		t.Errorf("Expected status code 200 got %v", w.Result().StatusCode)
	}

}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	mock.ExpectQuery("SELECT").
		WillReturnRows(sqlmock.NewRows([]string{"email", "password", "name"}).
			AddRow("vasya@gmail.com", "$2a$10$kpesPmh7axOtLIfz8WGXxeJnJI2bJw4ZXUW0TAakMrpcpx2teEQgW", "Test_Vasya"))
	if err != nil {
		t.Fatalf("Unexpected error while bootstrap %v", err)
		return
	}
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "vasya@gmail.com",
		Password: "1234",
		Name:     "Test_Vasya",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/register", reader)
	handler.Register()(w, req)
	if w.Result().StatusCode != 200 {
		t.Errorf("Expected status code 200 got %v", w.Result().StatusCode)
	}

}
