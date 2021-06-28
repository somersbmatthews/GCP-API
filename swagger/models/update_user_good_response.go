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

// UpdateUserGoodResponse update user good response
//
// swagger:model UpdateUserGoodResponse
type UpdateUserGoodResponse struct {

	// degree
	// Required: true
	Degree *string `json:"degree"`

	// email
	// Required: true
	Email *string `json:"email"`

	// name
	// Required: true
	Name *string `json:"name"`

	// specialty
	// Required: true
	Specialty *string `json:"specialty"`

	// updated
	// Required: true
	Updated *bool `json:"updated"`

	// user Id
	// Required: true
	UserID *string `json:"userId"`

	// verified
	// Required: true
	Verified bool `json:"verified"`
}

// Validate validates this update user good response
func (m *UpdateUserGoodResponse) Validate(formats strfmt.Registry) error {
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

	if err := m.validateSpecialty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
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

func (m *UpdateUserGoodResponse) validateDegree(formats strfmt.Registry) error {

	if err := validate.Required("degree", "body", m.Degree); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUserGoodResponse) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUserGoodResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUserGoodResponse) validateSpecialty(formats strfmt.Registry) error {

	if err := validate.Required("specialty", "body", m.Specialty); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUserGoodResponse) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("updated", "body", m.Updated); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUserGoodResponse) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *UpdateUserGoodResponse) validateVerified(formats strfmt.Registry) error {

	if err := validate.Required("verified", "body", bool(m.Verified)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update user good response based on context it is used
func (m *UpdateUserGoodResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUserGoodResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserGoodResponse) UnmarshalBinary(b []byte) error {
	var res UpdateUserGoodResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
