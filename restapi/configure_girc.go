// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gircapp/api/auth"
	"github.com/gircapp/api/fba"
	"github.com/gircapp/api/models"
	"github.com/gircapp/api/pg"
	"github.com/gircapp/api/restapi/operations"
	"github.com/gircapp/api/restapi/operations/e_n_t_incident"
	"github.com/gircapp/api/restapi/operations/incident"
	"github.com/gircapp/api/restapi/operations/medical_expert"
	"github.com/gircapp/api/restapi/operations/user"

	//	"github.com/gircapp/api/restapi/operations/verify"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// ./cloud_sql_proxy -instances=gircapp:us-central1:gircapppostgres=tcp:5432
//go:generate swagger generate server --target ../../gircapp2 --name Girc --spec ../swagger.yaml --principal interface{}

// USE THIS ONE, THIS IS THE UPDATED COMMAND:
// swagger generate server -f ./swagger/swagger5.yaml --exclude-main -A girc
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

		ok := auth.VerifyOldDataToken(tokenStr)
		if !ok {
			response := user.NewCreateUserUnauthorized()

			response.WithPayload(&models.BadResponse{})
			return response
		}
		newUser := pg.User{
			UserID:    *params.User.ID,
			Email:     *params.User.Email,
			Specialty: *params.User.Specialty,
			Degree:    *params.User.Degree,
			Name:      *params.User.Name,
			Verified:  false,
		}
		_, ok = pg.CreateUser(ctx, newUser)
		if !ok {
			response := user.NewCreateUserBadRequest()

			response.WithPayload(&models.BadResponse{
				Message: fmt.Sprint("couldn't create user with user id: %s", params.User.ID),
			})
			return response
		}
		response := user.NewCreateUserOK()
		response.WithPayload(&models.GoodResponse{
			Message: fmt.Sprintf("successfully create user with user id: %s", params.User.ID),
		})
		return response
	})

	// api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
	// 	ctx := context.Background()
	// 	tokenStr := params.Authorization
	// 	ok := auth.VerifyOldDataToken(tokenStr)
	// 	if !ok {
	// 		return middleware.Error(401, models.BadResponse{
	// 			Message: "Validation of firebase idToken failed.",
	// 		})
	// 	}
	// 	payload, ok := pg.GetUser(ctx, userID)
	// 	if !ok {
	// 		response := user.NewGetUserNotFound()

	// 		response.WithPayload(&models.BadResponse{
	// 			Message: fmt.Sprintf("user could not be created with user id: %s", userID),
	// 		})
	// 		return response
	// 	}
	// 	response := user.NewGetUserOK()
	// 	response.WithPayload(payload)
	// 	return response
	// })

	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID := params.User
		ok := auth.VerifyOldDataToken(tokenStr)
		booleanFalse := false
		if !ok {
			return middleware.Error(401, models.BadResponse{
				Message: "Validation of firebase idToken failed.",
			})
		}

		_, ok = pg.DeleteUser(ctx, userID)
		if !ok {
			return middleware.Error(404, models.DeleteUserBadResponse{
				UserID:  &userID,
				Deleted: &booleanFalse,
			})
		}
		response := user.NewDeleteUserOK()
		response.WithPayload(&models.GoodResponse{
			Message: fmt.Sprintf("user successfully deleted with user id: %s", userID),
		})
		return response
	})

	// api.VerifyVerifyHandler = verify.VerifyHandlerFunc(func(params verify.VerifyParams) middleware.Responder {
	// 	ctx := context.Background()
	// 	tokenStr := params.Authorization
	// 	userID, ok := fba.VerifyToken(ctx, tokenStr)
	// 	if !ok {
	// 		return middleware.Error(401, models.BadResponse{
	// 			Message: "Validation of firebase idToken failed.",
	// 		})
	// 	}
	// 	verifiedUser := models.Verify{
	// 		Verified: params.Verified.Verified,
	// 	}
	// 	payload, ok := pg.VerifyUser(ctx, verifiedUser, userID)
	// 	if !ok {
	// 		return middleware.Error(404, verifyUserNotFoundResponse(params, userID))
	// 	}
	// 	response := verify.NewVerifyOK()
	// 	response.WithPayload(payload)
	// 	return response
	// })

	api.IncidentCreateIncidentHandler = incident.CreateIncidentHandlerFunc(func(params incident.CreateIncidentParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		ok := auth.VerifyOldDataToken(tokenStr)
		// booleanFalse := false
		if !ok {
			return middleware.Error(401, models.BadResponse{
				Message: "Validation of firebase idToken failed.",
			})
		}
		payload, ok := pg.CreateIncidents(ctx, params.Incident.Incidents, params.Incident.UserID)
		if !ok {
			return middleware.Error(409, models.BadResponse{
				Message: fmt.Sprintf("could not create incidents for userId: %s", params.Incident.UserID),
			})
		}
		response := incident.NewCreateIncidentOK()
		response.WithPayload(payload)
		return response
	})

	api.IncidentGetIncidentsHandler = incident.GetIncidentsHandlerFunc(func(params incident.GetIncidentsParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID := params.User
		ok := auth.VerifyOldDataToken(tokenStr)
		if !ok {
			return middleware.Error(401, models.BadResponse{
				Message: "Validation of firebase idToken failed.",
			})
		}
		payload, ok := pg.GetIncidents(ctx, userID)
		if !ok {
			//	response := incident.NewGetIncidentsNotFound()
			return middleware.Error(404, &models.BadResponse{
				Message: fmt.Sprintf("Incidents could not be found for user id : %s", params.User),
			})
		}
		response := incident.NewGetIncidentsOK()
		response.WithPayload(payload)
		return response
	})

	// api.IncidentUpdateIncidentsHandler = incident.UpdateIncidentsHandlerFunc(func(params incident.UpdateIncidentsParams) middleware.Responder {
	// 	ctx := context.Background()
	// 	tokenStr := params.Authorization
	// 	_, ok := fba.VerifyToken(ctx, tokenStr)
	// 	booleanFalse := false
	// 	if !ok {
	// 		return middleware.Error(401, models.BadResponse{
	// 			Message: "Validation of firebase idToken failed.",
	// 		})
	// 	}
	// 	payload, ok := pg.UpdateIncident(ctx, *params.Incident)
	// 	if !ok {
	// 		return middleware.Error(404, models.UpdateIncidentIncidentIDNotFoundResponse{
	// 			ID:                            params.Incident.ID,
	// 			DateOfIncident:                params.Incident.DateOfIncident,
	// 			ApproximatePatientAge:         params.Incident.ApproximatePatientAge,
	// 			Gender:                        params.Incident.Gender,
	// 			LongTermPrognosis:             params.Incident.LongTermPrognosis,
	// 			IncidentDescription:           params.Incident.IncidentDescription,
	// 			Anterior:                      params.Incident.Anterior,
	// 			ObjectConsistency:             params.Incident.ObjectConsistency,
	// 			ObjectBasicShape:              params.Incident.ObjectBasicShape,
	// 			WhatMaterialIsTheObjectMadeOf: params.Incident.WhatMaterialIsTheObjectMadeOf,
	// 			TheObjectIs:                   params.Incident.TheObjectIs,
	// 			LargestLength:                 params.Incident.LargestLength,
	// 			Updated:                       &booleanFalse,
	// 		})
	// 	}
	// 	response := incident.NewUpdateIncidentsOK()
	// 	response.WithPayload(payload)
	// 	return response
	// })

	api.IncidentDeleteIncidentsHandler = incident.DeleteIncidentsHandlerFunc(func(params incident.DeleteIncidentsParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID := params.User
		ok := auth.VerifyOldDataToken(tokenStr)

		if !ok {
			return middleware.Error(401, models.BadResponse{
				Message: "could not authenticate request",
			})
		}
		_, ok = pg.DeleteIncidents(ctx, userID)
		if !ok {
			return middleware.Error(404, models.BadResponse{
				Message: "could not delete incidents with that user id",
			})
		}
		payload := models.GoodResponse{
			Message: fmt.Sprintf("incidents for user id: %s", userID),
		}
		response := incident.NewDeleteIncidentsOK()
		response.WithPayload(&payload)
		return response
	})

	api.MedicalExpertRegisterExpertHandler = medical_expert.RegisterExpertHandlerFunc(func(params medical_expert.RegisterExpertParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return middleware.Error(401, models.BadResponse{
				Message: "Validation of firebase idToken failed.",
			})
		}
		payload, ok := pg.CreateExpertWithAutoDirectorAndEmailVerification(ctx, params.Expert, userID)
		if !ok {
			return middleware.Error(404, models.BadResponse{
				Message: "could not create medical Expert",
			})
		}
		response := medical_expert.NewCreateExpertOK()
		response.WithPayload(payload)
		return response
	})

	api.MedicalExpertGetExpertHandler = medical_expert.GetExpertHandlerFunc(func(params medical_expert.GetExpertParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return medical_expert.NewGetExpertUnauthorized()
		}
		payload, ok := pg.GetMedicalExpert(ctx, userID)
		if !ok {
			return medical_expert.NewGetExpertNotFound()
		}
		response := medical_expert.NewGetExpertOK()
		response.WithPayload(payload)
		return response
	})

	api.MedicalExpertUpdateExpertHandler = medical_expert.UpdateExpertHandlerFunc(func(params medical_expert.UpdateExpertParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return medical_expert.NewUpdateExpertUnauthorized()
		}
		payload, ok := pg.UpdateMedicalExpert(ctx, *params.Expert, userID)
		if !ok {
			return medical_expert.NewUpdateExpertNotFound()
		}
		response := medical_expert.NewUpdateExpertOK()
		response.WithPayload(payload)
		return response
	})

	api.MedicalExpertDeleteExpertHandler = medical_expert.DeleteExpertHandlerFunc(func(params medical_expert.DeleteExpertParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return medical_expert.NewDeleteExpertUnauthorized()
		}
		payload, ok := pg.DeleteMedicalExpert(ctx, userID)
		if !ok {
			return medical_expert.NewDeleteExpertNotFound()
		}
		response := medical_expert.NewDeleteExpertOK()
		response.WithPayload(payload)
		return response
	})

	api.EntIncidentCreateENTIncidentHandler = e_n_t_incident.CreateENTIncidentHandlerFunc(func(params e_n_t_incident.CreateENTIncidentParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return e_n_t_incident.NewCreateENTIncidentUnauthorized()
		}
		payload, ok := pg.CreateENTIncident(ctx, params.Incident, userID)
		if !ok {
			return e_n_t_incident.NewCreateENTIncidentNotFound()
		}
		response := e_n_t_incident.NewCreateENTIncidentOK()
		response.WithPayload(payload)
		return response

	})

	api.EntIncidentGetENTIncidentsHandler = e_n_t_incident.GetENTIncidentsHandlerFunc(func(params e_n_t_incident.GetENTIncidentsParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return e_n_t_incident.NewGetENTIncidentsUnauthorized()
		}
		payload, ok := pg.GetENTIncidents(ctx, userID)
		if !ok {
			return e_n_t_incident.NewGetENTIncidentsNotFound()
		}
		response := e_n_t_incident.NewGetENTIncidentsOK()
		response.WithPayload(payload)
		return response
	})

	api.EntIncidentUpdateENTIncidentHandler = e_n_t_incident.UpdateENTIncidentHandlerFunc(func(params e_n_t_incident.UpdateENTIncidentParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return e_n_t_incident.NewUpdateENTIncidentUnauthorized()
		}
		payload, ok := pg.UpdateENTIncident(ctx, *params.Incident, userID)
		if !ok {
			return e_n_t_incident.NewGetENTIncidentsNotFound()
		}
		response := e_n_t_incident.NewUpdateENTIncidentOK()
		response.WithPayload(payload)
		return response
	})

	api.EntIncidentDeleteENTIncidentHandler = e_n_t_incident.DeleteENTIncidentHandlerFunc(func(params e_n_t_incident.DeleteENTIncidentParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		_, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return e_n_t_incident.NewDeleteENTIncidentUnauthorized()
		}
		payload, ok := pg.DeleteENTIncident(ctx, *params.Incident.ENTIncidentID)
		if !ok {
			return e_n_t_incident.NewDeleteENTIncidentNotFound()
		}
		response := e_n_t_incident.NewDeleteENTIncidentOK()
		response.WithPayload(payload)
		return response
	})

	api.MedicalExpertUpdateFCMtokenHandler = medical_expert.UpdateFCMtokenHandlerFunc(func(params medical_expert.UpdateFCMtokenParams) middleware.Responder {
		ctx := context.Background()
		tokenStr := params.Authorization
		userID, ok := fba.VerifyToken(ctx, tokenStr)
		if !ok {
			return medical_expert.NewUpdateFCMtokenUnauthorized()
		}
		payload, ok := pg.UpdateFCMToken(ctx, *params.FCMToken, userID)
		if !ok {
			return medical_expert.NewUpdateFCMtokenNotFound()
		}
		response := medical_expert.NewUpdateFCMtokenOK()
		response.WithPayload(payload)
		return response
	})

	// api.AdminVerifyExpertHandler = admin.VerifyExpertHandlerFunc(func(params admin.VerifyExpertParams) middleware.Responder {
	// 	ctx := context.Background()
	// 	tokenStr := params.Authorization
	//_, ok := fba.VerifyAdminToken(ctx, tokenStr)
	// })

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {
		pg.Conn.Close()
	}
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

