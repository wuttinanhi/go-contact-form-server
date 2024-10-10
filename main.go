package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Name    string
	Email   string
	Message string
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Automatically migrate the schema
	db.AutoMigrate(&Contact{})
	return db
}

func main() {
	db := initDB()
	r := gin.Default()

	// Serve the contact form
	r.GET("/", func(c *gin.Context) {
		c.File("pages/index.html")
	})

	// Serve the success page
	r.GET("/success", func(c *gin.Context) {
		c.File("pages/success.html")
	})

	// Handle form submission and redirect to success page
	r.POST("/submit", func(c *gin.Context) {
		var contact Contact
		if err := c.ShouldBind(&contact); err != nil {
			c.String(http.StatusBadRequest, "Invalid input")
			return
		}

		db.Create(&contact)

		// Redirect to success page
		c.Redirect(http.StatusFound, "/success")
	})

	r.Run(":3000")
}
