package repositories

import (
	"github.com/greglownes/gogin1/models"
	"gorm.io/gorm"
)

type ProductRepoInterface interface {
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(product *models.Product)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepoInterface {
	return &productRepo{
		db: db,
	}
}

func (t *productRepo) GetAll() ([]models.Product, error) {
	var products []models.Product
	t.db.Find(&products)
	return products, nil
}

func (t *productRepo) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := t.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (t *productRepo) Create(product *models.Product) error {
	// spew.Dump(product)
	return t.db.Create(product).Error

	// hold := &models.Product{
	// 	Title: "title99a",
	// 	Price: decimal.NewFromFloat(99.99),
	// 	Status: models.Status{
	// 		Status: "status99a",
	// 	},
	// }
	// spew.Dump(hold)
	// return t.db.Create(hold)).Error

	// return t.db.Create(
	//   &User{
	//     Name: "jinzhu",
	//     CreditCard: CreditCard{
	//       Number: "411111111111"
	//     }
	//   }
	// )

}

func (t *productRepo) Update(product *models.Product) error {
	return t.db.Save(product).Error
}

func (t *productRepo) Delete(product *models.Product) {
	t.db.Delete(product)
}
