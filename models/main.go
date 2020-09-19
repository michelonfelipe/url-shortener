package models

import "github.com/felipe-michelon/url-shortener/database"

func Migrate() {
	database.DB.AutoMigrate(&Url{})
}
