package main

import (
	"context"
	"fmt"
	"github.com/auth0/go-auth0/authentication"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/internal/infrastructure/http"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/app/handlers"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/app/service"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/constants"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/utils"
	"io"
)

//	@title			Authentication API
//	@version		1.0
//	@description	This is the authentication api.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Ricardo Romero
//	@contact.email	ricardo.jonathan.romero@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1/auth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						X-Api-Key

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	log := logger.New(logger.Opts{
		LogLevel: utils.GetEnv(constants.LogLevel, constants.DefaultLogLevel),
		AppName:  constants.AppName,
	})

	log.Debug("configuring auth client")
	auth, err := authentication.New(
		context.TODO(),
		utils.GetEnv(constants.Auth0Domain, constants.Empty),
		authentication.WithClientID(utils.GetEnv(constants.Auth0ClientId, constants.Empty)),
		authentication.WithClientSecret(utils.GetEnv(constants.Auth0ClientSecret, constants.Empty)),
	)

	if err != nil {
		log.Fatalf("error configuring auth0 client: %v", err)
	}

	// initialize business core
	srv := service.New(auth, log)
	handles := handlers.New(srv, log)

	// initialize server
	server, jaegerTracer := http.NewServer(handles)
	if jaegerTracer != nil {
		defer func(jaegerTracer io.Closer) {
			err := jaegerTracer.Close()
			if err != nil {
				log.Errorf("error closing jaeger io closer: %v", err)
			}
		}(jaegerTracer)
	}

	port := utils.GetEnv(constants.ServerPort, constants.DefaultServerPort)
	if err = server.Start(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("error starting server on port %s: %v", port, err)
	}
}
