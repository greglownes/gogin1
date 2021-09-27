package repositories

import (
	"github.com/greglownes/gogin1/models"
	"gorm.io/gorm"
)

type StatusRepoInterface interface {
	GetAll() ([]models.Status, error)
	GetByID(id uint) (*models.Status, error)
	Create(status *models.Status) error
	Update(status *models.Status) error
	Delete(status *models.Status)
}

type statusRepo struct {
	db *gorm.DB
}

func NewStatusRepo(db *gorm.DB) StatusRepoInterface {
	return &statusRepo{
		db: db,
	}
}

func (t *statusRepo) GetAll() ([]models.Status, error) {
	var statuses []models.Status
	t.db.Find(&statuses)
	return statuses, nil
}

func (t *statusRepo) GetByID(id uint) (*models.Status, error) {
	var status models.Status
	if err := t.db.First(&status, id).Error; err != nil {
		return nil, err
	}
	return &status, nil
}

func (t *statusRepo) Create(status *models.Status) error {
	return t.db.Create(status).Error
}

func (t *statusRepo) Update(status *models.Status) error {
	return t.db.Save(status).Error
}

func (t *statusRepo) Delete(status *models.Status) {
	t.db.Delete(status)
}
