package repository

import (
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepository interface {
}

type productsRepositoryImpl struct {
	client *mongo.Client
	log    logger.Logger
}

func New(client *mongo.Client, log logger.Logger) ProductsRepository {
	return &productsRepositoryImpl{client: client, log: log}
}
