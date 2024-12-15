package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thesayedirfan/kart-challenge/domain"
	"github.com/thesayedirfan/kart-challenge/usecase"
)

type OrderAPI struct {
	OrderUseCase *usecase.OrderUseCase
}

func NewOrderAPI(orderUsecase *usecase.OrderUseCase) *OrderAPI {
	return &OrderAPI{OrderUseCase: orderUsecase}
}

func (h *OrderAPI) PlaceOrder(c *gin.Context) {
	var orderReq domain.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	order, err := h.OrderUseCase.PlaceOrder(orderReq.Items, orderReq.CouponCode)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
