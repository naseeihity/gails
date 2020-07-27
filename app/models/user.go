package models

import (
	"github.com/jinzhu/gorm"
)

//User Model Strcut
type User struct {
	gorm.Model
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//GetUserByID query from database
func GetUserByID(ID int) (*User, error) {
	var user User
	res := db.First(&user, ID)
	return &user, res.Error
}
