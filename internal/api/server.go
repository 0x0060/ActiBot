package api

import (
	"net/http"
	actibot "ActiBot/internal/logger"
	"ActiBot/internal/discord"

	"github.com/gin-gonic/gin"
)


func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/api/v1/activities/:user_id", getUserActivity)

	return r
}

func getUserActivity(c *gin.Context) {
	userID := c.Param("user_id")
	activity, exists := discord.GetActivities(userID)
	if !exists {
		actibot.Log(actibot.API, "Activity not found for user "+userID)
		c.JSON(http.StatusNotFound, gin.H{"error": "activity not found"})
		return
	}

	actibot.Logf(actibot.API, "Returning activity for user %s", userID)
	c.JSON(http.StatusOK, activity)
}



