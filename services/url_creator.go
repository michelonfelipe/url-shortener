package services

import (
	"github.com/michelonfelipe/url-shortener/database"
	"github.com/michelonfelipe/url-shortener/models"

	"errors"
	"os"
	"regexp"
	"strconv"

	"gorm.io/gorm"
)

type CreateParams struct {
	Original string `form:"original" json:"original" binding:"required"`
}

var validUrl = regexp.MustCompile(`^(http[s]?:\/\/(www\.)?|ftp:\/\/(www\.)?|www\.){1}([0-9A-Za-z-\.@:%_\+~#=]+)+((\.[a-zA-Z]{2,3})+)(\/(.)*)?(\?(.)*)?`)

func UrlCreator(createParams CreateParams) (models.Url, error) {
	var url models.Url
	var err error

	if !validUrl.MatchString(createParams.Original) {
		return url, errors.New("original url needs to be valid")
	}

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
