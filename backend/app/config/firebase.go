package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	opt := option.WithCredentialsFile("service.json")

	config := &firebase.Config{
		DatabaseURL: "https://dbspending-default-rtdb.firebaseio.com",
	}

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("Firebase init error: %v", err)
	}

	App = app
}
