package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/index", Index)
	r.GET("/show_all", ShowAllBlog)

	return r
}