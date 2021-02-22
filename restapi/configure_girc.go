// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/gircapp/api/fba"
	"github.com/gircapp/api/models"
	"github.com/gircapp/api/pg"
	"github.com/gircapp/api/restapi/operations"
	"github.com/gircapp/api/restapi/operations/incident"
	"github.com/gircapp/api/restapi/operations/user"
	"github.com/gircapp/api/restapi/operations/verify"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

//go:generate swagger generate server --target ../../gircapp2 --name Girc --spec ../swagger.yaml --principal interface{}
// swagger generate server -f swagger5.yaml --exclude-main -A girc
func configureFlags(api *operations.GircAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GircAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf

	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
		ctx := context.Background()

		tokenStr := params.Authorization

		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			response := user.NewCreateUserBadRequest()
			booleanFalse := false
			createUserBadResponse := models.CreateUserBadResponse{
				Created:    &booleanFalse,
				Email:      params.User.Email,
				Degree:     params.User.Degree,
				Speciality: params.User.Speciality,
				Name:       params.User.Name,
			}
			response.WithPayload(&createUserBadResponse)
			return response
		}
		newUser := pg.User{
			UserID:     userID,
			Email:      *params.User.Email,
			Speciality: *params.User.Speciality,
			Degree:     *params.User.Degree,
			Name:       *params.User.Name,
			Verified:   false,
		}
		payload, ok := pg.CreateUser(ctx, newUser)
		if !ok {
			response := user.NewCreateUserBadRequest()
			booleanFalse := false
			createUserBadResponse := models.CreateUserBadResponse{
				Created:    &booleanFalse,
				Email:      params.User.Email,
				Degree:     params.User.Degree,
				Speciality: params.User.Speciality,
				Name:       params.User.Name,
			}
			response.WithPayload(&createUserBadResponse)
			return response
		}
		response := user.NewCreateUserOK()
		response.WithPayload(payload)
		return response
	})

	api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			response := user.NewGetUserBadRequest()
			getUserBadResponse := models.GetUserBadResponse{
				UserID: &userID,
			}
			response.WithPayload(&getUserBadResponse)
			return response
		}
		payload, ok := pg.GetUser(ctx, userID)
		if !ok {
			response := user.NewGetUserNotFound()
			getUserBadResponse := models.GetUserBadResponse{
				UserID: &userID,
			}
			response.WithPayload(&getUserBadResponse)
			return response
		}
		response := user.NewGetUserOK()
		response.WithPayload(payload)
		return response
	})

	api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return middleware.Error(400, updateUserInvalidResponse(params))
		}
		updatedUser := pg.User{
			UserID:     userID,
			Email:      params.User.Email,
			Speciality: params.User.Speciality,
			Degree:     params.User.Degree,
			Name:       params.User.Name,
		}
		payload, ok := pg.UpdateUser(ctx, updatedUser)
		if !ok {
			return middleware.Error(404, updateUserNotFoundResponse(params, userID))
		}
		response := user.NewUpdateUserOK()
		response.WithPayload(payload)
		return response
	})

	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		booleanFalse := false
		if !ok {
			return middleware.Error(400, models.DeleteUserBadResponse{
				Deleted: &booleanFalse,
				UserID:  &userID,
			})
		}

		payload, ok := pg.DeleteUser(ctx, userID)
		if !ok {
			return middleware.Error(404, models.DeleteUserBadResponse{
				UserID:  &userID,
				Deleted: &booleanFalse,
			})
		}
		response := user.NewDeleteUserOK()
		response.WithPayload(payload)
		return response
	})

	api.VerifyVerifyHandler = verify.VerifyHandlerFunc(func(params verify.VerifyParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		verifiedUser := models.Verify{
			Verified: params.Verified.Verified,
		}
		payload, ok := pg.VerifyUser(ctx, verifiedUser, userID)
		if !ok {
			return middleware.Error(404, verifyUserNotFoundResponse(params, userID))
		}
		response := verify.NewVerifyOK()
		response.WithPayload(payload)
		return response
	})

	api.IncidentCreateIncidentHandler = incident.CreateIncidentHandlerFunc(func(params incident.CreateIncidentParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		booleanFalse := false
		if !ok {
			return middleware.Error(400, models.CreateIncidentInvalidIncidentResponse{
				ID:                            params.Incident.ID,
				DateOfIncident:                params.Incident.DateOfIncident,
				ApproximatePatientAge:         params.Incident.ApproximatePatientAge,
				Gender:                        params.Incident.Gender,
				LongTermPrognosis:             params.Incident.LongTermPrognosis,
				IncidentDescription:           params.Incident.IncidentDescription,
				Anterior:                      params.Incident.Anterior,
				ObjectConsistency:             params.Incident.ObjectConsistency,
				ObjectBasicShape:              params.Incident.ObjectBasicShape,
				WhatMaterialIsTheObjectMadeOf: params.Incident.WhatMaterialIsTheObjectMadeOf,
				TheObjectIs:                   params.Incident.TheObjectIs,
				LargestLength:                 params.Incident.LargestLength,
				Created:                       &booleanFalse,
			})
		}
		payload := pg.CreateIncident(ctx, *params.Incident, userID)
		response := incident.NewCreateIncidentOK()
		response.WithPayload(payload)
		return response
	})

	// api.IncidentGetIncidentsHandler = incident.GetIncidentsHandlerFunc(func(params incident.GetIncidentsParams) middleware.Responder {
	// 	return middleware.NotImplemented("operation incident.GetIncidents has not yet been implemented")
	// })

	api.IncidentUpdateIncidentsHandler = incident.UpdateIncidentsHandlerFunc(func(params incident.UpdateIncidentsParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		_, ok := fba.VerifyToken(ctx, tokenStr)
		booleanFalse := false
		if !ok {
			return middleware.Error(400, models.UpdateIncidentIncidentIDNotFoundResponse{
				ID:                            params.Incident.ID,
				DateOfIncident:                params.Incident.DateOfIncident,
				ApproximatePatientAge:         params.Incident.ApproximatePatientAge,
				Gender:                        params.Incident.Gender,
				LongTermPrognosis:             params.Incident.LongTermPrognosis,
				IncidentDescription:           params.Incident.IncidentDescription,
				Anterior:                      params.Incident.Anterior,
				ObjectConsistency:             params.Incident.ObjectConsistency,
				ObjectBasicShape:              params.Incident.ObjectBasicShape,
				WhatMaterialIsTheObjectMadeOf: params.Incident.WhatMaterialIsTheObjectMadeOf,
				TheObjectIs:                   params.Incident.TheObjectIs,
				LargestLength:                 params.Incident.LargestLength,
				Updated:                       &booleanFalse,
			})
		}
		payload, ok := pg.UpdateIncident(ctx, *params.Incident)
		if !ok {
			return middleware.Error(404, models.UpdateIncidentIncidentIDNotFoundResponse{
				ID:                            params.Incident.ID,
				DateOfIncident:                params.Incident.DateOfIncident,
				ApproximatePatientAge:         params.Incident.ApproximatePatientAge,
				Gender:                        params.Incident.Gender,
				LongTermPrognosis:             params.Incident.LongTermPrognosis,
				IncidentDescription:           params.Incident.IncidentDescription,
				Anterior:                      params.Incident.Anterior,
				ObjectConsistency:             params.Incident.ObjectConsistency,
				ObjectBasicShape:              params.Incident.ObjectBasicShape,
				WhatMaterialIsTheObjectMadeOf: params.Incident.WhatMaterialIsTheObjectMadeOf,
				TheObjectIs:                   params.Incident.TheObjectIs,
				LargestLength:                 params.Incident.LargestLength,
				Updated:                       &booleanFalse,
			})
		}
		response := incident.NewUpdateIncidentsOK()
		response.WithPayload(payload)
		return response
	})

	api.IncidentDeleteIncidentsHandler = incident.DeleteIncidentsHandlerFunc(func(params incident.DeleteIncidentsParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		_, ok := fba.VerifyToken(ctx, tokenStr)
		booleanFalse := false
		if !ok {
			return middleware.Error(404, models.DeleteIncidentIncidentIDNotFoundResponse{
				Deleted: &booleanFalse,
				ID:      params.Incident.ID,
			})
		}
		payload, ok := pg.DeleteIncident(ctx, *params.Incident.ID)
		if !ok {
			return middleware.Error(404, models.DeleteIncidentIncidentIDNotFoundResponse{
				ID:      params.Incident.ID,
				Deleted: &booleanFalse,
			})
		}
		response := incident.NewDeleteIncidentsOK()
		response.WithPayload(payload)
		return response
	})
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

func updateUserNotFoundResponse(params user.UpdateUserParams, userID string) models.UpdateUserNotFoundResponse {
	booleanFalse := false
	updateUserNotFoundResponse := models.UpdateUserNotFoundResponse{
		UserID:  &userID,
		Updated: &booleanFalse,
	}
	return updateUserNotFoundResponse
}

func verifyUserNotFoundResponse(params verify.VerifyParams, userID string) models.UpdateUserNotFoundResponse {
	booleanFalse := false
	updateUserNotFoundResponse := models.UpdateUserNotFoundResponse{
		UserID:  &userID,
		Updated: &booleanFalse,
	}
	return updateUserNotFoundResponse
}
