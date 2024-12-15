package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thesayedirfan/kart-challenge/errors"
	"github.com/thesayedirfan/kart-challenge/usecase"
	"github.com/thesayedirfan/kart-challenge/utils"
)

type ProductAPI struct {
	productUsecase *usecase.ProductUseCase
}

func NewProductAPI(productUseCase *usecase.ProductUseCase) *ProductAPI {
	return &ProductAPI{productUsecase: productUseCase}
}

func (h *ProductAPI) ListProducts(c *gin.Context) {
	products := h.productUsecase.ListProducts()
	c.JSON(http.StatusOK, products)
}

func (h *ProductAPI) GetProduct(c *gin.Context) {

	id := c.Param("id")

	if !utils.IsValidProductID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"message": errors.ErrorInvalidIDSupplied.Error()})
		return
	}

	product, err := h.productUsecase.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
