package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/controllers"
	"github.com/greglownes/gogin1/repositories"
	"github.com/greglownes/gogin1/services/pingservice"
	"github.com/greglownes/gogin1/services/topicservice"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, router *gin.Engine) {
	// setup repositories
	topicRepo := repositories.NewTopicRepo(db)

	// setup services
	pingService := pingservice.NewPingService()
	topicService := topicservice.NewTopicService(topicRepo)

	// setup controllers
	pingController := controllers.NewPingController(pingService)
	topicController := controllers.NewTopicController(topicService)

	// define routes
	api := router.Group("/api")
	api.GET("/ping", pingController.Ping)

	api.GET("/topic", topicController.GetAll)
	api.GET("/topic/:id", topicController.GetByID)
	api.POST("/topic", topicController.Create)
	api.PUT("/topic/:id", topicController.Update)
	api.DELETE("/topic/:id", topicController.Delete)
}
