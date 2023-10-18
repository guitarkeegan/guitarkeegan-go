package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	data "guitarkeegan.com/go/controller"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func Portfolio(c *gin.Context) {
	c.HTML(http.StatusOK, "portfolio.tmpl", nil)
}

func PortfolioPass(c *gin.Context) {
	c.HTML(http.StatusOK, "portfolio-pass.tmpl", nil)
}

func Project(c *gin.Context) {
	projectId := c.Param("id")
	num, err := strconv.Atoi(projectId)

	project, err := data.GetProjectById(num)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve project"})
	}
	c.HTML(http.StatusOK, "project.tmpl", project)
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
