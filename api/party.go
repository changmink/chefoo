package api

import (
	"strconv"

	"github.com/changmink/shafoo/model"
	"github.com/gin-gonic/gin"
)

func SearchParties(g *gin.RouterGroup) {
	g.GET("", func(c *gin.Context) {
		latitude := c.Query("latitude")
		longitude := c.Query("longitude")
		distence := c.Query("distence")
		dist, err := strconv.ParseFloat(distence, 4)
		if err != nil {
			panic(err)
		}
		parties := model.SearchParties(latitude, longitude, dist)
		c.JSON(200, parties)
	})
}

func CreateParty(g *gin.RouterGroup) {
	g.POST("", func(c *gin.Context) {
		var party model.PartyForm
		if c.ShouldBindJSON(&party) == nil {
			id := model.CreateParty(party)
			c.JSON(201, gin.H{
				"message": "Party is created",
				"id":      id,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Invalid party form",
			})
		}
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
		partyId := c.Param("id")
		userId := c.Query("userId")
		if model.JoinPartyById(partyId, userId) == nil {
			c.JSON(200, gin.H{
				"message": "Join Party",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Over total People",
			})
		}
	})
}

func GetPartyById(g *gin.RouterGroup) {
	g.Group("/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is test",
		})
	})
}
