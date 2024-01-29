package main

import (
	"fmt"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/internal/infrastructure/http"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/app/handlers"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/app/services"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/constants"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/products/pkg/domain/repository"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/db/mongodb"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/utils"
	"io"
	"strings"
)

//	@title			Products API
//	@version		1.0
//	@description	This is the catalog of products API
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Ricardo Romero
//	@contact.email	ricardo.jonathan.romero@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/v1/products

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						X-Api-Key

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	log := logger.New(logger.Opts{
		LogLevel:  utils.GetEnv(constants.LogLevel, constants.DefaultLogLevel),
		BlackList: strings.Split(utils.GetEnv(constants.CustomBlackList, constants.Empty), constants.Comma),
		AppName:   constants.AppName,
	})

	db, err := mongodb.New(mongodb.MongoOpts{
		Uri:        utils.GetEnv(constants.MongoDBUri, constants.Empty),
		DB:         utils.GetEnv(constants.MongoDBName, constants.Empty),
		Collection: utils.GetEnv(constants.MongoDBCollection, constants.Empty),
	})
	if err != nil {
		log.Fatalf("error configuring mongodb client: %v", err)
	}

	// initialize business core
	repo := repository.New(db, log)
	srv := services.New(repo, log)
	handles := handlers.New(srv, log)

	// initialize server
	server, jaegerTracer, err := http.NewServer(handles)
	if err != nil {
		log.Fatalf("error generating server: %v", err)
	}

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
