package main

import (
	"github.com/felipe-michelon/url-shortener/database"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const RELEASE_ENV_NAME = "release"

func main() {
	if os.Getenv("GIN_MODE") != RELEASE_ENV_NAME {
		_ = godotenv.Load()
	}

	database.SetupDB()

	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", hello)

	return r
}

func hello(c *gin.Context) {
	c.String(200, "Hello there")
}
