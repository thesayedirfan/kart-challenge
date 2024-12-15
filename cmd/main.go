package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thesayedirfan/kart-challenge/api"
	"github.com/thesayedirfan/kart-challenge/middleware"
	"github.com/thesayedirfan/kart-challenge/repository"
	"github.com/thesayedirfan/kart-challenge/usecase"
	"github.com/thesayedirfan/kart-challenge/utils"
)

func main() {

	couponCountInFiles := utils.Process()

	err := os.Setenv("API_KEY", "fca87a7f")
	if err != nil {
		log.Fatalf("Error setting environment variable: %v", err)
	}

	router := gin.Default()

	router.Use(middleware.APIMiddleware())

	productRepository := repository.NewProductRepository()
	orderRepository := repository.NewOrderRepository(couponCountInFiles)

	productUseCase := usecase.NewProductUsecase(productRepository)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, productRepository)

	productAPIs := api.NewProductAPI(productUseCase)
	orderAPIs := api.NewOrderAPI(orderUseCase)

	router.GET("/product", productAPIs.ListProducts)
	router.GET("/product/:id", productAPIs.GetProduct)

	router.POST("/order", orderAPIs.PlaceOrder)

	router.Run(":8080")

}
