// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GoodResponse good response
//
// swagger:model GoodResponse
type GoodResponse struct {

	// message
	Message string `json:"Message,omitempty"`
}

// Validate validates this good response
func (m *GoodResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this good response based on context it is used
func (m *GoodResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GoodResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoodResponse) UnmarshalBinary(b []byte) error {
	var res GoodResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
