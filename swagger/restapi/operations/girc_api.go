// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

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

// NewGircAPI creates a new Girc instance
func NewGircAPI(spec *loads.Document) *GircAPI {
	return &GircAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CoronerIncidentCreateCorIncidentHandler: coroner_incident.CreateCorIncidentHandlerFunc(func(params coroner_incident.CreateCorIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation coroner_incident.CreateCorIncident has not yet been implemented")
		}),
		DermIncidentCreateDermIncidentHandler: derm_incident.CreateDermIncidentHandlerFunc(func(params derm_incident.CreateDermIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation derm_incident.CreateDermIncident has not yet been implemented")
		}),
		EmtIncidentCreateEMTIncidentHandler: e_m_t_incident.CreateEMTIncidentHandlerFunc(func(params e_m_t_incident.CreateEMTIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation e_m_t_incident.CreateEMTIncident has not yet been implemented")
		}),
		EntIncidentCreateENTIncidentHandler: e_n_t_incident.CreateENTIncidentHandlerFunc(func(params e_n_t_incident.CreateENTIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation e_n_t_incident.CreateENTIncident has not yet been implemented")
		}),
		ErpIncidentCreateERPIncidentHandler: e_r_p_incident.CreateERPIncidentHandlerFunc(func(params e_r_p_incident.CreateERPIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation e_r_p_incident.CreateERPIncident has not yet been implemented")
		}),
		FireIncidentCreateFireIncidentHandler: fire_incident.CreateFireIncidentHandlerFunc(func(params fire_incident.CreateFireIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation fire_incident.CreateFireIncident has not yet been implemented")
		}),
		IncidentCreateIncidentHandler: incident.CreateIncidentHandlerFunc(func(params incident.CreateIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.CreateIncident has not yet been implemented")
		}),
		SurgIncidentCreateSurgIncidentHandler: surg_incident.CreateSurgIncidentHandlerFunc(func(params surg_incident.CreateSurgIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation surg_incident.CreateSurgIncident has not yet been implemented")
		}),
		UserCreateUserHandler: user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		}),
		VetIncidentCreateVetIncidentHandler: vet_incident.CreateVetIncidentHandlerFunc(func(params vet_incident.CreateVetIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation vet_incident.CreateVetIncident has not yet been implemented")
		}),
		MedicalExpertDeleteExpertHandler: medical_expert.DeleteExpertHandlerFunc(func(params medical_expert.DeleteExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.DeleteExpert has not yet been implemented")
		}),
		IncidentDeleteIncidentsHandler: incident.DeleteIncidentsHandlerFunc(func(params incident.DeleteIncidentsParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.DeleteIncidents has not yet been implemented")
		}),
		UserDeleteUserHandler: user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		}),
		MedicalExpertGetExpertHandler: medical_expert.GetExpertHandlerFunc(func(params medical_expert.GetExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.GetExpert has not yet been implemented")
		}),
		IncidentGetIncidentsHandler: incident.GetIncidentsHandlerFunc(func(params incident.GetIncidentsParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.GetIncidents has not yet been implemented")
		}),
		UserGetUserHandler: user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		}),
		MedicalExpertLoginExpertHandler: medical_expert.LoginExpertHandlerFunc(func(params medical_expert.LoginExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.LoginExpert has not yet been implemented")
		}),
		MedicalExpertLogoutExpertHandler: medical_expert.LogoutExpertHandlerFunc(func(params medical_expert.LogoutExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.LogoutExpert has not yet been implemented")
		}),
		MedicalExpertRegisterExpertHandler: medical_expert.RegisterExpertHandlerFunc(func(params medical_expert.RegisterExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.RegisterExpert has not yet been implemented")
		}),
		MedicalExpertUpdateExpertHandler: medical_expert.UpdateExpertHandlerFunc(func(params medical_expert.UpdateExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation medical_expert.UpdateExpert has not yet been implemented")
		}),
		IncidentUpdateIncidentsHandler: incident.UpdateIncidentsHandlerFunc(func(params incident.UpdateIncidentsParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.UpdateIncidents has not yet been implemented")
		}),
		UserUpdateUserHandler: user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		}),
		VerifyVerifyHandler: verify.VerifyHandlerFunc(func(params verify.VerifyParams) middleware.Responder {
			return middleware.NotImplemented("operation verify.Verify has not yet been implemented")
		}),
		VerifyVerifyExpertHandler: verify.VerifyExpertHandlerFunc(func(params verify.VerifyExpertParams) middleware.Responder {
			return middleware.NotImplemented("operation verify.VerifyExpert has not yet been implemented")
		}),
	}
}

