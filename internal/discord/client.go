package discord

import (
	"fmt"
	"sync"
	"time"

	actibot "ActiBot/internal/logger"
	"ActiBot/internal/models"
	"github.com/bwmarrin/discordgo"
)

var (
	activities = make(map[string]models.Activity)
	mutex      = &sync.Mutex{}
)

func NewDiscordClient(token string) (*discordgo.Session, error) {
	if token == "" {
		return nil, fmt.Errorf("discord token is empty")
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		actibot.Log(actibot.Error, "Failed to create Discord client: "+err.Error())
		return nil, err
	}

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildPresences)

	dg.AddHandler(activityHandler)

	err = dg.Open()
	if err != nil {
		actibot.Log(actibot.Error, "Error opening connection: "+err.Error())
		return nil, err
	}
	actibot.Log(actibot.System, "Bot is now running")
	return dg, nil
}

func activityHandler(s *discordgo.Session, event *discordgo.PresenceUpdate) {
	mutex.Lock()
	defer mutex.Unlock()

	userActivity := models.Activity{
		UserID:   event.User.ID,
		Username: event.User.Username,
		Status:   string(event.Status),
		Platform: "Discord",
		Updated:  time.Now(),
	}

	for _, activity := range event.Activities {
		if activity.Name == "Spotify" {
			userActivity.Activity = activity.Name
			userActivity.Type = activityTypeToString(activity.Type)
			userActivity.Details = activity.Details
			userActivity.State = activity.State
		}
	}

	if event.ClientStatus.Desktop != "" {
		userActivity.Platform = "PC"
	} else if event.ClientStatus.Mobile != "" {
		userActivity.Platform = "Mobile"
	} else if event.ClientStatus.Web != "" {
		userActivity.Platform = "Web"
	}

	activities[event.User.ID] = userActivity

	actibot.Logf(actibot.API, "Updated activity for user %s (%s): Type: %s, Details: %s", 
		userActivity.UserID, userActivity.Username, userActivity.Type, userActivity.Details)
}
func activityTypeToString(activityType discordgo.ActivityType) string {
	switch activityType {
	case 0: // Playing
		return "Playing"
	case 1: // Streaming
		return "Streaming"
	case 2: // Listening
		return "Listening"
	case 3: // Watching
		return "Watching"
	default:
		return "Unknown"
	}
}

func GetActivities(userID string) (models.Activity, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	activity, exists := activities[userID]
	return activity, exists
}
