package main

import (
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/configs"
	"github.com/greglownes/gogin1/models"
	"github.com/greglownes/gogin1/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// config
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()
	spew.Dump(config)

	// database setup
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	// auto migrate
	db.AutoMigrate(
		&models.User{},
		&models.PasswordReset{},
		&models.Topic{},
		&models.Status{},
		&models.Product{},
	)

	// setup router
	router := gin.Default()
	routes.SetupRoutes(db, router)
	router.Run() // localhost:8080
}
