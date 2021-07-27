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
