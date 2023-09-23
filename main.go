package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "home")
	})
	r.GET("/portfolio", func(c *gin.Context) {
		c.String(http.StatusOK, "portfolio")
	})
	r.GET("/blog", func(c *gin.Context) {
		c.String(http.StatusOK, "blog")
	})
	r.GET("/project", func(c *gin.Context) {
		c.String(http.StatusOK, "project")
	})
	r.GET("/blog/:title", func(c *gin.Context) {
		title := c.Params.ByName("title")

		var ok = true
		if ok {
			c.JSON(http.StatusOK, gin.H{"title": title})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": "undefined", "status": "no value"})
		}
	})
	r.GET("/api/contact", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name": "formname", "email": "address"})
	})
	r.GET("/api/chat", func(c *gin.Context) {
		c.String(http.StatusOK, "contact")
	})
	return r
}

// api routes
// r.POST("/api/contact", func(c *gin.Context){
// 	return
// })

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
