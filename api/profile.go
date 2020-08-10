package api

import (
	"github.com/changmink/shafoo/model"
	"github.com/gin-gonic/gin"
)

func GetProfileById(g *gin.RouterGroup) {
	g.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		profile := model.GetProfileById(id)
		c.JSON(200, gin.H{
			"message": "Found",
			"profile": profile,
		})
	})
}

func EditProfileById(g *gin.RouterGroup) {
	g.PUT("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Edited",
		})
	})
}
