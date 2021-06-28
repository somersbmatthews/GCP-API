// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gircapp/api/swagger/restapi/operations"
	"github.com/gircapp/api/swagger/restapi/operations/coroner_incident"
	"github.com/gircapp/api/swagger/restapi/operations/derm_incident"
	"github.com/gircapp/api/swagger/restapi/operations/e_m_t_incident"
	"github.com/gircapp/api/swagger/restapi/operations/e_n_t_incident"
	"github.com/gircapp/api/swagger/restapi/operations/e_r_p_incident"
	"github.com/gircapp/api/swagger/restapi/operations/fire_incident"
	"github.com/gircapp/api/swagger/restapi/operations/incident"
	"github.com/gircapp/api/swagger/restapi/operations/medical_expert"
	"github.com/gircapp/api/swagger/restapi/operations/surg_incident"
	"github.com/gircapp/api/swagger/restapi/operations/user"
	"github.com/gircapp/api/swagger/restapi/operations/verify"
	"github.com/gircapp/api/swagger/restapi/operations/vet_incident"
)

//go:generate swagger generate server --target ../../swagger --name Girc --spec ../swagger18.yaml --principal interface{} --exclude-main

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

	if api.CoronerIncidentCreateCorIncidentHandler == nil {
		api.CoronerIncidentCreateCorIncidentHandler = coroner_incident.CreateCorIncidentHandlerFunc(func(params coroner_incident.CreateCorIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation coroner_incident.CreateCorIncident has not yet been implemented")
		})
	}
	if api.DermIncidentCreateDermIncidentHandler == nil {
		api.DermIncidentCreateDermIncidentHandler = derm_incident.CreateDermIncidentHandlerFunc(func(params derm_incident.CreateDermIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation derm_incident.CreateDermIncident has not yet been implemented")
		})
	}
	if api.EmtIncidentCreateEMTIncidentHandler == nil {
		api.EmtIncidentCreateEMTIncidentHandler = e_m_t_incident.CreateEMTIncidentHandlerFunc(func(params e_m_t_incident.CreateEMTIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation e_m_t_incident.CreateEMTIncident has not yet been implemented")
		})
	}
	if api.EntIncidentCreateENTIncidentHandler == nil {
		api.EntIncidentCreateENTIncidentHandler = e_n_t_incident.CreateENTIncidentHandlerFunc(func(params e_n_t_incident.CreateENTIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation e_n_t_incident.CreateENTIncident has not yet been implemented")
		})
	}
	if api.ErpIncidentCreateERPIncidentHandler == nil {
		api.ErpIncidentCreateERPIncidentHandler = e_r_p_incident.CreateERPIncidentHandlerFunc(func(params e_r_p_incident.CreateERPIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation e_r_p_incident.CreateERPIncident has not yet been implemented")
		})
	}
	if api.FireIncidentCreateFireIncidentHandler == nil {
		api.FireIncidentCreateFireIncidentHandler = fire_incident.CreateFireIncidentHandlerFunc(func(params fire_incident.CreateFireIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation fire_incident.CreateFireIncident has not yet been implemented")
		})
	}
	if api.IncidentCreateIncidentHandler == nil {
		api.IncidentCreateIncidentHandler = incident.CreateIncidentHandlerFunc(func(params incident.CreateIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.CreateIncident has not yet been implemented")
		})
	}
	if api.SurgIncidentCreateSurgIncidentHandler == nil {
		api.SurgIncidentCreateSurgIncidentHandler = surg_incident.CreateSurgIncidentHandlerFunc(func(params surg_incident.CreateSurgIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation surg_incident.CreateSurgIncident has not yet been implemented")
		})
	}
	if api.UserCreateUserHandler == nil {
		api.UserCreateUserHandler = user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		})
	}
	if api.VetIncidentCreateVetIncidentHandler == nil {
		api.VetIncidentCreateVetIncidentHandler = vet_incident.CreateVetIncidentHandlerFunc(func(params vet_incident.CreateVetIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation vet_incident.CreateVetIncident has not yet been implemented")
		})
	}
	if api.MedicalExpertDeleteExpertHandler == nil {
		api.MedicalExpertDeleteExpertHandler = medical_expert.DeleteExpertHandlerFunc(func(params medical_expert.DeleteExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.DeleteExpert has not yet been implemented")
		})
	}
	if api.IncidentDeleteIncidentsHandler == nil {
		api.IncidentDeleteIncidentsHandler = incident.DeleteIncidentsHandlerFunc(func(params incident.DeleteIncidentsParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.DeleteIncidents has not yet been implemented")
		})
	}
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.MedicalExpertGetExpertHandler == nil {
		api.MedicalExpertGetExpertHandler = medical_expert.GetExpertHandlerFunc(func(params medical_expert.GetExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.GetExpert has not yet been implemented")
		})
	}
	if api.IncidentGetIncidentsHandler == nil {
		api.IncidentGetIncidentsHandler = incident.GetIncidentsHandlerFunc(func(params incident.GetIncidentsParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.GetIncidents has not yet been implemented")
		})
	}
	if api.UserGetUserHandler == nil {
		api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		})
	}
	if api.MedicalExpertLoginExpertHandler == nil {
		api.MedicalExpertLoginExpertHandler = medical_expert.LoginExpertHandlerFunc(func(params medical_expert.LoginExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.LoginExpert has not yet been implemented")
		})
	}
	if api.MedicalExpertLogoutExpertHandler == nil {
		api.MedicalExpertLogoutExpertHandler = medical_expert.LogoutExpertHandlerFunc(func(params medical_expert.LogoutExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.LogoutExpert has not yet been implemented")
		})
	}
	if api.MedicalExpertRegisterExpertHandler == nil {
		api.MedicalExpertRegisterExpertHandler = medical_expert.RegisterExpertHandlerFunc(func(params medical_expert.RegisterExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.RegisterExpert has not yet been implemented")
		})
	}
	if api.MedicalExpertUpdateExpertHandler == nil {
		api.MedicalExpertUpdateExpertHandler = medical_expert.UpdateExpertHandlerFunc(func(params medical_expert.UpdateExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.UpdateExpert has not yet been implemented")
		})
	}
	if api.IncidentUpdateIncidentsHandler == nil {
		api.IncidentUpdateIncidentsHandler = incident.UpdateIncidentsHandlerFunc(func(params incident.UpdateIncidentsParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.UpdateIncidents has not yet been implemented")
		})
	}
	if api.UserUpdateUserHandler == nil {
		api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		})
	}
	if api.VerifyVerifyHandler == nil {
		api.VerifyVerifyHandler = verify.VerifyHandlerFunc(func(params verify.VerifyParams) middleware.Responder {
			return middleware.NotImplemented("operation verify.Verify has not yet been implemented")
		})
	}
	if api.VerifyVerifyExpertHandler == nil {
		api.VerifyVerifyExpertHandler = verify.VerifyExpertHandlerFunc(func(params verify.VerifyExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation verify.VerifyExpert has not yet been implemented")
		})
	}

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
