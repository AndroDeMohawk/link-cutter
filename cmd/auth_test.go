package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AndroDeMohawk/link-cutter/internal/auth"
	"github.com/AndroDeMohawk/link-cutter/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLoginSuccess(t *testing.T) {

	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "vasya2@gmail.com",
		Password: "1234",
	})

	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {

		t.Fatalf("resp.StatusCode = %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatalf("token is empty")
	}
	removeData(db)
}

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "vasya2@gmail.com",
		Password: "$2a$10$kpesPmh7axOtLIfz8WGXxeJnJI2bJw4ZXUW0TAakMrpcpx2teEQgW",
		Username: "Test_Vasya2",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "vasya2@gmail.com").
		Delete(&user.User{})
}

func TestLoginFail(t *testing.T) {

	db := initDb()
	initData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "vasya2@gmail.com",
		Password: "1",
	})

	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 401 {
		t.Fatalf("resp.StatusCode = %d", resp.StatusCode)
	}
	removeData(db)
}
