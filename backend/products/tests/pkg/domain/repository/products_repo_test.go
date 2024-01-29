package repository_test

import (
	"context"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/repository"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/valueobjects"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var documents = []bson.D{
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "1"}, {Key: "price", Value: 23.34}, {Key: "name", Value: "test_1"}, {Key: "description", Value: "test_test_1"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "2"}, {Key: "price", Value: 10.23}, {Key: "name", Value: "test_2"}, {Key: "description", Value: "test_test_2"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "3"}, {Key: "price", Value: 232.12}, {Key: "name", Value: "test_3"}, {Key: "description", Value: "test_test_3"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "4"}, {Key: "price", Value: 10.00}, {Key: "name", Value: "test_4"}, {Key: "description", Value: "test_test_4"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "5"}, {Key: "price", Value: 23.41}, {Key: "name", Value: "test_5"}, {Key: "description", Value: "test_test_5"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "6"}, {Key: "price", Value: 78.29}, {Key: "name", Value: "test_6"}, {Key: "description", Value: "test_test_6"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "7"}, {Key: "price", Value: 54.90}, {Key: "name", Value: "test_7"}, {Key: "description", Value: "test_test_7"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "8"}, {Key: "price", Value: 65.76}, {Key: "name", Value: "test_8"}, {Key: "description", Value: "test_test_8"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "9"}, {Key: "price", Value: 54.34}, {Key: "name", Value: "test_9"}, {Key: "description", Value: "test_test_9"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: "10"}, {Key: "price", Value: 321.09}, {Key: "name", Value: "test_10"}, {Key: "description", Value: "test_test_10"}, {Key: "quantity", Value: 10}, {Key: "image", Value: "https://mock.image.com"}, {Key: "category", Value: "test"}},
}

var doesNotMatch = []bson.D{
	{{Key: "_id", Value: primitive.NewObjectID()}, {Key: "sku", Value: 1}, {Key: "price", Value: "23.34"}, {Key: "name", Value: "test_1"}, {Key: "description", Value: "test_test_1"}, {Key: "quantity", Value: "10"}},
}

