// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/somersbmatthews/gircapp2/fba"
	"github.com/somersbmatthews/gircapp2/models"
	"github.com/somersbmatthews/gircapp2/pg"
	"github.com/somersbmatthews/gircapp2/restapi/operations"
	"github.com/somersbmatthews/gircapp2/restapi/operations/user"
	"github.com/somersbmatthews/gircapp2/restapi/operations/verify"
)

//go:generate swagger generate server --target ../../gircapp2 --name Girc --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.GircAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GircAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// if api.IncidentCreateIncidentHandler == nil {
	// 	api.IncidentCreateIncidentHandler = incident.CreateIncidentHandlerFunc(func(params incident.CreateIncidentParams) middleware.Responder {
	// 		return middleware.NotImplemented("operation incident.CreateIncident has not yet been implemented")
	// 	})
	// }

	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
		ctx := context.Background()

		tokenStr := params.Authorization

		ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return middleware.Error(400, createUserBadResponse(params))
		}
		newUser := pg.User{
			UserId:     *params.User.UserID,
			Email:      *params.User.Email,
			Speciality: *params.User.Speciality,
			Degree:     *params.User.Degree,
			Name:       *params.User.Name,
			Verified:   false,
		}
		payload := pg.CreateUser(ctx, newUser)
		response := user.NewCreateUserOK()
		response.WithPayload(payload)
		return response

	})

	api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
		ctx := context.Background()

		tokenStr := params.Authorization

		ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return middleware.Error(400, nil)
		}

		payload, ok := pg.GetUser(ctx, *params.User.UserID)
		if !ok {
			return middleware.Error(404, nil)
		}
		response := user.NewGetUserOK()
		response.WithPayload(payload)
		return response
	})

	api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
		ctx := context.Background()

		tokenStr := params.Authorization

		ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return middleware.Error(404, updateUserInvalidResponse(params))
		}

		updatedUser := pg.User{
			UserId:     *params.User.UserID,
			Email:      params.User.Email,
			Speciality: params.User.Speciality,
			Degree:     params.User.Degree,
			Name:       params.User.Name,
		}

		payload, ok := pg.UpdateUser(ctx, updatedUser)
		if !ok {
			return middleware.Error(404, nil)
		}
		response := user.NewUpdateUserOK()
		response.WithPayload(payload)
		return response

		//	return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
	})

	api.VerifyVerifyHandler = verify.VerifyHandlerFunc(func(params verify.VerifyParams) middleware.Responder {
		ctx := context.Background()

		// tokenStr := params.Authorization

		// ok := fba.VerifyToken(ctx, tokenStr)
		// if !ok {
		// 	return middleware.Error(404, updateUserInvalidResponse(params))
		// }

		verifiedUser := models.Verify{
			UserID:   params.Verified.UserID,
			Verified: params.Verified.Verified,
		}

		payload, ok := pg.VerifyUser(ctx, verifiedUser)
		if !ok {
			return middleware.Error(404, nil)
		}
		response := verify.NewVerifyOK()
		response.WithPayload(payload)
		return response
	})

	// api.IncidentDeleteIncidentsHandler = incident.DeleteIncidentsHandlerFunc(func(params incident.DeleteIncidentsParams) middleware.Responder {

	// 	return middleware.NotImplemented("operation incident.DeleteIncidents has not yet been implemented")
	// })

	// if api.IncidentGetIncidentsHandler == nil {
	// 	api.IncidentGetIncidentsHandler = incident.GetIncidentsHandlerFunc(func(params incident.GetIncidentsParams) middleware.Responder {
	// 		return middleware.NotImplemented("operation incident.GetIncidents has not yet been implemented")
	// 	})
	// }

	// if api.IncidentUpdateIncidentsHandler == nil {
	// 	api.IncidentUpdateIncidentsHandler = incident.UpdateIncidentsHandlerFunc(func(params incident.UpdateIncidentsParams) middleware.Responder {
	// 		return middleware.NotImplemented("operation incident.UpdateIncidents has not yet been implemented")
	// 	})
	// }

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

func createUserBadResponse(params user.CreateUserParams) models.CreateUserBadResponse {
	booleanFalse := false
	createUserBadResponse := models.CreateUserBadResponse{
		Created:    &booleanFalse,
		Email:      params.User.Email,
		Degree:     params.User.Degree,
		Speciality: params.User.Speciality,
		Name:       params.User.Name,
	}
	return createUserBadResponse
}

func updateUserInvalidResponse(params user.UpdateUserParams) models.UpdateUserInvalidResponse {
	booleanFalse := false
	updateUserInvalidResponse := models.UpdateUserInvalidResponse{
		Degree:   &params.User.Degree,
		Email:    &params.User.Email,
		Name:     &params.User.Name,
		Verified: &booleanFalse,
		Updated:  &booleanFalse,
	}
	return updateUserInvalidResponse

}