/*GircAPI This is test server for GIRC app */
type GircAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// CoronerIncidentCreateCorIncidentHandler sets the operation handler for the create cor incident operation
	CoronerIncidentCreateCorIncidentHandler coroner_incident.CreateCorIncidentHandler
	// DermIncidentCreateDermIncidentHandler sets the operation handler for the create derm incident operation
	DermIncidentCreateDermIncidentHandler derm_incident.CreateDermIncidentHandler
	// EmtIncidentCreateEMTIncidentHandler sets the operation handler for the create e m t incident operation
	EmtIncidentCreateEMTIncidentHandler e_m_t_incident.CreateEMTIncidentHandler
	// EntIncidentCreateENTIncidentHandler sets the operation handler for the create e n t incident operation
	EntIncidentCreateENTIncidentHandler e_n_t_incident.CreateENTIncidentHandler
	// ErpIncidentCreateERPIncidentHandler sets the operation handler for the create e r p incident operation
	ErpIncidentCreateERPIncidentHandler e_r_p_incident.CreateERPIncidentHandler
	// FireIncidentCreateFireIncidentHandler sets the operation handler for the create fire incident operation
	FireIncidentCreateFireIncidentHandler fire_incident.CreateFireIncidentHandler
	// IncidentCreateIncidentHandler sets the operation handler for the create incident operation
	IncidentCreateIncidentHandler incident.CreateIncidentHandler
	// SurgIncidentCreateSurgIncidentHandler sets the operation handler for the create surg incident operation
	SurgIncidentCreateSurgIncidentHandler surg_incident.CreateSurgIncidentHandler
	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// VetIncidentCreateVetIncidentHandler sets the operation handler for the create vet incident operation
	VetIncidentCreateVetIncidentHandler vet_incident.CreateVetIncidentHandler
	// MedicalExpertDeleteExpertHandler sets the operation handler for the delete expert operation
	MedicalExpertDeleteExpertHandler medical_expert.DeleteExpertHandler
	// IncidentDeleteIncidentsHandler sets the operation handler for the delete incidents operation
	IncidentDeleteIncidentsHandler incident.DeleteIncidentsHandler
	// UserDeleteUserHandler sets the operation handler for the delete user operation
	UserDeleteUserHandler user.DeleteUserHandler
	// MedicalExpertGetExpertHandler sets the operation handler for the get expert operation
	MedicalExpertGetExpertHandler medical_expert.GetExpertHandler
	// IncidentGetIncidentsHandler sets the operation handler for the get incidents operation
	IncidentGetIncidentsHandler incident.GetIncidentsHandler
	// UserGetUserHandler sets the operation handler for the get user operation
	UserGetUserHandler user.GetUserHandler
	// MedicalExpertLoginExpertHandler sets the operation handler for the login expert operation
	MedicalExpertLoginExpertHandler medical_expert.LoginExpertHandler
	// MedicalExpertLogoutExpertHandler sets the operation handler for the logout expert operation
	MedicalExpertLogoutExpertHandler medical_expert.LogoutExpertHandler
	// MedicalExpertRegisterExpertHandler sets the operation handler for the register expert operation
	MedicalExpertRegisterExpertHandler medical_expert.RegisterExpertHandler
	// MedicalExpertUpdateExpertHandler sets the operation handler for the update expert operation
	MedicalExpertUpdateExpertHandler medical_expert.UpdateExpertHandler
	// IncidentUpdateIncidentsHandler sets the operation handler for the update incidents operation
	IncidentUpdateIncidentsHandler incident.UpdateIncidentsHandler
	// UserUpdateUserHandler sets the operation handler for the update user operation
	UserUpdateUserHandler user.UpdateUserHandler
	// VerifyVerifyHandler sets the operation handler for the verify operation
	VerifyVerifyHandler verify.VerifyHandler
	// VerifyVerifyExpertHandler sets the operation handler for the verify expert operation
	VerifyVerifyExpertHandler verify.VerifyExpertHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *GircAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *GircAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *GircAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *GircAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *GircAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *GircAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *GircAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *GircAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *GircAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the GircAPI
