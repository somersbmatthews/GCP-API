// Code generated by go-swagger; DO NOT EDIT.

package medical_expert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RegisterExpertHandlerFunc turns a function with the right signature into a register expert handler
type RegisterExpertHandlerFunc func(RegisterExpertParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterExpertHandlerFunc) Handle(params RegisterExpertParams) middleware.Responder {
	return fn(params)
}

// RegisterExpertHandler interface for that can handle valid register expert params
type RegisterExpertHandler interface {
	Handle(RegisterExpertParams) middleware.Responder
}

// NewRegisterExpert creates a new http.Handler for the register expert operation
func NewRegisterExpert(ctx *middleware.Context, handler RegisterExpertHandler) *RegisterExpert {
	return &RegisterExpert{Context: ctx, Handler: handler}
}

/* RegisterExpert swagger:route POST /v3/expert Medical Expert registerExpert

register expert

Use this to register a expert

*/
type RegisterExpert struct {
	Context *middleware.Context
	Handler RegisterExpertHandler
}

func (o *RegisterExpert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRegisterExpertParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
