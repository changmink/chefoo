package api

import (
	"github.com/gin-gonic/gin"
)

func GetProfileById(g *gin.RouterGroup) {
	g.GET("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is test",
		})
	})
}

func EditProfile(g *gin.RouterGroup) {
	g.PUT("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is test",
		})
	})
}
