package errors_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/errors"
)

var _ = Describe("Errors", func() {
	Context("create dictionary", func() {
		It("default error", func() {
			e := errors.NewError(9999999, fmt.Errorf("generic error"))
			Expect(e).To(HaveOccurred())
		})
	})
})
