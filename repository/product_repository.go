package repository

import (
	"github.com/thesayedirfan/kart-challenge/domain"
	"github.com/thesayedirfan/kart-challenge/errors"
)

type ProductRepository struct {
	products map[string]domain.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: map[string]domain.Product{
			"1": {ID: "1", Name: "Chicken Waffle", Price: 10, Category: "Waffle"},
		},
	}
}

func (p *ProductRepository) GetAll() []domain.Product {

	var products []domain.Product

	for _, product := range p.products {
		products = append(products, product)
	}

	return products
}

func (p *ProductRepository) GetByID(id string) (*domain.Product, error) {
	product, exists := p.products[id]
	if !exists {
		return nil, errors.ErrorProductNotFound
	}
	return &product, nil
}
