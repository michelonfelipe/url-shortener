package models

import "github.com/michelonfelipe/url-shortener/database"

func Migrate() {
	database.DB.AutoMigrate(&Url{})
}
