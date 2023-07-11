package repositories

import (
	"github.com/google/uuid"
	"go-rest-api-sample/pkg/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db gorm.DB
}

func NewRepository(db gorm.DB) ProductRepository {
	return ProductRepository{
		db: db,
	}
}

func (repo ProductRepository) GetAll() []models.Product {
	var products []models.Product
	repo.db.Find(&products)

	return products
}

func (repo ProductRepository) Create(product models.Product) {
	repo.db.Create(&product)
}

func (repo ProductRepository) GetById(id uuid.UUID) models.Product {
	var product models.Product
	repo.db.Find(&product, id)

	return product
}

func (repo ProductRepository) Delete(id uuid.UUID) {

	var product models.Product
	repo.db.Find(&product, id)
	repo.db.Delete(&product)
}
