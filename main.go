package main

import "log"

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}

	resp, err := app.me()
	if err != nil {
		log.Fatalf("failed to get me: %v", err)
	}

	log.Printf("me: %+v", resp)

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}
