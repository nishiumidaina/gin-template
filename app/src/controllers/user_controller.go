package controller

import (
	"gin-template/models"

	"github.com/gin-gonic/gin"
)

func ShowAllBlog(c *gin.Context) {
	data := model.GetAll()
	c.JSON(200, gin.H{
		"message": data,
	})
}