package main

import (
	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/configs"
	"github.com/joho/godotenv"
	"github.com/davecgh/go-spew/spew"
	"log"
)

func main() {
	// config
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()
	spew.Dump(config)
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
