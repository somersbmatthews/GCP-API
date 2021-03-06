// Code generated by go-swagger; DO NOT EDIT.

package e_n_t_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteENTIncidentHandlerFunc turns a function with the right signature into a delete e n t incident handler
type DeleteENTIncidentHandlerFunc func(DeleteENTIncidentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteENTIncidentHandlerFunc) Handle(params DeleteENTIncidentParams) middleware.Responder {
	return fn(params)
}

// DeleteENTIncidentHandler interface for that can handle valid delete e n t incident params
type DeleteENTIncidentHandler interface {
	Handle(DeleteENTIncidentParams) middleware.Responder
}

// NewDeleteENTIncident creates a new http.Handler for the delete e n t incident operation
func NewDeleteENTIncident(ctx *middleware.Context, handler DeleteENTIncidentHandler) *DeleteENTIncident {
	return &DeleteENTIncident{Context: ctx, Handler: handler}
}

/* DeleteENTIncident swagger:route DELETE /v3/entincident ENT Incident deleteENTIncident

Delete ENT Incident

Use this to delete an ENT incident

*/
type DeleteENTIncident struct {
	Context *middleware.Context
	Handler DeleteENTIncidentHandler
}

func (o *DeleteENTIncident) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteENTIncidentParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
