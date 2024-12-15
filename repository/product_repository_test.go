package repository_test

import (
	"reflect"

	"testing"

	"github.com/thesayedirfan/kart-challenge/domain"
	"github.com/thesayedirfan/kart-challenge/errors"
	"github.com/thesayedirfan/kart-challenge/repository"
)

func TestProductRepository_GetAll(t *testing.T) {
	repo := repository.NewProductRepository()
	products := repo.GetAll()
	expectedProducts := []domain.Product{
		{ID: "1", Name: "Chicken Waffle", Price: 10, Category: "Waffle"},
	}
	if !reflect.DeepEqual(products, expectedProducts) {
		t.Errorf("expected %v, got %v", expectedProducts, products)
	}
}

func TestProductRepository_GetByID_Found(t *testing.T) {
	repo := repository.NewProductRepository()
	product, err := repo.GetByID("1")
	expectedProduct := &domain.Product{ID: "1", Name: "Chicken Waffle", Price: 10, Category: "Waffle"}

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(product, expectedProduct) {
		t.Errorf("expected %v, got %v", expectedProduct, product)
	}
}

func TestProductRepository_GetByID_NotFound(t *testing.T) {
	repo := repository.NewProductRepository()
	product, err := repo.GetByID("2")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err != errors.ErrorProductNotFound {
		t.Errorf("expected error %v, got %v", errors.ErrorProductNotFound, err)
	}
	if product != nil {
		t.Errorf("expected nil product, got %v", product)
	}
}
