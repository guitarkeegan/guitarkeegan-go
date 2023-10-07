package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"guitarkeegan.com/go/routes"
)

func setupRouter() *gin.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	r.Use(gin.ErrorLogger())

	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/public", "./public")

	r.GET("/", routes.Home)
	r.GET("/portfolio", routes.Portfolio)
	r.GET("/portfolio/pass", routes.PortfolioPass)
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
