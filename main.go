package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/configs"
	"github.com/greglownes/gogin1/models"
	"github.com/greglownes/gogin1/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	db.AutoMigrate(
		&models.User{},
		&models.PasswordReset{},
		&models.Topic{},
	)

	// setup router
	router := gin.Default()
	routes.SetupRoutes(db, router)
	router.Run() // localhost:8080
}
