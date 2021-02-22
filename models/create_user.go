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

// CreateUser create user
//
// swagger:model CreateUser
type CreateUser struct {

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

// Validate validates this create user
func (m *CreateUser) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *CreateUser) validateDegree(formats strfmt.Registry) error {

	if err := validate.Required("degree", "body", m.Degree); err != nil {
		return err
	}

	return nil
}

func (m *CreateUser) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *CreateUser) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateUser) validateSpeciality(formats strfmt.Registry) error {

	if err := validate.Required("speciality", "body", m.Speciality); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create user based on context it is used
func (m *CreateUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateUser) UnmarshalBinary(b []byte) error {
	var res CreateUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
