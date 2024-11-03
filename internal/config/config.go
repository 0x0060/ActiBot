package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DiscordToken string `json:"discord_token"`
	APIPort      string `json:"api_port"`
}

var AppConfig Config

func LoadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return fmt.Errorf("error decoding config file: %v", err)
	}

	return nil
}