var _ = Describe("Repositories", func() {
	var mongoTest *mtest.T
	var log logger.Logger

	BeforeEach(func() {
		mongoTest = mtest.New(globalTest, mtest.NewOptions().ClientType(mtest.Mock))

		log = logger.New(logger.Opts{
			LogLevel: "debug",
			AppName:  "products-repo-test",
			Tags:     map[string]string{"env": "test"},
		})
	})

	Describe("find documents", func() {
		Context("with mongodb mock server", func() {
			When("pagination is used", func() {
				It("can read documents", func() {
					mongoTest.Run("with limit", func(mt *mtest.T) {
						// reply documents
						ns := fmt.Sprintf("%s.%s", mt.DB.Name(), mt.Coll.Name())
						//mt.AddMockResponses(bson.D{{"ok", 1}, {ns, true}, {"n", 10}})
						for i, document := range documents {
							if i == 0 {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.FirstBatch, document))
							} else {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.NextBatch, document))
							}
						}
						mt.AddMockResponses(mtest.CreateCursorResponse(0, ns, mtest.NextBatch))

						repo := repository.New(mt.Coll, log)
						pagination := &valueobjects.Filters{
							Limit:  10,
							Offset: 0,
						}

						docs, err := repo.
							FindAllWithPagination(context.Background(), pagination)

						Expect(err).NotTo(HaveOccurred())
						Expect(docs).To(HaveLen(10))
					})

					mongoTest.Run("with limit and skip", func(mt *mtest.T) {
						// reply documents
						ns := fmt.Sprintf("%s.%s", mt.DB.Name(), mt.Coll.Name())
						//mt.AddMockResponses(bson.D{{"ok", 1}, {ns, true}, {"n", 10}})
						for i, document := range documents[:5] {
							if i == 0 {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.FirstBatch, document))
							} else {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.NextBatch, document))
							}
						}
						mt.AddMockResponses(mtest.CreateCursorResponse(0, ns, mtest.NextBatch))

						repo := repository.New(mt.Coll, log)
						pagination := &valueobjects.Filters{
							Limit:  5,
							Offset: 5,
						}

						docs, err := repo.
							FindAllWithPagination(context.Background(), pagination)

						Expect(err).NotTo(HaveOccurred())
						Expect(docs).To(HaveLen(5))
					})

					mongoTest.Run("with pagination configured", func(mt *mtest.T) {
						// reply documents
						ns := fmt.Sprintf("%s.%s", mt.DB.Name(), mt.Coll.Name())
						//mt.AddMockResponses(bson.D{{"ok", 1}, {ns, true}, {"n", 10}})
						for i, document := range documents {
							if i == 0 {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.FirstBatch, document))
							} else {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.NextBatch, document))
							}
						}
						mt.AddMockResponses(mtest.CreateCursorResponse(0, ns, mtest.NextBatch))

						repo := repository.New(mt.Coll, log)
						pagination := &valueobjects.Filters{
							Limit:    10,
							Offset:   10,
							FilterBy: "category=test",
							SortBy:   "_id",
						}

						docs, err := repo.
							FindAllWithPagination(context.Background(), pagination)

						Expect(err).NotTo(HaveOccurred())
						Expect(docs).To(HaveLen(10))
					})

					mongoTest.Run("with multi values in pagination", func(mt *mtest.T) {
						// reply documents
						ns := fmt.Sprintf("%s.%s", mt.DB.Name(), mt.Coll.Name())
						//mt.AddMockResponses(bson.D{{"ok", 1}, {ns, true}, {"n", 10}})
						for i, document := range documents {
							if i == 0 {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.FirstBatch, document))
							} else {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.NextBatch, document))
							}
						}
						mt.AddMockResponses(mtest.CreateCursorResponse(0, ns, mtest.NextBatch))

						repo := repository.New(mt.Coll, log)
						pagination := &valueobjects.Filters{
							Limit:    10,
							Offset:   10,
							FilterBy: "category=test example",
							SortBy:   "-_id",
						}

						docs, err := repo.
							FindAllWithPagination(context.Background(), pagination)

						Expect(err).NotTo(HaveOccurred())
						Expect(docs).To(HaveLen(10))
					})

					mongoTest.Run("without pagination", func(mt *mtest.T) {
						// reply documents
						ns := fmt.Sprintf("%s.%s", mt.DB.Name(), mt.Coll.Name())
						//mt.AddMockResponses(bson.D{{"ok", 1}, {ns, true}, {"n", 10}})
						for i, document := range documents[:1] {
							if i == 0 {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.FirstBatch, document))
							} else {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.NextBatch, document))
							}
						}
						mt.AddMockResponses(mtest.CreateCursorResponse(0, ns, mtest.NextBatch))

						repo := repository.New(mt.Coll, log)

						docs, err := repo.
							FindAllWithPagination(context.Background(), nil)

						Expect(err).NotTo(HaveOccurred())
						Expect(docs).To(HaveLen(1))
					})
				})

				It("cannot return documents", func() {
					mongoTest.Run("invalid db response", func(mt *mtest.T) {
						mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
							Index:   1,
							Code:    11000,
							Message: "invalid connection",
						}))

						repo := repository.New(mt.Coll, log)

						docs, err := repo.
							FindAllWithPagination(context.Background(), nil)

						Expect(docs).To(BeNil())
						Expect(err).To(HaveOccurred())
					})

					mongoTest.Run("structure does not match", func(mt *mtest.T) {
						// reply documents
						ns := fmt.Sprintf("%s.%s", mt.DB.Name(), mt.Coll.Name())
						//mt.AddMockResponses(bson.D{{"ok", 1}, {ns, true}, {"n", 10}})
						for i, document := range doesNotMatch {
							if i == 0 {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.FirstBatch, document))
							} else {
								mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.NextBatch, document))
							}
						}
						mt.AddMockResponses(mtest.CreateCursorResponse(0, ns, mtest.NextBatch))

						repo := repository.New(mt.Coll, log)

						docs, err := repo.
							FindAllWithPagination(context.Background(), nil)

						Expect(docs).To(BeNil())
						Expect(err).To(HaveOccurred())
					})
				})
			})
		})
	})

	Describe("find document", func() {
		Context("with mongodb mock server", func() {
			It("return result", func() {
				mongoTest.Run("by id", func(mt *mtest.T) {
					// reply document
					ns := fmt.Sprintf("%s.%s", mt.DB.Name(), mt.Coll.Name())
					mt.AddMockResponses(mtest.CreateCursorResponse(1, ns, mtest.FirstBatch, documents[0]))
					mt.AddMockResponses(mtest.CreateCursorResponse(0, ns, mtest.NextBatch))

					repo := repository.New(mt.Coll, log)

					doc, err := repo.FindById(context.Background(), primitive.NewObjectID())

					Expect(err).NotTo(HaveOccurred())
					Expect(doc).NotTo(BeNil())
				})
			})

			It("produces an error", func() {
				mongoTest.Run("by id", func(mt *mtest.T) {
					// reply error
					mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{
						Index:   1,
						Code:    11000,
						Message: "invalid connection",
					}))

					repo := repository.New(mt.Coll, log)

					doc, err := repo.FindById(context.Background(), primitive.NewObjectID())

					Expect(doc).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
