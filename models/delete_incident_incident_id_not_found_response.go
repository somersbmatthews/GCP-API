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

// DeleteIncidentIncidentIDNotFoundResponse delete incident incident Id not found response
//
// swagger:model DeleteIncidentIncidentIdNotFoundResponse
type DeleteIncidentIncidentIDNotFoundResponse struct {

	// deleted
	Deleted *bool `json:"deleted,omitempty"`

	// incident Id
	// Required: true
	IncidentID *string `json:"incidentId"`
}

// Validate validates this delete incident incident Id not found response
func (m *DeleteIncidentIncidentIDNotFoundResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncidentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteIncidentIncidentIDNotFoundResponse) validateIncidentID(formats strfmt.Registry) error {

	if err := validate.Required("incidentId", "body", m.IncidentID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete incident incident Id not found response based on context it is used
func (m *DeleteIncidentIncidentIDNotFoundResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteIncidentIncidentIDNotFoundResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteIncidentIncidentIDNotFoundResponse) UnmarshalBinary(b []byte) error {
	var res DeleteIncidentIncidentIDNotFoundResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
