package controller

import (
	"gin-template/controllers/users"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/index", Index)
	r.GET("/show_all", user.ShowAllBlog)

	return r
}