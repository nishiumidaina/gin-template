package user

import (
	"gin-template/models/users"

	"github.com/gin-gonic/gin"
)

func ShowAllBlog(c *gin.Context) {
	data := user.GetAll()
	c.JSON(200, gin.H{
		"message": data,
	})
}