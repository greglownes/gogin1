package productservice

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/greglownes/gogin1/models"
	"github.com/greglownes/gogin1/repositories"
)

type ProductServiceInterface interface {
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Create(*models.Product) error
	Update(*models.Product) error
	Delete(*models.Product) error

	CreateModelForAddFromRawInput(c *gin.Context) (models.Product, error)
	MapProductToProductOutput(product *models.Product) models.ProductOutput
}

// hold repos and other dependencies
type productService struct {
	ProductRepo repositories.ProductRepoInterface
}

func NewProductService(repo repositories.ProductRepoInterface) ProductServiceInterface {
	return &productService{
		ProductRepo: repo,
	}
}

func (service *productService) GetAll() ([]models.Product, error) {
	products, err := service.ProductRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (service *productService) GetByID(id uint) (*models.Product, error) {
	if id == 0 {
		return nil, errors.New("id param is required")
	}
	product, err := service.ProductRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (service *productService) Create(product *models.Product) error {
	spew.Dump(product)
	hold := service.ProductRepo.Create(product)
	spew.Dump(hold)
	return hold
}

func (service *productService) Update(product *models.Product) error {
	return service.ProductRepo.Update(product)
}

func (service *productService) Delete(product *models.Product) error {
	service.ProductRepo.Delete(product)
	return nil
}

func (service *productService) CreateModelForAddFromRawInput(c *gin.Context) (models.Product, error) {
	// validate input
	var rawInput models.ProductCreateInput
	if err := c.ShouldBindJSON(&rawInput); err != nil {
		spew.Dump(err)
		return models.Product{}, err
	}
	spew.Dump(rawInput)

	// convert raw input to model
	// skip ID and 3 date fields, that is handled by gorm
	product := models.Product{
		Title:  rawInput.Title,
		Price:  rawInput.Price,
		Status: rawInput.Status,
		// StatusID int
	}
	spew.Dump(product)
	return product, nil
}

func (service *productService) MapProductToProductOutput(product *models.Product) models.ProductOutput {
	return models.ProductOutput{
		ID:    product.ID,
		Title: product.Title,
		Price: product.Price,
		//Status: product.Status,
	}
}
