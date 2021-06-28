// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteIncidentGoodResponse delete incident good response
//
// swagger:model DeleteIncidentGoodResponse
type DeleteIncidentGoodResponse struct {

	// deleted
	Deleted *bool `json:"Deleted,omitempty"`

	// ID
	// Required: true
	ID *string `json:"ID"`
}

// Validate validates this delete incident good response
func (m *DeleteIncidentGoodResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteIncidentGoodResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete incident good response based on context it is used
func (m *DeleteIncidentGoodResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteIncidentGoodResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteIncidentGoodResponse) UnmarshalBinary(b []byte) error {
	var res DeleteIncidentGoodResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
