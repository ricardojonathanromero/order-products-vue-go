package http

type Method string

const (
	MethodGet    Method = "GET"
	MethodPost   Method = "POST"
	MethodPut    Method = "PUT"
	MethodPatch  Method = "PATCH"
	MethodHead   Method = "HEAD"
	MethodDelete Method = "DELETE"
)

var defaultHeaders = map[string]string{
	"User-Agent":   "go github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/clients/http",
	"Content-Type": "application/json",
}
