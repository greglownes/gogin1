package repositories

import (
	"github.com/greglownes/gogin1/models"

	"gorm.io/gorm"
)

type TopicRepoInterface interface {
	GetAll() ([]models.Topic, error)
	GetByID(id uint) (*models.Topic, error)
	Create(topic *models.Topic) error
	Update(topic *models.Topic) error
	Delete(topic *models.Topic)
}

type topicRepo struct {
	db *gorm.DB
}

func NewTopicRepo(db *gorm.DB) TopicRepoInterface {
	return &topicRepo{
		db: db,
	}
}

func (t *topicRepo) GetAll() ([]models.Topic, error) {
	var topics []models.Topic
	t.db.Find(&topics)
	// spew.Dump(topics)
	return topics, nil
}

func (t *topicRepo) GetByID(id uint) (*models.Topic, error) {
	var topic models.Topic
	if err := t.db.First(&topic, id).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func (t *topicRepo) Create(topic *models.Topic) error {
	return t.db.Create(topic).Error
}

func (t *topicRepo) Update(topic *models.Topic) error {
	return t.db.Save(topic).Error
}

func (t *topicRepo) Delete(topic *models.Topic) {
	t.db.Delete(topic)
}
