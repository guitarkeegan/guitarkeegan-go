package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func Portfolio(c *gin.Context) {
	c.String(http.StatusOK, "portfolio")
}

func Project(c *gin.Context) {
	c.String(http.StatusOK, "project")
}

func Blog(c *gin.Context) {
	c.String(http.StatusOK, "blog")
}

func BlogTitle(c *gin.Context) {
	title := c.Params.ByName("title")

	var ok = true
	if ok {
		c.JSON(http.StatusOK, gin.H{"title": title})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": "undefined", "status": "no value"})
	}
}
