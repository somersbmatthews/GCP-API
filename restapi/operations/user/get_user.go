// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserHandlerFunc turns a function with the right signature into a get user handler
type GetUserHandlerFunc func(GetUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserHandlerFunc) Handle(params GetUserParams) middleware.Responder {
	return fn(params)
}

// GetUserHandler interface for that can handle valid get user params
type GetUserHandler interface {
	Handle(GetUserParams) middleware.Responder
}

// NewGetUser creates a new http.Handler for the get user operation
func NewGetUser(ctx *middleware.Context, handler GetUserHandler) *GetUser {
	return &GetUser{Context: ctx, Handler: handler}
}

/* GetUser swagger:route GET /v2/user user getUser

Get user's information.

Get a user's information.

*/
type GetUser struct {
	Context *middleware.Context
	Handler GetUserHandler
}

func (o *GetUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
