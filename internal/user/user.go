package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"profiler/internal/loader"
	"profiler/internal/message"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email    string `json: "email" validate: "email"`
	Password string `json: "password" validate: "gt=6"`
	Token    string `json: "token"`
}

var validate *validator.Validate

func (user *User) Create() map[string]interface{} {
	validate = validator.New()
	err := validate.Struct(user)

	if err != nil {
		return message.Message(false, "NG0001", "Validation error")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	loader.GetDatabase().Create(user)

	if user.ID <= 0 {
		message.Message(true, "NG0001", "Failed to crate user")
	}

	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, _ := token.SignedString([]byte(os.Getenv("token")))
	user.Token = tokenString
	// Why we need to delete password here?
	user.Password = ""

	res := message.Message(true, "OK0001", "User has been created")
	res["user"] = user
	return res
}

func GetUser(u uint) *User {
	// TODO(Trung) Implement logic
	return nil
}
