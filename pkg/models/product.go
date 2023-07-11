package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	id          uuid.UUID
	name        string
	price       float64
	createdDate time.Time
}

func NewProduct(name string, price float64, releaseDate time.Time) Product {
	return Product{
		id:          uuid.New(),
		name:        name,
		price:       price,
		createdDate: time.Now(),
	}
}
