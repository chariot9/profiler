package information

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"profiler/internal/loader"
	"profiler/internal/message"
)

type Information struct {
	gorm.Model
	Name   string `json: "name"`
	Phone  string `json: "phone"`
	UserId uint   `json: "userId"`
}

func (i *Information) Create() map[string]interface{} {
	validate := validator.New()

	err := validate.Struct(i)

	if err != nil {
		return message.Message(false, "NG0001", "User is not recognized")
	}

	loader.GetDatabase().Create(i)
	res := message.Message(true, "OK0001", "success")
	res["result"] = i
	return res
}

func GetInformation(user uint) []*Information {
	info := make([]*Information, 0)

	err := loader.GetDatabase().Table("information").Where("user_id = ?", user).Find(&info).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return info
}
