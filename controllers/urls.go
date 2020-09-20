package controllers

import (
	"github.com/felipe-michelon/url-shortener/database"
	"github.com/felipe-michelon/url-shortener/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUrl(c *gin.Context) {
	var url models.Url

	err := database.DB.Where("shortened = ?", c.Param("shortened")).First(&url).Error

	if err != nil {
		c.String(http.StatusNotFound, "Url not found")
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Original)
}
