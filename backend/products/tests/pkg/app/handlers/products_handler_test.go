package handlers_test

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/app/handlers"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/aggregates"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/entities"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/errors"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/valueobjects"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/validator"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) LookingForDocuments(ctx context.Context, pagination *valueobjects.Filters) (*aggregates.ProductsAggregate, error) {
	args := m.Called(ctx, pagination)
	return args.Get(0).(*aggregates.ProductsAggregate), args.Error(1)
}

func (m *MockService) LookingForDocument(ctx context.Context, id primitive.ObjectID) (*entities.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Product), args.Error(1)
}

var _ = Describe("Handlers", func() {
	var mockService *MockService
	var log logger.Logger
	var e *echo.Echo

	element := &entities.Product{
		Id:          primitive.NewObjectID(),
		Sku:         "1234567890",
		Price:       19.99,
		Name:        "product_example",
		Description: "product example",
		Quantity:    1,
		Image:       "https://mock.image.com",
		Category:    "test",
	}

	BeforeEach(func() {
		mockService = new(MockService)
		log = logger.New(logger.Opts{
			LogLevel: "debug",
			AppName:  "products-handler-test",
			Tags: map[string]string{
				"env":   "test",
				"owner": "unit_tests",
			},
		})

		e = echo.New()
		e.Validator = validator.NewValidator()
	})

	Context("consume the handleListProducts endpoint", func() {
		When("return success response", func() {
			It("can return a list of products with the pagination applied", func() {
				// declare echo request
				q := make(url.Values)
				q.Set("limit", "10")
				req := httptest.NewRequest(http.MethodGet, "/v1/products?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				mockService.On("LookingForDocuments", mock.Anything, mock.Anything).
					Times(1).
					Return(&aggregates.ProductsAggregate{
						Products:   []*entities.Product{element},
						Pagination: &valueobjects.Filters{Limit: 10},
					}, nil)

				hdl := handlers.New(mockService, log)

				err := hdl.HandleListProducts(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusOK))

				// unmarshal response in aggregate
				var result aggregates.ProductsAggregate
				err = json.Unmarshal(rec.Body.Bytes(), &result)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(HaveExistingField("Products"))
				Expect(result).To(HaveExistingField("Pagination"))
				Expect(result.Products).NotTo(BeNil())
				Expect(result.Pagination).NotTo(BeNil())
				Expect(result.Products).To(HaveLen(1))
				Expect(result.Pagination.Limit).To(Equal(int64(10)))
			})
		})

		When("sent invalid request", func() {
			It("produces an http bad request caused by binding query params", func() {
				// declare echo request
				req := httptest.NewRequest(http.MethodGet, "/v1/products?limit=invalid&offset=invalid", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				hdl := handlers.New(mockService, log)

				err := hdl.HandleListProducts(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
			})

			It("produces an http bad request by limit out of range", func() {
				// declare echo request
				req := httptest.NewRequest(http.MethodGet, "/v1/products?limit=200", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				hdl := handlers.New(mockService, log)

				err := hdl.HandleListProducts(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
			})
		})

		When("service layer produces an error", func() {
			It("return an invalid http code", func() {
				// declare echo request
				q := make(url.Values)
				q.Set("limit", "10")
				req := httptest.NewRequest(http.MethodGet, "/v1/products?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				var empty *aggregates.ProductsAggregate

				mockService.On("LookingForDocuments", mock.Anything, mock.Anything).
					Times(1).
					Return(empty, errors.NewError(errors.InvalidReqBind))

				hdl := handlers.New(mockService, log)

				err := hdl.HandleListProducts(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusConflict))
			})
		})
	})

	Context("consume the handleProduct endpoint", func() {
		When("return success response", func() {
			It("can return the single response", func() {
				req := httptest.NewRequest(http.MethodGet, "/v1/products", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetParamNames("productId")
				c.SetParamValues(element.Id.Hex())

				mockService.On("LookingForDocument", mock.Anything, element.Id).
					Times(1).
					Return(element, nil)

				hdl := handlers.New(mockService, log)

				err := hdl.HandleProduct(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusOK))

				// unmarshal response in aggregate
				var result *entities.Product
				err = json.Unmarshal(rec.Body.Bytes(), &result)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(element))
			})
		})

		When("sent invalid request", func() {
			It("produces an http bad request caused by binding invalid id value", func() {
				// declare echo request
				req := httptest.NewRequest(http.MethodGet, "/v1/products", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetParamNames("productId")
				c.SetParamValues("some-value")

				hdl := handlers.New(mockService, log)

				err := hdl.HandleProduct(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
			})
		})

		When("service layer produces an error", func() {
			It("cause an invalid http error", func() {
				var empty *entities.Product
				req := httptest.NewRequest(http.MethodGet, "/v1/products", nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetParamNames("productId")
				c.SetParamValues(element.Id.Hex())

				mockService.On("LookingForDocument", mock.Anything, element.Id).
					Times(1).
					Return(empty, errors.NewError(errors.ErrNoDocument))

				hdl := handlers.New(mockService, log)

				err := hdl.HandleProduct(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusConflict))
			})
		})
	})
})
