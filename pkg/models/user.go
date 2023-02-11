package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint64 `json:"id" gorm:"primary_key"`
	UserName string `json:"username"`
	Email    string `json:"email,omitempty" gorm:"unique"`
	Password string `json:"-" gorm:"size:255;not null"`
	Task     []Task `json:"tasks,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
}

type UserRegister struct {
	UserName        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return nil
}
