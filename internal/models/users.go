package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type Token struct {
	UserId int
}

type User struct {
	gorm.Model
	Email    string `json: "email" validate: "email"`
	Password string `json: "password" validate: "gt>=6"`
	Token    string `json: "token"`
}

var validate *validator.Validate

func (user *User) Validate(u User) error {
	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) Create() {
	// TODO(Trung) Implement logic
}

func GetUser(u uint) *User {
	// TODO(Trung) Implement logic
	return nil
}
