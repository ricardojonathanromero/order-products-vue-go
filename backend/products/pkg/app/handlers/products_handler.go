package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/app/services"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/errors"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/valueobjects"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ProductsHandler interface {
	HandleListProducts(c echo.Context) error
	HandleProduct(c echo.Context) error
}

type productsHandlerImpl struct {
	srv services.ProductsService
	log logger.Logger
}

func New(srv services.ProductsService, log logger.Logger) ProductsHandler {
	return &productsHandlerImpl{
		srv: srv,
		log: log,
	}
}

// HandleListProducts godoc
// @Summary List products
// @Description  list all the products existing in the catalog
// @Tags         catalogs
// @Accept       json
// @Produce      json
// @Param		 Authorization	header	string				true	"Bearer token"
// @Param        limit	query	int	false "number of elements to return"
// @Param        offset	query	int	false "number of elements to skip"
// @Param        filter_by	query	string	false "field value used to filter. e.g. category=soda"
// @Param        sort_by	query	string	false "indicates fields to sort by, by default id asc. e.g. -category,-_id"
// @Success      200  {object}  aggregates.ProductsAggregate
// @Failure      400  {object}  errors.CustomError
// @Failure      409  {object}  errors.CustomError
// @Failure      403  {object}  errors.CustomError
// @Failure      424  {object}  errors.CustomError
// @Failure      500  {object}  errors.CustomError
// @Security		ApiKeyAuth
// @Router       / [get]
func (h *productsHandlerImpl) HandleListProducts(c echo.Context) error {
	h.log.Debug("start handle list products")
	ctx := c.Request().Context()

	var pagination valueobjects.Filters

	h.log.Debug("binding request")
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &pagination); err != nil {
		h.log.Errorf("error binding query parameters: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewError(errors.InvalidReqBind, err))
	}

	h.log.Debug("validating request")
	if err := c.Validate(pagination); err != nil {
		h.log.Errorf("error invalid fields: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewError(errors.InvalidReqBind, err))
	}

	h.log.Info("processing look for documents trx")
	products, err := h.srv.LookingForDocuments(ctx, &pagination)
	if err != nil {
		h.log.Errorf("error retrieveing documents: %v", err)
		return c.JSON(http.StatusConflict, err)
	}

	h.log.Info("products retrieved correctly")
	return c.JSON(http.StatusOK, products)
}

// HandleProduct godoc
// @Summary Return product information
// @Description  return the product information filtering by product id
// @Tags         catalogs
// @Accept       json
// @Produce      json
// @Param		 Authorization	header	string				true	"Bearer token"
// @Param        id		path	string	true	"product ID" Format(primitive.ObjectID)
// @Success      200  {object}  aggregates.ProductsAggregate
// @Failure      400  {object}  errors.CustomError
// @Failure      409  {object}  errors.CustomError
// @Failure      403  {object}  errors.CustomError
// @Failure      424  {object}  errors.CustomError
// @Failure      500  {object}  errors.CustomError
// @Security		ApiKeyAuth
// @Router       /{id} [get]
func (h *productsHandlerImpl) HandleProduct(c echo.Context) error {
	h.log.Debug("start handle product by id")
	ctx := c.Request().Context()

	id := c.Param("productId")

	productId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		h.log.Errorf("error invalid product id: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewError(errors.InvalidId, err))
	}

	h.log.Info("processing look for document trx")
	product, err := h.srv.LookingForDocument(ctx, productId)
	if err != nil {
		h.log.Errorf("error retrieveing document: %v", err)
		return c.JSON(http.StatusConflict, err)
	}

	h.log.Info("product retrieved correctly")
	return c.JSON(http.StatusOK, product)
}
