// Code generated by go-swagger; DO NOT EDIT.

package incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetIncidentsHandlerFunc turns a function with the right signature into a get incidents handler
type GetIncidentsHandlerFunc func(GetIncidentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIncidentsHandlerFunc) Handle(params GetIncidentsParams) middleware.Responder {
	return fn(params)
}

// GetIncidentsHandler interface for that can handle valid get incidents params
type GetIncidentsHandler interface {
	Handle(GetIncidentsParams) middleware.Responder
}

// NewGetIncidents creates a new http.Handler for the get incidents operation
func NewGetIncidents(ctx *middleware.Context, handler GetIncidentsHandler) *GetIncidents {
	return &GetIncidents{Context: ctx, Handler: handler}
}

/* GetIncidents swagger:route GET /incident incident getIncidents

Get incidents

Use this to get all incidents created by a user

*/
type GetIncidents struct {
	Context *middleware.Context
	Handler GetIncidentsHandler
}

func (o *GetIncidents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetIncidentsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
