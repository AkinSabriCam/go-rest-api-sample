package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id          uuid.UUID `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"column:name"`
	Price       float64   `gorm:"column:price"`
	CreatedDate time.Time `gorm:"column:createdDate"`
}

func NewProduct(name string, price float64, releaseDate time.Time) Product {
	return Product{
		Name:        name,
		Price:       price,
		CreatedDate: time.Now(),
	}
}
