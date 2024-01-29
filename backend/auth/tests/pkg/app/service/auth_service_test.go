package service_test

import (
	"context"
	"errors"
	"github.com/auth0/go-auth0/authentication"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/app/service"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/constants"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/entities"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"net/http"
	"os"
)

const (
	domain       = "my.domain.us.auth0.com"
	clientId     = "dummy"
	clientSecret = "dummy"
	scope        = "test"
	realm        = "DBTest"
	algorithm    = "HS256"
)

var sessResp = `{"access_token":"my-access-token","refresh_token":"my-refresh-token","expires_in":3600,"token_type":"Bearer","scope":"test"}`

var _ = Describe("Service", func() {
	var httpClient *http.Client
	var auth0Auth *authentication.Authentication
	var log logger.Logger

	BeforeEach(func() {
		httpClient = http.DefaultClient
		httpmock.Reset()
		log = logger.New(logger.Opts{
			LogLevel:  "debug",
			AppName:   "auth-service-test",
			BlackList: []string{"access_token"},
			Tags: map[string]string{
				"env":   "test",
				"owner": "unit_tests",
			},
		})

		_ = os.Setenv(constants.Auth0Scope, scope)
		_ = os.Setenv(constants.Auth0Audience, "https://"+domain+"/api/v2/")
		_ = os.Setenv(constants.Auth0Realm, realm)

		// configure validator
		jwksRes := `{"keys":[{"kty":"RSA","use":"sig","e":"AQAB","n":"","kid":"id","x5t":"","x5c":[""],"alg":"RS256"}]}`
		resp := httpmock.NewStringResponder(http.StatusOK, jwksRes)
		httpmock.RegisterResponder(http.MethodGet, "https://"+domain+"/.well-known/jwks.json", resp)

		auth, err := authentication.New(
			context.Background(),
			domain,
			authentication.WithClientID(clientId),
			authentication.WithClientSecret(clientSecret),
			authentication.WithClient(httpClient),
			authentication.WithNoRetries(),
		)
		Expect(err).NotTo(HaveOccurred())

		auth0Auth = auth
	})

	AfterEach(func() {
		os.Clearenv()
	})

	Context("generate sessions", func() {
		When("http client is configured", func() {
			It("creates a valid session", func() {
				ctx := context.TODO()

				// configure http mock listener
				resp := httpmock.NewStringResponder(http.StatusOK, sessResp).Times(1)
				httpmock.RegisterResponder(http.MethodPost, "https://"+domain+"/oauth/token", resp)

				// set request
				loginReq := &entities.LoginReq{
					Username: "test",
					Password: "my-p4$w0rd",
				}

				srv := service.New(auth0Auth, log)
				session, err := srv.Login(ctx, loginReq)
				Expect(err).NotTo(HaveOccurred())
				Expect(session).NotTo(BeNil())
				Expect(session).To(HaveExistingField("AccessToken"))
				Expect(session.AccessToken).To(Equal("my-access-token"))
			})

			It("auth0 returns an error", func() {
				ctx := context.TODO()

				// configure http mock listener
				resp := httpmock.NewStringResponder(
					http.StatusBadRequest,
					`{"statusCode":400,"error":"invalid_scope","error_description":"Scope must be an array or a string"}`).
					Times(1)
				httpmock.RegisterResponder(http.MethodPost, "https://"+domain+"/oauth/token", resp)

				// set request
				loginReq := &entities.LoginReq{
					Username: "another-test",
					Password: "simple-password",
				}

				srv := service.New(auth0Auth, log)
				session, err := srv.Login(ctx, loginReq)
				Expect(session).To(BeNil())
				Expect(err).NotTo(BeNil())
			})

			It("auth0 returns an unexpected error", func() {
				ctx := context.TODO()

				// configure http mock listener
				errResp := httpmock.NewErrorResponder(errors.New("generic error"))
				httpmock.RegisterResponder(http.MethodPost, "https://"+domain+"/oauth/token", errResp)

				// set request
				loginReq := &entities.LoginReq{
					Username: "another-test",
					Password: "simple-password",
				}

				srv := service.New(auth0Auth, log)
				session, err := srv.Login(ctx, loginReq)
				Expect(session).To(BeNil())
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Context("refresh token", func() {
		When("http client is configured", func() {
			It("renew session", func() {
				ctx := context.TODO()

				// configure http mock listener
				resp := httpmock.NewStringResponder(http.StatusOK, sessResp).Times(1)
				httpmock.RegisterResponder(http.MethodPost, "https://"+domain+"/oauth/token", resp)

				// set request
				refreshTokenReq := &entities.RefreshTokenReq{
					Token: "my-refresh-token",
				}

				srv := service.New(auth0Auth, log)
				session, err := srv.RenewSession(ctx, refreshTokenReq)
				Expect(err).NotTo(HaveOccurred())
				Expect(session).NotTo(BeNil())
				Expect(session).To(HaveExistingField("AccessToken"))
				Expect(session.AccessToken).To(Equal("my-access-token"))
			})

			It("auth0 returns an error", func() {
				ctx := context.TODO()

				// configure http mock listener
				resp := httpmock.NewStringResponder(
					http.StatusBadRequest,
					`{"statusCode":400,"error":"invalid_scope","error_description":"Scope must be an array or a string"}`).
					Times(1)
				httpmock.RegisterResponder(http.MethodPost, "https://"+domain+"/oauth/token", resp)

				// set request
				refreshTokenReq := &entities.RefreshTokenReq{
					Token: "my-refresh-token",
				}

				srv := service.New(auth0Auth, log)
				session, err := srv.RenewSession(ctx, refreshTokenReq)
				Expect(session).To(BeNil())
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
