package main

import (
	"github.com/joho/godotenv"
	"log"
	"messgraph.com/m/modules/engine"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := engine.New()
	app.Run()
}
