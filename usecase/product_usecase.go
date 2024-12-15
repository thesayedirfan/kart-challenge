package usecase

import (
	"github.com/thesayedirfan/kart-challenge/domain"
	"github.com/thesayedirfan/kart-challenge/repository"
)

type ProductUseCase struct {
	productRepo *repository.ProductRepository
}

func NewProductUsecase(productRepo *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepo: productRepo}
}

func (p *ProductUseCase) GetProductByID(id string) (*domain.Product, error) {
	return p.productRepo.GetByID(id)
}

func (p *ProductUseCase) ListProducts() []domain.Product {
	return p.productRepo.GetAll()
}
