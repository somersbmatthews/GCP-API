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

// CreateUserBadResponse create user bad response
//
// swagger:model CreateUserBadResponse
type CreateUserBadResponse struct {

	// created
	// Example: false
	// Required: true
	Created *bool `json:"created"`

	// degree
	// Required: true
	Degree *string `json:"degree"`

	// email
	// Required: true
	Email *string `json:"email"`

	// name
	// Required: true
	Name *string `json:"name"`

	// speciality
	// Required: true
	Speciality *string `json:"speciality"`
}

// Validate validates this create user bad response
func (m *CreateUserBadResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDegree(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeciality(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateUserBadResponse) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *CreateUserBadResponse) validateDegree(formats strfmt.Registry) error {

	if err := validate.Required("degree", "body", m.Degree); err != nil {
		return err
	}

	return nil
}

func (m *CreateUserBadResponse) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *CreateUserBadResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateUserBadResponse) validateSpeciality(formats strfmt.Registry) error {

	if err := validate.Required("speciality", "body", m.Speciality); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create user bad response based on context it is used
func (m *CreateUserBadResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateUserBadResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUserBadResponse) UnmarshalBinary(b []byte) error {
	var res CreateUserBadResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
