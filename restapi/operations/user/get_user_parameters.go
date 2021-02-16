// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetUserParams creates a new GetUserParams object
//
// There are no default values defined in the spec.
func NewGetUserParams() GetUserParams {

	return GetUserParams{}
}

// GetUserParams contains all the bound params for the get user operation
// typically these are obtained from a http.Request
//
// swagger:parameters getUser
type GetUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*authorization header contains firebase ID token
	  Required: true
	  In: header
	*/
	Authorization string
	/*header that contains userId of user you want to get
	  Required: true
	  In: header
	*/
	UserID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUserParams() beforehand.
func (o *GetUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindUserID(r.Header[http.CanonicalHeaderKey("userId")], true, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorization binds and validates parameter Authorization from header.
func (o *GetUserParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindUserID binds and validates parameter UserID from header.
func (o *GetUserParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("userId", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("userId", "header", raw); err != nil {
		return err
	}
	o.UserID = raw

	return nil
}