// func createUserBadResponse(params user.CreateUserParams) models.CreateUserBadResponse {
// 	booleanFalse := false
// 	createUserBadResponse := models.CreateUserBadResponse{
// 		Created:   &booleanFalse,
// 		Email:     params.User.Email,
// 		Degree:    params.User.Degree,
// 		Specialty: params.User.Specialty,
// 		Name:      params.User.Name,
// 	}
// 	return createUserBadResponse

// }

// func updateUserInvalidResponse(params user.UpdateUserParams) models.UpdateUserInvalidResponse {
// 	booleanFalse := false
// 	updateUserInvalidResponse := models.UpdateUserInvalidResponse{
// 		Degree:   &params.User.Degree,
// 		Email:    &params.User.Email,
// 		Name:     &params.User.Name,
// 		Verified: &booleanFalse,
// 		Updated:  &booleanFalse,
// 	}
// 	return updateUserInvalidResponse
// }

// func updateUserNotFoundResponse(params user.UpdateUserParams, userID string) models.UpdateUserNotFoundResponse {
// 	booleanFalse := false
// 	updateUserNotFoundResponse := models.UpdateUserNotFoundResponse{
// 		UserID:  &userID,
// 		Updated: &booleanFalse,
// 	}
// 	return updateUserNotFoundResponse
// }

// func verifyUserNotFoundResponse(params verify.VerifyParams, userID string) models.UpdateUserNotFoundResponse {
// 	booleanFalse := false
// 	updateUserNotFoundResponse := models.UpdateUserNotFoundResponse{
// 		UserID:  &userID,
// 		Updated: &booleanFalse,
// 	}
// 	return updateUserNotFoundResponse
// }
