package api

import (
	"github.com/gin-gonic/gin"
)

func SignUp(g *gin.RouterGroup) {
	g.POST("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is test",
		})
	})
}

func Login(g *gin.RouterGroup) {
	g.PUT("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is test",
		})
	})
}

func Logout(g *gin.RouterGroup) {
	g.DELETE("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "this is test",
		})
	})
}
