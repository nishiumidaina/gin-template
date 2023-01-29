package main

import (
	"gin-template/controllers"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8000")
}