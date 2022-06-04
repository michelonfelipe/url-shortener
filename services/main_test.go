package services

import (
	"github.com/michelonfelipe/url-shortener/database"
	"github.com/michelonfelipe/url-shortener/models"

	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load("../.env")
	database.SetupDB()
	models.Migrate()

	exitVal := m.Run()

	os.Exit(exitVal)
}
