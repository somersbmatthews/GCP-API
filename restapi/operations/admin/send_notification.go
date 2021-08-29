// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SendNotificationHandlerFunc turns a function with the right signature into a send notification handler
type SendNotificationHandlerFunc func(SendNotificationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SendNotificationHandlerFunc) Handle(params SendNotificationParams) middleware.Responder {
	return fn(params)
}

// SendNotificationHandler interface for that can handle valid send notification params
type SendNotificationHandler interface {
	Handle(SendNotificationParams) middleware.Responder
}

// NewSendNotification creates a new http.Handler for the send notification operation
func NewSendNotification(ctx *middleware.Context, handler SendNotificationHandler) *SendNotification {
	return &SendNotification{Context: ctx, Handler: handler}
}

/* SendNotification swagger:route POST /v3/admin/sendnotification Admin sendNotification

send notification to medical expert

use this to send notification to a medical expert with id

*/
type SendNotification struct {
	Context *middleware.Context
	Handler SendNotificationHandler
}

func (o *SendNotification) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSendNotificationParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}