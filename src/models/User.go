package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email string `json:"email"`
	Password string `json:"password"`
	
}

func (u *User) hashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB)(err error){
	return u.hashPassword()
}