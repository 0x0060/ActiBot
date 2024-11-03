package main

import (
	"log"
	actibot "ActiBot/internal/logger"
	"ActiBot/internal/api"
	"ActiBot/internal/config"
	"ActiBot/internal/discord"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)

	dg, err := discord.NewDiscordClient(config.AppConfig.DiscordToken)
	if err != nil {
		actibot.Log(actibot.Error, "Failed to start Discord client: "+err.Error())
		return
	}
	defer dg.Close()

	server := api.NewServer()
	actibot.Log(actibot.System, "Starting API server on "+config.AppConfig.APIPort)

	if err := server.Run(config.AppConfig.APIPort); err != nil {
		actibot.Log(actibot.Error, "Failed to start API server: "+err.Error())
	}
}
