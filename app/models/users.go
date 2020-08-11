package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

//Users Model Strcut
type Users struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//GetUserByID query from database
func GetUserByID(ID interface{}) (*Users, error) {
	var user Users
	// Get record with primary key
	// 如果用了gorm.Model这里查询到的id字段会放在model里的ID上
	res := db.First(&user, ID)
	return &user, res.Error
}

//FindUserByEmail query by email
func FindUserByEmail(email string) (*Users, error) {
	var user Users
	res := db.Where("email = ?", email).First(&user)
	return &user, res.Error
}

//Authenticate password check
func (user *Users) Authenticate(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
