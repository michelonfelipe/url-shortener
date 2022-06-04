package controllers

import (
	"github.com/michelonfelipe/url-shortener/database"
	"github.com/michelonfelipe/url-shortener/models"
	"github.com/michelonfelipe/url-shortener/services"

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

func CreateUrl(c *gin.Context) {
	var params services.CreateParams
	var err error
	var url models.Url

	err = c.ShouldBind(&params)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	url, err = services.UrlCreator(params)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, url)
}
