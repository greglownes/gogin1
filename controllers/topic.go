package controllers

import (
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/models"
	"github.com/greglownes/gogin1/services/topicservice"
)

// interface to act as a table of content
type TopicControllerInterface interface {
	GetAll(*gin.Context)
	GetByID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

// controller struct and new function used in setup
type topicController struct {
	topicService topicservice.TopicServiceInterface
}

func NewTopicController(ts topicservice.TopicServiceInterface) TopicControllerInterface {
	return &topicController{ts}
}

// input and out structs
type createTopicInput struct {
	Title string `json:"title" binding:"required"`
}

// leaving out required
type updateTopicInput struct {
	Title string `json:"title"`
}

// topicOutput is data returned to user
type topicOutput struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

// route handlers
func (controller *topicController) GetAll(c *gin.Context) {
	topics, err := controller.topicService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": topics})

}

func (controller *topicController) GetByID(c *gin.Context) {
	// param id will by id from route or blank if not available, no possible error
	id, err := getIDParam(c.Param(("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	topics, err := controller.topicService.GetByID(id)
	if err != nil {
		es := err.Error()
		if strings.Contains(es, "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": es})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": es})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": topics})
}

func (controller *topicController) Create(c *gin.Context) {
	// validate input
	var cti createTopicInput
	if err := c.ShouldBindJSON(&cti); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	spew.Dump(cti)

	// convert input to model
	topic := models.Topic{
		Title: cti.Title,
	}
	spew.Dump(topic)

	// create topic using service
	if err := controller.topicService.Create(&topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": topic})
}

func (controller *topicController) Update(c *gin.Context) {
	// param id will by id from route or blank if not available, no possible error
	id, err := getIDParam(c.Param(("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// make sure that topic exists
	topic, err := controller.topicService.GetByID(id)
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
	var input updateTopicInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update record
	topic.Title = input.Title
	if err := controller.topicService.Update(topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	topicOutput := topicOutput{
		ID:    topic.ID,
		Title: topic.Title,
	}
	c.JSON(http.StatusOK, gin.H{"data": topicOutput})
}

func (controller *topicController) Delete(c *gin.Context) {
	// OPEN make a common function here

	// param id will by id from route or blank if not available, no possible error
	id, err := getIDParam(c.Param(("id")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// make sure that topic exists
	topic, err := controller.topicService.GetByID(id)
	if err != nil {
		es := err.Error()
		if strings.Contains(es, "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": es})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": es})
		return
	}

	if err := controller.topicService.Delete(topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return data that was deleted
	// response
	topicOutput := topicOutput{
		ID:    topic.ID,
		Title: topic.Title,
	}
	c.JSON(http.StatusOK, gin.H{"data": topicOutput})
}
