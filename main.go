package main

import (
	"SahamRakyatTechnicalTest-Golang/api"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file not FOUND")
	}
	api.Init()
}
