package services

import (
	"context"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/aggregates"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/constants"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/entities"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/repository"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/valueobjects"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsService interface {
	LookingForDocuments(ctx context.Context, pagination *valueobjects.Filters) (*aggregates.ProductsAggregate, error)
	LookingForDocument(ctx context.Context, id primitive.ObjectID) (*entities.Product, error)
}

type productsServiceImpl struct {
	repo repository.ProductsRepository
	log  logger.Logger
}

func New(repo repository.ProductsRepository, log logger.Logger) ProductsService {
	return &productsServiceImpl{
		repo: repo,
		log:  log,
	}
}

func (s *productsServiceImpl) LookingForDocuments(ctx context.Context, pagination *valueobjects.Filters) (*aggregates.ProductsAggregate, error) {
	s.log.Debug("start LookingForDocuments service")

	s.log.Info("validating max limit")
	if pagination.Limit > constants.MaxLimit || pagination.Limit <= 0 {
		s.log.Debugf("replacing limit %d by max limit %d", pagination.Limit, constants.MaxLimit)
		pagination.Limit = constants.MaxLimit
	}

	s.log.Debug("looking for documents")
	products, err := s.repo.FindAllWithPagination(ctx, pagination)
	if err != nil {
		s.log.Errorf("error find by documents: %v", err)
		return nil, err
	}

	s.log.Debug("building aggregate")
	agg := &aggregates.ProductsAggregate{Products: products, Pagination: pagination}

	s.log.Info("returning documents")
	return agg, nil
}

func (s *productsServiceImpl) LookingForDocument(ctx context.Context, id primitive.ObjectID) (*entities.Product, error) {
	s.log.Debugf("looking for document with id: %v", id)

	product, err := s.repo.FindById(ctx, id)
	if err != nil {
		s.log.Errorf("error retrieving document: %v", err)
		return nil, err
	}

	s.log.Info("document found")
	return product, nil
}
