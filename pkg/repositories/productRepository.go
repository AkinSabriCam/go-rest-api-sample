package repositories

import (
	"github.com/google/uuid"
	"go-rest-api-sample/pkg/models"
	"gorm.io/gorm"
)

var ProductRepositoryInstant productRepository

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) {
	ProductRepositoryInstant = productRepository{
		db: db,
	}
}

func (repo productRepository) GetAll() []models.Product {
	var products []models.Product
	repo.db.Find(&products)

	return products
}

func (repo productRepository) Create(product models.Product) {
	repo.db.Create(&product)
}

func (repo productRepository) GetById(id uuid.UUID) models.Product {
	var product models.Product
	repo.db.Find(&product, id)

	return product
}

func (repo productRepository) Delete(id uuid.UUID) {

	var product models.Product
	repo.db.Find(&product, id)
	repo.db.Delete(&product)
}
