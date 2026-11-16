package service

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ClientID string
var ClientSecret string

func GenDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ClientID = os.Getenv("ClientID")
	ClientSecret = os.Getenv("Secret")
}
