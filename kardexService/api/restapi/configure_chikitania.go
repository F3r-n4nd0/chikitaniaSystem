// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	modelsApi "kardexService/api/models"
	"kardexService/eventstore"
	"kardexService/kardex/usecase"
	"kardexService/models"
	"kardexService/repository"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"kardexService/api/restapi/operations"
	"kardexService/api/restapi/operations/kardex"
	"kardexService/api/restapi/operations/purchase"
	"kardexService/api/restapi/operations/sale"
)

//go:generate swagger generate kardexService --target ../../api --name Chikitania --spec ../swagger/swagger.json

func configureFlags(api *operations.ChikitaniaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ChikitaniaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.KardexRegisterKardexHandler = kardex.RegisterKardexHandlerFunc(func(params kardex.RegisterKardexParams) middleware.Responder {
		gRpcClient :=  eventstore.NewGRpcClient("localhost:50051")
		userCaseKardex :=  usecase.NewKardexUseCase(2, gRpcClient)
		newKardex := models.NewKardex{
			ProductId: params.Body.ProductID,
		}
		if err := userCaseKardex.Register(context.Background(), newKardex); err != nil {
			log.Printf("error while registering kardex: %v", err)
			return kardex.NewRegisterKardexMethodNotAllowed()
		}
		return kardex.NewRegisterKardexCreated()
	})

	api.KardexGetKardexByIDHandler = kardex.GetKardexByIDHandlerFunc(func(params kardex.GetKardexByIDParams) middleware.Responder {
		repository := repository.NewMongoDbClient("mongodb://localhost:27017")
		userCaseQueryKardex :=  usecase.NewQueryKardexUseCase(2, repository )
		result, err := userCaseQueryKardex.GetById(context.Background(), params.KardexID)
		if  err != nil {
			return kardex.NewGetKardexByIDBadRequest()
		}
		if result == nil {
			return  kardex.NewGetKardexByIDNotFound()
		}
		dataKardex :=  modelsApi.Kardex{
			ID:        result.KardexId,
			ProductID: result.ProductId,
		}
		return kardex.NewGetKardexByIDOK().WithPayload(&dataKardex)
	})



	if api.PurchasePlacePurchaseHandler == nil {
		api.PurchasePlacePurchaseHandler = purchase.PlacePurchaseHandlerFunc(func(params purchase.PlacePurchaseParams) middleware.Responder {
			return middleware.NotImplemented("operation purchase.PlacePurchase has not yet been implemented")
		})
	}
	if api.SalePlaceSaleHandler == nil {
		api.SalePlaceSaleHandler = sale.PlaceSaleHandlerFunc(func(params sale.PlaceSaleParams) middleware.Responder {
			return middleware.NotImplemented("operation sale.PlaceSale has not yet been implemented")
		})
	}
	if api.PurchaseReturnPurchaseHandler == nil {
		api.PurchaseReturnPurchaseHandler = purchase.ReturnPurchaseHandlerFunc(func(params purchase.ReturnPurchaseParams) middleware.Responder {
			return middleware.NotImplemented("operation purchase.ReturnPurchase has not yet been implemented")
		})
	}
	if api.SaleReturnSaleHandler == nil {
		api.SaleReturnSaleHandler = sale.ReturnSaleHandlerFunc(func(params sale.ReturnSaleParams) middleware.Responder {
			return middleware.NotImplemented("operation sale.ReturnSale has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}
	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS kardexService starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as kardexService is initialized but not run yet, this function will be called.
// If you need to modify a config, store kardexService instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return middleware.Redoc(middleware.RedocOpts{},handler)
}
