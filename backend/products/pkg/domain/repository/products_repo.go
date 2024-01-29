package repository

import (
	"context"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/entities"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/errors"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/valueobjects"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type ProductsRepository interface {
	FindAllWithPagination(ctx context.Context, pagination *valueobjects.Filters) ([]*entities.Product, error)
	FindById(ctx context.Context, id primitive.ObjectID) (*entities.Product, error)
}

type productsRepositoryImpl struct {
	col *mongo.Collection
	log logger.Logger
}

func New(col *mongo.Collection, log logger.Logger) ProductsRepository {
	return &productsRepositoryImpl{col: col, log: log}
}

func (repo *productsRepositoryImpl) FindAllWithPagination(ctx context.Context, pagination *valueobjects.Filters) ([]*entities.Product, error) {
	var products []*entities.Product
	repo.log.Debug("processing pagination")

	repo.log.Debug("building filter and options")
	filter, opts := buildOptions(pagination)

	repo.log.Debug("querying documents")
	cur, err := repo.col.Find(ctx, filter, opts)
	if err != nil {
		repo.log.Errorf("error retrieving documents from db: %v", err)
		return products, errors.NewError(errors.ErrNoDocuments)
	}

	repo.log.Debug("binding documents in entity")
	err = cur.All(ctx, &products)
	if err != nil {
		repo.log.Errorf("error binding documents: %v", err)
		return products, errors.NewError(errors.ErrBindDBDocuments)
	}

	repo.log.Debug("returning documents")
	return products, nil
}

func (repo *productsRepositoryImpl) FindById(ctx context.Context, id primitive.ObjectID) (*entities.Product, error) {
	var result *entities.Product
	repo.log.Debug("find by id")

	err := repo.col.FindOne(ctx, bson.M{"_id": id}, options.FindOne()).Decode(&result)
	if err != nil {
		repo.log.Errorf("document not found: %v", err)
		return result, errors.NewError(errors.ErrNoDocument)
	}

	repo.log.Debug("document retrieved")
	return result, nil
}

func buildOptions(pagination *valueobjects.Filters) (bson.D, *options.FindOptions) {
	filter := bson.D{}
	opts := options.Find()

	if pagination == nil {
		return filter, opts
	}

	if len(pagination.FilterBy) > 0 {
		// split filters
		filter = buildFilter(strings.Split(pagination.FilterBy, ",")...)
	}

	// set limit
	if pagination.Limit > 0 {
		opts.SetLimit(pagination.Limit)
	}

	// set offset
	if pagination.Offset > 0 {
		opts.SetSkip(pagination.Offset)
	}

	// configure sort
	if pagination.SortBy != "" {
		opts.SetSort(buildSort(strings.Split(pagination.SortBy, ",")...))
	}

	return filter, opts
}

func buildFilter(conditions ...string) bson.D {
	filter := bson.D{}

	for _, condition := range conditions {
		c := strings.Split(condition, "=")
		key := c[0]
		values := strings.Split(c[1], " ")

		if len(values) == 1 {
			filter = append(filter, bson.E{Key: key, Value: values[0]})
			continue
		}

		// build filter
		filter = append(filter, bson.E{Key: key, Value: values})
	}

	return filter
}

func buildSort(sort ...string) any {
	sortBy := make(map[string]int, len(sort))

	for _, s := range sort {
		if s[0] == '-' {
			sortBy[s[1:]] = -1
		} else {
			sortBy[s] = 1
		}
	}

	return sortBy
}
