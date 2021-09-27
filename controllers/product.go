package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/services/productservice"
)

// interface to act as a table of content
type ProductControllerInterface interface {
	GetAll(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

// controller struct and new function used in setup
type productController struct {
	productService productservice.ProductServiceInterface
}

func NewProductController(ss productservice.ProductServiceInterface) ProductControllerInterface {
	return &productController{ss}
}

// route handlers

func (controller *productController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "tobedone"})
}

func (controller *productController) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "tobedone"})
}

func (controller *productController) Create(c *gin.Context) {
	// OPEN authentication/authority

	// raw input -> model
	product, err := controller.productService.CreateModelForAddFromRawInput(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// add to db
	if err := controller.productService.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create response with clean/santitize from model and return
	productOutput := controller.productService.MapProductToProductOutput(&product)
	c.JSON(http.StatusOK, gin.H{"data": productOutput})
}

func (controller *productController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "tobedone"})
}

func (controller *productController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "tobedone"})
}