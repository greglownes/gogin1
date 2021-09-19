package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/configs"
	"github.com/greglownes/gogin1/models"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	// config
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()
	spew.Dump(config)

	// database setup
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto migrate
	db.AutoMigrate(&models.User{}, &models.PasswordReset{})

	// setup gin routing
	r := gin.Default()
	// routing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong, hello world, greg was here",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
