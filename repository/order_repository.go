package repository

import (
	"fmt"

	"github.com/thesayedirfan/kart-challenge/domain"
)

type OrderRepository struct {
	Orders  map[string]domain.Order
	Coupons map[string]int
}

func (o *OrderRepository) Save(order domain.Order) domain.Order {
	order.ID = fmt.Sprintf("%d", len(o.Orders)+1)
	o.Orders[order.ID] = order
	return order
}

func NewOrderRepository(coupons map[string]int) *OrderRepository {
	return &OrderRepository{
		Orders:  make(map[string]domain.Order),
		Coupons: coupons,
	}
}
