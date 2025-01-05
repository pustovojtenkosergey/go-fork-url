package main

import (
	"log"
	"myserv/app"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	app := app.NewApp(app.NewConfig())

	app.Start()
}
