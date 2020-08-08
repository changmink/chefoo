package api

import (
	"github.com/changmink/shafoo/model"
	"github.com/gin-gonic/gin"
)

func SignUp(g *gin.RouterGroup) {
	g.POST("", func(c *gin.Context) {
		var user model.UserForm
		if c.ShouldBindJSON(&user) == nil {
			model.AddUser(user)
			c.JSON(201, gin.H{
				"message": "Created",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "invalid user info",
			})
		}
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
