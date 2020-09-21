package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", home)
	r.GET("/:shortened", FindUrl)
	r.POST("/urls", CreateUrl)

	return r
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
