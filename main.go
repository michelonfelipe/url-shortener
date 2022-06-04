package main

import (
	"github.com/michelonfelipe/url-shortener/controllers"
	"github.com/michelonfelipe/url-shortener/database"
	"github.com/michelonfelipe/url-shortener/models"

	"os"

	"github.com/joho/godotenv"
)

const RELEASE_ENV_NAME = "release"

func main() {
	if os.Getenv("GIN_MODE") != RELEASE_ENV_NAME {
		_ = godotenv.Load()
	}

	database.SetupDB()

	models.Migrate()

	controllers.SetupRouter().Run()
}
