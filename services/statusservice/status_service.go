package statusservice

import (
	"errors"

	"github.com/greglownes/gogin1/models"
	"github.com/greglownes/gogin1/repositories"
)

type StatusServiceInterface interface {
	GetAll() ([]models.Status, error)
	GetByID(id uint) (*models.Status, error)
	Create(*models.Status) error
	Update(*models.Status) error
	Delete(*models.Status) error
}

// hold repos and other dependencies
type statusService struct {
	StatusRepo repositories.StatusRepoInterface
}

func NewStatusService(tr repositories.StatusRepoInterface) StatusServiceInterface {
	return &statusService{
		StatusRepo: tr,
	}
}

func (ts *statusService) GetAll() ([]models.Status, error) {
	statuses, err := ts.StatusRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return statuses, nil
}

func (ts *statusService) GetByID(id uint) (*models.Status, error) {
	if id == 0 {
		return nil, errors.New("id param is required")
	}
	status, err := ts.StatusRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (ts *statusService) Create(status *models.Status) error {
	return ts.StatusRepo.Create(status)
}

func (ts *statusService) Update(status *models.Status) error {
	return ts.StatusRepo.Update(status)
}

func (ts *statusService) Delete(status *models.Status) error {
	ts.StatusRepo.Delete(status)
	return nil
}
