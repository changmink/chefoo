package api

import (
	"github.com/gin-gonic/gin"
)

func SearchParties(g *gin.RouterGroup) {
	g.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is test",
		})
	})
}

func CreateParty(g *gin.RouterGroup) {
	g.POST("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is test",
		})
	})
}

func LeaveParty(g *gin.RouterGroup) {
	g.DELETE("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is test",
		})
	})
}

func JoinPartyById(g *gin.RouterGroup) {
	g.PUT("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is test",
		})
	})
}

func GetPartyById(g *gin.RouterGroup) {
	g.Group("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is test",
		})
	})
}
