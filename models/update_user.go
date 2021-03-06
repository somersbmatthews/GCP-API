// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateUser update user
//
// swagger:model UpdateUser
type UpdateUser struct {

	// degree
	Degree string `json:"degree,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// specialty
	Specialty string `json:"specialty,omitempty"`
}

// Validate validates this update user
func (m *UpdateUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update user based on context it is used
func (m *UpdateUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUser) UnmarshalBinary(b []byte) error {
	var res UpdateUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
