package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"guitarkeegan.com/go/forms"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
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
	r.POST("/api/contact", func(c *gin.Context) {
		var data forms.ContactForm
		if c.ShouldBind(&data) == nil {
			log.Println(data.Email)
			log.Println(data.Message)
		}
		// sets the content type as application json
		c.JSON(http.StatusOK, gin.H{"email": data.Email, "message": data.Message})
	})
	r.POST("/api/chat", func(c *gin.Context) {
		c.String(http.StatusOK, "chat")
	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
