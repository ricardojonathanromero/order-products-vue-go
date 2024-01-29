package handlers_test

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/app/handlers"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/entities"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/errors"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/validator"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Login(ctx context.Context, req *entities.LoginReq) (*entities.Token, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*entities.Token), args.Error(1)
}

func (m *MockService) RenewSession(ctx context.Context, req *entities.RefreshTokenReq) (*entities.Token, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*entities.Token), args.Error(1)
}

var _ = Describe("Auth", func() {
	var mockService *MockService
	var log logger.Logger
	var e *echo.Echo

	BeforeEach(func() {
		mockService = new(MockService)
		log = logger.New(logger.Opts{
			LogLevel: "debug",
			AppName:  "auth-handler-test",
			Tags: map[string]string{
				"env":   "test",
				"owner": "unit_tests",
			},
		})

		e = echo.New()
		e.Validator = validator.NewValidator()
	})

	Context("consume the handleLogin endpoint", func() {
		When("return success response", func() {
			It("can return a new session", func() {
				body := `{"username":"user-test","password":"1234567890"}`

				req := httptest.NewRequest(http.MethodPost, "/v1/auth/login", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				mockService.On("Login", mock.Anything, mock.Anything).
					Times(1).
					Return(&entities.Token{
						AccessToken:  "my-access-token",
						RefreshToken: "my-refresh-token",
						Type:         "Bearer",
						ExpiresIn:    3600,
					}, nil)

				hdl := handlers.New(mockService, log)

				err := hdl.HandleLogin(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusOK))

				var result *entities.Token
				err = json.Unmarshal(rec.Body.Bytes(), &result)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(HaveExistingField("AccessToken"))
				Expect(result.AccessToken).To(Equal("my-access-token"))
			})
		})

		When("sent invalid request", func() {
			It("produces an http bad request caused by bind params", func() {
				body := `{"username":"user-test","password":"1234567890"`
				req := httptest.NewRequest(http.MethodPost, "/v1/auth/login", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				hdl := handlers.New(mockService, log)
				err := hdl.HandleLogin(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
			})

			It("produces an http bad request caused by validations", func() {
				body := `{"username":"test","password":"12345"}`
				req := httptest.NewRequest(http.MethodPost, "/v1/auth/login", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				hdl := handlers.New(mockService, log)
				err := hdl.HandleLogin(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
			})
		})

		When("service layer produces an error", func() {
			It("return an invalid http code", func() {
				body := `{"username":"user-test","password":"1234567890"}`

				req := httptest.NewRequest(http.MethodPost, "/v1/auth/login", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				var empty *entities.Token
				mockService.On("Login", mock.Anything, mock.Anything).
					Times(1).
					Return(empty, errors.NewError(errors.ErrInvalidCredentials))

				hdl := handlers.New(mockService, log)

				err := hdl.HandleLogin(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusConflict))
			})
		})
	})

	Context("consume the handleRefreshToken endpoint", func() {
		When("return success response", func() {
			It("can return a new session", func() {
				body := `{"token":"my-refresh-token"}`

				req := httptest.NewRequest(http.MethodPost, "/v1/auth/refresh-token", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				mockService.On("RenewSession", mock.Anything, mock.Anything).
					Times(1).
					Return(&entities.Token{
						AccessToken:  "my-access-token",
						RefreshToken: "my-refresh-token",
						Type:         "Bearer",
						ExpiresIn:    3600,
					}, nil)

				hdl := handlers.New(mockService, log)

				err := hdl.HandleRefreshToken(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusOK))

				var result *entities.Token
				err = json.Unmarshal(rec.Body.Bytes(), &result)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(HaveExistingField("AccessToken"))
				Expect(result).To(HaveExistingField("RefreshToken"))
				Expect(result.AccessToken).To(Equal("my-access-token"))
				Expect(result.RefreshToken).To(Equal("my-refresh-token"))
			})
		})

		When("sent invalid request", func() {
			It("produces an http bad request caused by bind params", func() {
				body := `{"token":"my-refresh-token"`
				req := httptest.NewRequest(http.MethodPost, "/v1/auth/refresh-token", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				hdl := handlers.New(mockService, log)
				err := hdl.HandleRefreshToken(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
			})

			It("produces an http bad request caused by validations", func() {
				body := `{"refresh_token":"my-refresh-token"}`
				req := httptest.NewRequest(http.MethodPost, "/v1/auth/refresh-token", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				hdl := handlers.New(mockService, log)
				err := hdl.HandleRefreshToken(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusBadRequest))
			})
		})

		When("service layer produces an error", func() {
			It("return an invalid http code", func() {
				body := `{"token":"my-refresh-token"}`

				req := httptest.NewRequest(http.MethodPost, "/v1/auth/refresh-token", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				var empty *entities.Token
				mockService.On("RenewSession", mock.Anything, mock.Anything).
					Times(1).
					Return(empty, errors.NewError(errors.ErrRefreshToken))

				hdl := handlers.New(mockService, log)

				err := hdl.HandleRefreshToken(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(http.StatusConflict))
			})
		})
	})
})
