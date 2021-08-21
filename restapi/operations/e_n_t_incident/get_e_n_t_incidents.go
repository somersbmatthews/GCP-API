// Code generated by go-swagger; DO NOT EDIT.

package e_n_t_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetENTIncidentsHandlerFunc turns a function with the right signature into a get e n t incidents handler
type GetENTIncidentsHandlerFunc func(GetENTIncidentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetENTIncidentsHandlerFunc) Handle(params GetENTIncidentsParams) middleware.Responder {
	return fn(params)
}

// GetENTIncidentsHandler interface for that can handle valid get e n t incidents params
type GetENTIncidentsHandler interface {
	Handle(GetENTIncidentsParams) middleware.Responder
}

// NewGetENTIncidents creates a new http.Handler for the get e n t incidents operation
func NewGetENTIncidents(ctx *middleware.Context, handler GetENTIncidentsHandler) *GetENTIncidents {
	return &GetENTIncidents{Context: ctx, Handler: handler}
}

/* GetENTIncidents swagger:route GET /v3/entincident ENT Incident getENTIncidents

get all medical experts entincidents

use this to get all of a users entIncidents

*/
type GetENTIncidents struct {
	Context *middleware.Context
	Handler GetENTIncidentsHandler
}

func (o *GetENTIncidents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetENTIncidentsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
