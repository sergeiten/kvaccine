package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TelegramToken  string `yaml:"telegram_token"`
	TelegramChatID string `yaml:"telegram_chat_id"`
	Sleep          int    `yaml:"sleep"`
	NotifyAfter    int    `yaml:"notify_after"`
	Cookie         string `yaml:"cookie"`
	Locations      []struct {
		Name        string `yaml:"name"`
		BottomRight struct {
			X float64 `yaml:"x"`
			Y float64 `yaml:"y"`
		} `yaml:"bottom_right"`
		TopLeft struct {
			X float64 `yaml:"x"`
			Y float64 `yaml:"y"`
		} `yaml:"top_left"`
	} `yaml:"locations"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	file, err := os.Open("config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
