package topicservice

import (
	"errors"

	"github.com/greglownes/gogin1/models"
	"github.com/greglownes/gogin1/repositories"
)

type TopicServiceInterface interface {
	GetAll() ([]models.Topic, error)
	GetByID(id uint) (*models.Topic, error)
	Create(*models.Topic) error
	Update(*models.Topic) error
	Delete(*models.Topic) error
}

// hold repos and other dependencies
type topicService struct {
	TopicRepo repositories.TopicRepoInterface
}

func NewTopicService(tr repositories.TopicRepoInterface) TopicServiceInterface {
	return &topicService{
		TopicRepo: tr,
	}
}

func (ts *topicService) GetAll() ([]models.Topic, error) {
	topics, err := ts.TopicRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return topics, nil
}

func (ts *topicService) GetByID(id uint) (*models.Topic, error) {
	if id == 0 {
		return nil, errors.New("id param is required")
	}
	topic, err := ts.TopicRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func (ts *topicService) Create(topic *models.Topic) error {
	return ts.TopicRepo.Create(topic)
}

func (ts *topicService) Update(topic *models.Topic) error {
	return ts.TopicRepo.Update(topic)
}

func (ts *topicService) Delete(topic *models.Topic) error {
	ts.TopicRepo.Delete(topic)
	return nil
}
