package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const TELEGRAM_API = "https://api.telegram.org/bot%s/sendMessage"

type TelegramLogger struct {
	url    string
	chatID string
	email  string
	client *http.Client
}

func NewTelegramLogger(token string, chatID string) *TelegramLogger {
	return &TelegramLogger{
		url:    fmt.Sprintf(TELEGRAM_API, token),
		chatID: chatID,
		client: NewTimeoutClient(),
	}
}

func (s *TelegramLogger) Logf(format string, args ...interface{}) {
	text := fmt.Sprintf("%s", fmt.Sprintf(format, args...))

	data := url.Values{}
	data.Add("chat_id", s.chatID)
	data.Add("text", text)

	req, err := http.NewRequest("POST", s.url, strings.NewReader(data.Encode()))
	if err != nil {
		log.Printf("failed to create request: %v\n", err)
		return
	}

	req.Header.Set("content-length", strconv.Itoa(len(data.Encode())))
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		log.Printf("failed to make request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("failed to get response: %v\n", resp)
		return
	}
}

