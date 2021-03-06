// Code generated by go-swagger; DO NOT EDIT.

package medical_expert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateExpertHandlerFunc turns a function with the right signature into a update expert handler
type UpdateExpertHandlerFunc func(UpdateExpertParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateExpertHandlerFunc) Handle(params UpdateExpertParams) middleware.Responder {
	return fn(params)
}

// UpdateExpertHandler interface for that can handle valid update expert params
type UpdateExpertHandler interface {
	Handle(UpdateExpertParams) middleware.Responder
}

// NewUpdateExpert creates a new http.Handler for the update expert operation
func NewUpdateExpert(ctx *middleware.Context, handler UpdateExpertHandler) *UpdateExpert {
	return &UpdateExpert{Context: ctx, Handler: handler}
}

/* UpdateExpert swagger:route PATCH /v3/expert Medical Expert updateExpert

register expert

Use this to register a expert

*/
type UpdateExpert struct {
	Context *middleware.Context
	Handler UpdateExpertHandler
}

func (o *UpdateExpert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateExpertParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
