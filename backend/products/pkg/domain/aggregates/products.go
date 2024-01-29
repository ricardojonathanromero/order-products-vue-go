package aggregates

import (
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/entities"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/valueobjects"
)

type ProductsAggregate struct {
	Products   []*entities.Product   `json:"products"`
	Pagination *valueobjects.Filters `json:"pagination"`
}
