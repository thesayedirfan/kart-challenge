package usecase

import (
	"github.com/thesayedirfan/kart-challenge/domain"
	"github.com/thesayedirfan/kart-challenge/errors"
	"github.com/thesayedirfan/kart-challenge/repository"
	"github.com/thesayedirfan/kart-challenge/utils"
)

type OrderUseCase struct {
	orderRepository   *repository.OrderRepository
	productRepository *repository.ProductRepository
}

func (o *OrderUseCase) PlaceOrder(items []domain.Item, couponCode string) (*domain.Order, error) {
	var total float64
	var discount float64
	var products []domain.Product

	for _, item := range items {
		product, err := o.productRepository.GetByID(item.ProductID)
		if err != nil {
			return nil, errors.ErrorProductNotFound
		}
		total += float64(product.Price) * float64(item.Quantity)
		products = append(products, *product)
	}

	if couponCode != "" {
		ok, err := utils.ValidateDiscountCoupon(couponCode, o.orderRepository.Coupons)
		if err != nil {
			return nil, err
		}
		discount = utils.GetDiscountedPrice(total, 10)
		if ok {
			total -= discount
		}
	}

	order := domain.Order{
		Total:    total,
		Discount: discount,
		Items:    items,
		Products: products,
	}
	savedOrder := o.orderRepository.Save(order)
	return &savedOrder, nil
}

func NewOrderUseCase(orderRepo *repository.OrderRepository, productRepo *repository.ProductRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepository:   orderRepo,
		productRepository: productRepo,
	}
}
