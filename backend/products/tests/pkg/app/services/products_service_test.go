package services_test

import (
	"context"
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/app/services"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/entities"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/valueobjects"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindAllWithPagination(ctx context.Context, pagination *valueobjects.Filters) ([]*entities.Product, error) {
	args := m.Called(ctx, pagination)
	return args.Get(0).([]*entities.Product), args.Error(1)
}

func (m *MockRepository) FindById(ctx context.Context, id primitive.ObjectID) (*entities.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Product), args.Error(1)
}

var _ = Describe("Services", func() {
	var mockRepository *MockRepository
	var log logger.Logger
	var ctx context.Context

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
		mockRepository = new(MockRepository)
		log = logger.New(logger.Opts{
			LogLevel: "debug",
			AppName:  "products-srv-test",
			Tags: map[string]string{
				"env":   "test",
				"owner": "unit_tests",
			},
		})

		ctx = context.Background()
	})

	Context("find documents", func() {
		It("can return one documents", func() {
			pagination := &valueobjects.Filters{Limit: 10}

			mockRepository.
				On("FindAllWithPagination", ctx, pagination).
				Times(1).
				Return([]*entities.Product{element}, nil)

			srv := services.New(mockRepository, log)

			documents, err := srv.LookingForDocuments(ctx, pagination)
			Expect(err).NotTo(HaveOccurred())
			Expect(documents).NotTo(BeNil())
			Expect(documents).To(HaveExistingField("Products"))
			Expect(documents).To(HaveExistingField("Pagination"))
			Expect(documents.Products).To(HaveLen(1))
		})

		When("limit pagination is exceeded", func() {
			It("can return one documents", func() {
				pagination := &valueobjects.Filters{Limit: 100}

				mockRepository.
					On("FindAllWithPagination", ctx, pagination).
					Times(1).
					Return([]*entities.Product{element}, nil)

				srv := services.New(mockRepository, log)

				documents, err := srv.LookingForDocuments(ctx, pagination)
				Expect(err).NotTo(HaveOccurred())
				Expect(documents).NotTo(BeNil())
				Expect(documents).To(HaveExistingField("Products"))
				Expect(documents).To(HaveExistingField("Pagination"))
				Expect(documents.Products).To(HaveLen(1))
			})
		})

		When("occurs an error", func() {
			It("can process the error", func() {
				pagination := &valueobjects.Filters{Limit: 10}

				mockRepository.
					On("FindAllWithPagination", ctx, pagination).
					Times(1).
					Return([]*entities.Product{}, errors.New("generic error occurs"))

				srv := services.New(mockRepository, log)

				documents, err := srv.LookingForDocuments(ctx, pagination)
				Expect(documents).To(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("find document by id", func() {
		When("produces a success response", func() {
			It("can return a product document", func() {
				mockRepository.
					On("FindById", ctx, element.Id).
					Times(1).
					Return(element, nil)

				srv := services.New(mockRepository, log)

				product, err := srv.LookingForDocument(ctx, element.Id)
				Expect(err).NotTo(HaveOccurred())
				Expect(product).NotTo(BeNil())
				Expect(product).To(HaveExistingField("Id"))
				Expect(product).To(Equal(element))
			})
		})

		When("produces an error", func() {
			It("return an error", func() {
				var empty *entities.Product
				mockRepository.
					On("FindById", ctx, element.Id).
					Times(1).
					Return(empty, errors.New("generic error"))

				srv := services.New(mockRepository, log)

				product, err := srv.LookingForDocument(ctx, element.Id)
				Expect(product).To(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
