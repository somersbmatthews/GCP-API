// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/gircapp/api/models"
)

// NewSendNotificationParams creates a new SendNotificationParams object
//
// There are no default values defined in the spec.
func NewSendNotificationParams() SendNotificationParams {

	return SendNotificationParams{}
}

// SendNotificationParams contains all the bound params for the send notification operation
// typically these are obtained from a http.Request
//
// swagger:parameters sendNotification
type SendNotificationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*authorization header Firebase id token
	  Required: true
	  In: header
	*/
	Authorization string
	/*
	  Required: true
	  In: body
	*/
	Incident *models.Notification
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSendNotificationParams() beforehand.
func (o *SendNotificationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Notification
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("incident", "body", ""))
			} else {
				res = append(res, errors.NewParseError("incident", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Incident = &body
			}
		}
	} else {
		res = append(res, errors.Required("incident", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorization binds and validates parameter Authorization from header.
func (o *SendNotificationParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Authorization", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Authorization", "header", raw); err != nil {
		return err
	}
	o.Authorization = raw

	return nil
}
