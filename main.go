package main

import (
	"github.com/gin-gonic/gin"
	"guitarkeegan.com/go/routes"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/public", "./public")

	r.GET("/", routes.Home)
	r.GET("/portfolio", routes.Portfolio)
	r.GET("/blog", routes.Blog)
	r.GET("/project", routes.Project)
	r.GET("/blog/:title", routes.BlogTitle)

	r.POST("/api/contact", routes.Contact)
	r.POST("/api/chat", routes.Chat)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
