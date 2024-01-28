package services

import "github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"

type ProductsService interface {
}

type productsServiceImpl struct {
	log logger.Logger
}
