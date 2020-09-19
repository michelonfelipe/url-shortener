package main

import (
	"github.com/felipe-michelon/url-shortener/controllers"
	"github.com/felipe-michelon/url-shortener/database"
	"github.com/felipe-michelon/url-shortener/models"

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
