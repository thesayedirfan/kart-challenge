package repository_test

import (
	"testing"

	"github.com/thesayedirfan/kart-challenge/domain"
	"github.com/thesayedirfan/kart-challenge/repository"
)

func TestOrderRepository_Save(t *testing.T) {
	coupons := map[string]int{}

	repo := repository.NewOrderRepository(coupons)

	order := domain.Order{
		Total:    100.0,
		Discount: 0.0,
		Items: []domain.Item{
			{ProductID: "p1", Quantity: 2},
		},
		Products: []domain.Product{
			{ID: "p1", Name: "Product 1", Price: 50, Category: "Category 1"},
		},
	}

	savedOrder := repo.Save(order)

	if savedOrder.ID == "" {
		t.Fatalf("expected a non-empty ID, got empty ID")
	}

	if len(repo.Orders) != 1 {
		t.Fatalf("expected 1 order in repository, got %d", len(repo.Orders))
	}

	if savedOrder.ID != "1" {
		t.Fatalf("expected order id to be '%s' ", savedOrder.ID)
	}
}
