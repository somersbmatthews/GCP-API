// Code generated by go-swagger; DO NOT EDIT.

package medical_expert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetExpertParams creates a new GetExpertParams object
//
// There are no default values defined in the spec.
func NewGetExpertParams() GetExpertParams {

	return GetExpertParams{}
}

// GetExpertParams contains all the bound params for the get expert operation
// typically these are obtained from a http.Request
//
// swagger:parameters getExpert
type GetExpertParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*App version from ios App
	  Required: true
	  In: header
	*/
	AppVersion string
	/*if using google oauth, set that token here
	  Required: true
	  In: header
	*/
	Authorization string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetExpertParams() beforehand.
func (o *GetExpertParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAppVersion(r.Header[http.CanonicalHeaderKey("App-Version")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAppVersion binds and validates parameter AppVersion from header.
func (o *GetExpertParams) bindAppVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("App-Version", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("App-Version", "header", raw); err != nil {
		return err
	}
	o.AppVersion = raw

	return nil
}

// bindAuthorization binds and validates parameter Authorization from header.
func (o *GetExpertParams) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
