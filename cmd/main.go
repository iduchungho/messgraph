package main

import (
	"log"

	"github.com/joho/godotenv"
	"messgraph.com/m/internal/engine"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	engine := engine.NewEngine()
	engine.Run()
}