package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	TelegramToken  string `yaml:"telegram-token"`
	TelegramChatID string `yaml:"telegram-chat-id"`
}

func NewConfig() (*Config, error){
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
