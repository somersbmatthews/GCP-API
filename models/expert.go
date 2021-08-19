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

// Expert expert
//
// swagger:model Expert
type Expert struct {

	// degree
	// Required: true
	Degree *string `json:"degree"`

	// email
	// Required: true
	Email *string `json:"email"`

	// expertise
	// Required: true
	Expertise *string `json:"expertise"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this expert
func (m *Expert) Validate(formats strfmt.Registry) error {
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Expert) validateDegree(formats strfmt.Registry) error {

	if err := validate.Required("degree", "body", m.Degree); err != nil {
		return err
	}

	return nil
}

func (m *Expert) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *Expert) validateExpertise(formats strfmt.Registry) error {

	if err := validate.Required("expertise", "body", m.Expertise); err != nil {
		return err
	}

	return nil
}

func (m *Expert) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this expert based on context it is used
func (m *Expert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Expert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Expert) UnmarshalBinary(b []byte) error {
	var res Expert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
