package handlers

import (
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/app/services"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
)

type ProductsHandler interface {
}

type productsHandlerImpl struct {
	srv services.ProductsService
	log logger.Logger
}
