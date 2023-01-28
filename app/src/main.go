package main

import (
	"nothing-behind.com/sample_gin/controllers"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8000")
}