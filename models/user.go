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

// User user
//
// swagger:model User
type User struct {

	// ID
	// Required: true
	ID *string `json:"ID"`

	// degree
	// Required: true
	Degree *string `json:"degree"`

	// email
	// Required: true
	Email *string `json:"email"`

	// email confirmed
	// Example: false
	EmailConfirmed bool `json:"emailConfirmed,omitempty"`

	// expertise
	// Required: true
	Expertise *string `json:"expertise"`

	// name
	// Required: true
	Name *string `json:"name"`

	// verified
	// Example: false
	// Required: true
	Verified *bool `json:"verified"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDegree(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpertise(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateDegree(formats strfmt.Registry) error {

	if err := validate.Required("degree", "body", m.Degree); err != nil {
		return err
	}

	return nil
}

func (m *User) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *User) validateExpertise(formats strfmt.Registry) error {

	if err := validate.Required("expertise", "body", m.Expertise); err != nil {
		return err
	}

	return nil
}

func (m *User) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *User) validateVerified(formats strfmt.Registry) error {

	if err := validate.Required("verified", "body", m.Verified); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user based on context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