func (o *GircAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CoronerIncidentCreateCorIncidentHandler == nil {
		unregistered = append(unregistered, "coroner_incident.CreateCorIncidentHandler")
	}
	if o.DermIncidentCreateDermIncidentHandler == nil {
		unregistered = append(unregistered, "derm_incident.CreateDermIncidentHandler")
	}
	if o.EmtIncidentCreateEMTIncidentHandler == nil {
		unregistered = append(unregistered, "e_m_t_incident.CreateEMTIncidentHandler")
	}
	if o.EntIncidentCreateENTIncidentHandler == nil {
		unregistered = append(unregistered, "e_n_t_incident.CreateENTIncidentHandler")
	}
	if o.ErpIncidentCreateERPIncidentHandler == nil {
		unregistered = append(unregistered, "e_r_p_incident.CreateERPIncidentHandler")
	}
	if o.FireIncidentCreateFireIncidentHandler == nil {
		unregistered = append(unregistered, "fire_incident.CreateFireIncidentHandler")
	}
	if o.IncidentCreateIncidentHandler == nil {
		unregistered = append(unregistered, "incident.CreateIncidentHandler")
	}
	if o.SurgIncidentCreateSurgIncidentHandler == nil {
		unregistered = append(unregistered, "surg_incident.CreateSurgIncidentHandler")
	}
	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}
	if o.VetIncidentCreateVetIncidentHandler == nil {
		unregistered = append(unregistered, "vet_incident.CreateVetIncidentHandler")
	}
	if o.MedicalExpertDeleteExpertHandler == nil {
		unregistered = append(unregistered, "medical_expert.DeleteExpertHandler")
	}
	if o.IncidentDeleteIncidentsHandler == nil {
		unregistered = append(unregistered, "incident.DeleteIncidentsHandler")
	}
	if o.UserDeleteUserHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserHandler")
	}
	if o.MedicalExpertGetExpertHandler == nil {
		unregistered = append(unregistered, "medical_expert.GetExpertHandler")
	}
	if o.IncidentGetIncidentsHandler == nil {
		unregistered = append(unregistered, "incident.GetIncidentsHandler")
	}
	if o.UserGetUserHandler == nil {
		unregistered = append(unregistered, "user.GetUserHandler")
	}
	if o.MedicalExpertLoginExpertHandler == nil {
		unregistered = append(unregistered, "medical_expert.LoginExpertHandler")
	}
	if o.MedicalExpertLogoutExpertHandler == nil {
		unregistered = append(unregistered, "medical_expert.LogoutExpertHandler")
	}
	if o.MedicalExpertRegisterExpertHandler == nil {
		unregistered = append(unregistered, "medical_expert.RegisterExpertHandler")
	}
	if o.MedicalExpertUpdateExpertHandler == nil {
		unregistered = append(unregistered, "medical_expert.UpdateExpertHandler")
	}
	if o.IncidentUpdateIncidentsHandler == nil {
		unregistered = append(unregistered, "incident.UpdateIncidentsHandler")
	}
	if o.UserUpdateUserHandler == nil {
		unregistered = append(unregistered, "user.UpdateUserHandler")
	}
	if o.VerifyVerifyHandler == nil {
		unregistered = append(unregistered, "verify.VerifyHandler")
	}
	if o.VerifyVerifyExpertHandler == nil {
		unregistered = append(unregistered, "verify.VerifyExpertHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *GircAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *GircAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *GircAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *GircAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *GircAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *GircAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the girc API
func (o *GircAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *GircAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/corincident"] = coroner_incident.NewCreateCorIncident(o.context, o.CoronerIncidentCreateCorIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/dermincident"] = derm_incident.NewCreateDermIncident(o.context, o.DermIncidentCreateDermIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/emtincident"] = e_m_t_incident.NewCreateEMTIncident(o.context, o.EmtIncidentCreateEMTIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/entincident"] = e_n_t_incident.NewCreateENTIncident(o.context, o.EntIncidentCreateENTIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/erpincident"] = e_r_p_incident.NewCreateERPIncident(o.context, o.ErpIncidentCreateERPIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/fireincident"] = fire_incident.NewCreateFireIncident(o.context, o.FireIncidentCreateFireIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v2/incident"] = incident.NewCreateIncident(o.context, o.IncidentCreateIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/surgincident"] = surg_incident.NewCreateSurgIncident(o.context, o.SurgIncidentCreateSurgIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v2/user"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/vetincident"] = vet_incident.NewCreateVetIncident(o.context, o.VetIncidentCreateVetIncidentHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v3/expert"] = medical_expert.NewDeleteExpert(o.context, o.MedicalExpertDeleteExpertHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v2/incident"] = incident.NewDeleteIncidents(o.context, o.IncidentDeleteIncidentsHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v2/user"] = user.NewDeleteUser(o.context, o.UserDeleteUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v3/expert"] = medical_expert.NewGetExpert(o.context, o.MedicalExpertGetExpertHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/incident"] = incident.NewGetIncidents(o.context, o.IncidentGetIncidentsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v2/user"] = user.NewGetUser(o.context, o.UserGetUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v3/expert/login"] = medical_expert.NewLoginExpert(o.context, o.MedicalExpertLoginExpertHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v3/expert/logout"] = medical_expert.NewLogoutExpert(o.context, o.MedicalExpertLogoutExpertHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v3/expert"] = medical_expert.NewRegisterExpert(o.context, o.MedicalExpertRegisterExpertHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/v3/expert"] = medical_expert.NewUpdateExpert(o.context, o.MedicalExpertUpdateExpertHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/v2/incident"] = incident.NewUpdateIncidents(o.context, o.IncidentUpdateIncidentsHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/v2/user"] = user.NewUpdateUser(o.context, o.UserUpdateUserHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/v2/verify"] = verify.NewVerify(o.context, o.VerifyVerifyHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/v3/verify"] = verify.NewVerifyExpert(o.context, o.VerifyVerifyExpertHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *GircAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *GircAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *GircAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *GircAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *GircAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
