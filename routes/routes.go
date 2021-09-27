package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/controllers"
	"github.com/greglownes/gogin1/repositories"
	"github.com/greglownes/gogin1/services/pingservice"
	"github.com/greglownes/gogin1/services/productservice"
	"github.com/greglownes/gogin1/services/statusservice"
	"github.com/greglownes/gogin1/services/topicservice"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, router *gin.Engine) {
	// setup repositories
	topicRepo := repositories.NewTopicRepo(db)
	statusRepo := repositories.NewStatusRepo(db)
	productRepo := repositories.NewProductRepo(db)

	// setup services
	pingService := pingservice.NewPingService()
	topicService := topicservice.NewTopicService(topicRepo)
	statusService := statusservice.NewStatusService(statusRepo)
	productService := productservice.NewProductService(productRepo)

	// setup controllers
	pingController := controllers.NewPingController(pingService)
	topicController := controllers.NewTopicController(topicService)
	statusController := controllers.NewStatusController(statusService)
	productController := controllers.NewProductController(productService)

	// define routes
	api := router.Group("/api")
	api.GET("/ping", pingController.Ping)

	api.GET("/topic", topicController.GetAll)
	api.GET("/topic/:id", topicController.GetByID)
	api.POST("/topic", topicController.Create)
	api.PUT("/topic/:id", topicController.Update)
	api.DELETE("/topic/:id", topicController.Delete)

	api.GET("/status", statusController.GetAll)
	api.GET("/status/:id", statusController.GetByID)
	api.POST("/status", statusController.Create)
	api.PUT("/status/:id", statusController.Update)
	api.DELETE("/status/:id", statusController.Delete)

	api.GET("/product", productController.GetAll)
	api.GET("/product/:id", productController.GetByID)
	api.POST("/product", productController.Create)
	api.PUT("/product/:id", productController.Update)
	api.DELETE("/product/:id", productController.Delete)
}
