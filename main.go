package main

import "github.com/gin-gonic/gin"

func main() {
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
