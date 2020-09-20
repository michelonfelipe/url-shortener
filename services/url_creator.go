package services

import (
	"github.com/felipe-michelon/url-shortener/database"
	"github.com/felipe-michelon/url-shortener/models"

	"os"
	"strconv"

	"gorm.io/gorm"
)

type CreateParams struct {
	Original string `json:"original"`
}

func UrlCreator(createParams CreateParams) (models.Url, error) {
	var url models.Url
	var err error

	err = database.DB.First(&url, "original = ?", createParams.Original).Error

	if err != gorm.ErrRecordNotFound {
		return url, err
	}

	n, _ := strconv.Atoi(os.Getenv("SHORTENED_URL_CHARS_NUMBER"))
	url = models.Url{
		Original:  createParams.Original,
		Shortened: RandomStringGenerator(n),
	}

	err = database.DB.Create(&url).Error

	return url, err
}
