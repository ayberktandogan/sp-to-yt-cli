package main

import (
	"fmt"
	"log"

	app "github.com/ayberktandogan/melody/app/melody"
	"github.com/ayberktandogan/melody/config"
	"github.com/joho/godotenv"
)

var version string

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		log.Fatal("Error loading .env file")
	}

	config.InitEnvConfigs(version)
	app.Main()
}
