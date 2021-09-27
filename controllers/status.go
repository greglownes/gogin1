package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/models"
	"github.com/greglownes/gogin1/services/statusservice"
)

// interface to act as a table of content
type StatusControllerInterface interface {
	GetAll(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

// controller struct and new function used in setup
type statusController struct {
	statusService statusservice.StatusServiceInterface
}

func NewStatusController(ss statusservice.StatusServiceInterface) StatusControllerInterface {
	return &statusController{ss}
}

// input and out structs
// leaving out active for create, it will default to true
type createStatusInput struct {
	Status      string `json:"status" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// leaving out required
type updateStatusInput struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

// statusOutput is data returned to user
// it should be stripped of private data
// removed stuff: created_at, updated_at, deleted_at
type statusOutput struct {
	ID          uint   `json:"id"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

// route handlers

func (controller *statusController) GetAll(c *gin.Context) {
	statuses, err := controller.statusService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": statuses})
}

func (controller *statusController) GetByID(c *gin.Context) {
	// param id will by id from route or blank if not available, no possible error
	// getIDParam() in helper can return error
	id, err := getIDParam(c.Param(("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statuses, err := controller.statusService.GetByID(id)
	if err != nil {
		es := err.Error()
		if strings.Contains(es, "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": es})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": es})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": statuses})
}

func (controller *statusController) Create(c *gin.Context) {
	// validate input
	var rawinput createStatusInput
	if err := c.ShouldBindJSON(&rawinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// convert raw input to model
	status := models.Status{
		Status:      rawinput.Status,
		Description: rawinput.Description,
		Active:      true,
	}

	// create status using service
	if err := controller.statusService.Create(&status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": status})
}

func (controller *statusController) Update(c *gin.Context) {
	// param id will by id from route or blank if not available, no possible error
	id, err := getIDParam(c.Param(("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// make sure that status exists
	status, err := controller.statusService.GetByID(id)
	if err != nil {
		es := err.Error()
		if strings.Contains(es, "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": es})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": es})
		return
	}

	// validate input
	var rawinput updateStatusInput
	if err := c.ShouldBindJSON(&rawinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update record
	status.Status = rawinput.Status
	status.Description = rawinput.Description
	status.Active = rawinput.Active

	if err := controller.statusService.Update(status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// create response from model
	statusOutput := mapStatustoStatusOutput(status)
	c.JSON(http.StatusOK, gin.H{"data": statusOutput})
}

func (controller *statusController) Delete(c *gin.Context) {
	// OPEN make a common function here that can be shared between edit and delete

	// param id will by id from route or blank if not available, no possible error
	id, err := getIDParam(c.Param(("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// make sure that status exists
	status, err := controller.statusService.GetByID(id)
	if err != nil {
		es := err.Error()
		if strings.Contains(es, "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": es})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": es})
		return
	}

	if err := controller.statusService.Delete(status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return data that was deleted in response
	statusOutput := mapStatustoStatusOutput(status)
	c.JSON(http.StatusOK, gin.H{"data": statusOutput})
}

// private functions

func mapStatustoStatusOutput(status *models.Status) statusOutput {
	return statusOutput{
		ID:          status.ID,
		Status:      status.Status,
		Description: status.Description,
		Active:      status.Active,
	}
}
