package user

import (
	"gin-template/models"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func GetAll() (data []User) {
	result := model.Db.Find(&data)

	if result.Error != nil {
		panic(result.Error)
	}

	return
}