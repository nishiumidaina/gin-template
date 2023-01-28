package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func GetAll() (data []User) {
	result := Db.Find(&data)

	if result.Error != nil {
		panic(result.Error)
	}

	return
}