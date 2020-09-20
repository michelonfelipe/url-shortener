package services

import (
	"github.com/felipe-michelon/url-shortener/database"
	"github.com/felipe-michelon/url-shortener/models"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateExistingUrl(t *testing.T) {
	existing := models.Url{Shortened: "shortened", Original: "http://abc.com"}

	database.DB.Create(&existing)
	defer database.DB.Unscoped().Delete(&existing)

	c := CreateParams{Original: existing.Original}

	result, err := UrlCreator(c)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.ID, existing.ID)
}

func TestCreateNonExistingUrl(t *testing.T) {
	c := CreateParams{Original: "http://abcd.com"}

	result, err := UrlCreator(c)
	defer database.DB.Unscoped().Delete(&result)

	assert.Equal(t, err, nil)
	assert.Equal(t, result.Original, c.Original)
	assert.NotEqual(t, result.Shortened, "")
}
