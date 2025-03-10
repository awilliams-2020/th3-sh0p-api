// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/stripe/stripe-go/v81"

	"th3-sh0p-api/handlers"
	"th3-sh0p-api/internal"
	"th3-sh0p-api/restapi/operations"
)

//go:generate swagger generate server --target ../../th3-sh0p-api --name Th3Sh0pAPI --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.Th3Sh0pAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Th3Sh0pAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.BearerAuth = internal.ValidateToken

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	stripe.Key = os.Getenv("STRIPE_KEY")

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetPubKeyHandler = operations.GetPubKeyHandlerFunc(handlers.GetPubKey)

	api.GetUserCreditHandler = operations.GetUserCreditHandlerFunc(handlers.GetUserCredit)

	api.GetImagesHandler = operations.GetImagesHandlerFunc(handlers.GetImages)

	api.GetImagesPagesHandler = operations.GetImagesPagesHandlerFunc(handlers.GetImagesPages)

	api.GetGoogleProfileHandler = operations.GetGoogleProfileHandlerFunc(handlers.GetGoogleProfile)

	api.PostImageHandler = operations.PostImageHandlerFunc(handlers.CreateImage)

	api.PostImagePackHandler = operations.PostImagePackHandlerFunc(handlers.PostImagePack)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
