package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"guitarkeegan.com/go/forms"
)

func Contact(c *gin.Context) {
	var data forms.ContactForm
	if c.ShouldBind(&data) == nil {
		log.Println(data.Email)
		log.Println(data.Message)
	}
	// sets the content type as application json
	c.JSON(http.StatusOK, gin.H{"email": data.Email, "message": data.Message})
}

func Chat(c *gin.Context) {
	c.String(http.StatusOK, "chat")
}
