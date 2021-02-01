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

// LoginGoodResponse login good response
//
// swagger:model LoginGoodResponse
type LoginGoodResponse struct {

	// degree
	// Required: true
	Degree *string `json:"degree"`

	// email
	Email string `json:"email,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// registered
	Registered bool `json:"registered,omitempty"`

	// speciality
	// Required: true
	Speciality *string `json:"speciality"`

	// user Id
	// Required: true
	UserID *string `json:"userId"`

	// verified
	// Required: true
	Verified bool `json:"verified"`
}

// Validate validates this login good response
func (m *LoginGoodResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDegree(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeciality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
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

func (m *LoginGoodResponse) validateDegree(formats strfmt.Registry) error {

	if err := validate.Required("degree", "body", m.Degree); err != nil {
		return err
	}

	return nil
}

func (m *LoginGoodResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *LoginGoodResponse) validateSpeciality(formats strfmt.Registry) error {

	if err := validate.Required("speciality", "body", m.Speciality); err != nil {
		return err
	}

	return nil
}

func (m *LoginGoodResponse) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *LoginGoodResponse) validateVerified(formats strfmt.Registry) error {

	if err := validate.Required("verified", "body", bool(m.Verified)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login good response based on context it is used
func (m *LoginGoodResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginGoodResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginGoodResponse) UnmarshalBinary(b []byte) error {
	var res LoginGoodResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
