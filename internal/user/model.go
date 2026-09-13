package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"username" gorm:"unique" gorm:"index"`
	Password string `json:"password" gorm:"password"`
	Email    string `json:"email" gorm:"email" gorm:"unique"`
}
