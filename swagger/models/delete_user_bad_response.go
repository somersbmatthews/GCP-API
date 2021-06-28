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

// DeleteUserBadResponse delete user bad response
//
// swagger:model DeleteUserBadResponse
type DeleteUserBadResponse struct {

	// deleted
	// Example: false
	// Required: true
	Deleted *bool `json:"deleted"`

	// user Id
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this delete user bad response
func (m *DeleteUserBadResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteUserBadResponse) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", m.Deleted); err != nil {
		return err
	}

	return nil
}

func (m *DeleteUserBadResponse) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete user bad response based on context it is used
func (m *DeleteUserBadResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteUserBadResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteUserBadResponse) UnmarshalBinary(b []byte) error {
	var res DeleteUserBadResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
