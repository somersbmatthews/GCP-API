// Code generated by go-swagger; DO NOT EDIT.

package verify

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VerifyExpertHandlerFunc turns a function with the right signature into a verify expert handler
type VerifyExpertHandlerFunc func(VerifyExpertParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyExpertHandlerFunc) Handle(params VerifyExpertParams) middleware.Responder {
	return fn(params)
}

// VerifyExpertHandler interface for that can handle valid verify expert params
type VerifyExpertHandler interface {
	Handle(VerifyExpertParams) middleware.Responder
}

// NewVerifyExpert creates a new http.Handler for the verify expert operation
func NewVerifyExpert(ctx *middleware.Context, handler VerifyExpertHandler) *VerifyExpert {
	return &VerifyExpert{Context: ctx, Handler: handler}
}

/* VerifyExpert swagger:route PATCH /v3/verify verify verifyExpert

use this to verify or unverify a Medcal Expert, for testing only

verified field is true to verify, false to unverify

*/
type VerifyExpert struct {
	Context *middleware.Context
	Handler VerifyExpertHandler
}

func (o *VerifyExpert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVerifyExpertParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}