package main

import (
	"fmt"
	"os"

	"github.com/AndroDeMohawk/link-cutter/internal/link"
	"github.com/AndroDeMohawk/link-cutter/internal/stat"
	"github.com/AndroDeMohawk/link-cutter/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Overload(".env")
	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("DSN"))
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
	if err != nil {
		panic("Migration failed")
	}

}
