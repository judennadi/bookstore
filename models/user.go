package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

func (u *User) HashPassword() error {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(passwordByte)
	return nil
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func (u *User) CreateUser() (*User, error) {
	result := db.Create(&u)
	return u, result.Error
}

func GetUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUser(id string) (*User, *gorm.DB) {
	var user User
	db := db.Where("ID=?", id).Find(&user)
	return &user, db
}

func GetUserByEmail(email string) *User {
	var user User
	db.Where("Email=?", email).Find(&user)
	return &user
}

func DeleteUser(id uuid.UUID) User {
	var user User
	db.Where("ID=?", id).Delete(user)
	return user
}
