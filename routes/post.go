package routes

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"guitarkeegan.com/go/forms"
)

func Contact(c *gin.Context) {
	var data forms.ContactForm
	if c.ShouldBind(&data) == nil {
		log.Println(data.Email)
		log.Println(data.Message)
	}
	fmt.Println("open db connection")
	fp, err := filepath.Abs("mailingList.db")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "we messed up"})
		return
	}

	db, err := sql.Open("sqlite3", fp)

	if err != nil {
		log.Println(err)
		err := errors.New("Error accessing db")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"email": data.Email, "message": data.Message})
		return
	}

	const create string = `
  CREATE TABLE IF NOT EXISTS user (
  id INTEGER NOT NULL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  message TEXT
  );`

	_, err = db.Exec(create)
	if err != nil {
		log.Println(err)
		err := errors.New("Error accessing db")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "we messed up again"})
		return
	}
	// fmt.Println("prepair db check...")
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM user WHERE email = ?", data.Email).Scan(&count)
	if err != nil {
		log.Println(err)
		err := errors.New("Error creating sql select query")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "we messed up again"})
		return
	}

	if count > 0 {
		fmt.Println("already signed up")
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists\n"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO user (email, message) VALUES(?,?)")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	defer stmt.Close()

	userAdded, err := stmt.Exec(data.Email, data.Message)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}

	userRows, err := userAdded.RowsAffected()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}

	fmt.Printf("User was added! %v\n", userRows)

	// sets the content type as application json
	c.JSON(http.StatusCreated, gin.H{"message": "user added!"})
}

func Chat(c *gin.Context) {
	c.String(http.StatusOK, "chat")
}
