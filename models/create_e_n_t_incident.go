// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateENTIncident create e n t incident
//
// swagger:model CreateENTIncident
type CreateENTIncident struct {

	// age
	Age string `json:"Age,omitempty"`

	// anesthesia
	Anesthesia int64 `json:"Anesthesia,omitempty"`

	// complications
	Complications []int64 `json:"Complications"`

	// country
	Country string `json:"Country,omitempty"`

	// custom complications
	CustomComplications []string `json:"CustomComplications"`

	// custom procedures
	CustomProcedures []string `json:"CustomProcedures"`

	// custom symptoms
	CustomSymptoms []string `json:"CustomSymptoms"`

	// date
	Date string `json:"Date,omitempty"`

	// gender
	Gender string `json:"Gender,omitempty"`

	// ID
	// Required: true
	ID *string `json:"ID"`

	// info
	Info string `json:"Info,omitempty"`

	// procedures
	Procedures []int64 `json:"Procedures"`

	// prognosis
	Prognosis string `json:"Prognosis,omitempty"`

	// removal setting
	RemovalSetting string `json:"RemovalSetting,omitempty"`

	// swallowed objects
	SwallowedObjects []*CreateSwallowedObject `json:"SwallowedObjects"`

	// symptom severity
	SymptomSeverity int64 `json:"SymptomSeverity,omitempty"`

	// symptoms
	Symptoms []int64 `json:"Symptoms"`

	// time until removal
	TimeUntilRemoval int64 `json:"TimeUntilRemoval,omitempty"`
}

// Validate validates this create e n t incident
func (m *CreateENTIncident) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwallowedObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateENTIncident) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CreateENTIncident) validateSwallowedObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.SwallowedObjects) { // not required
		return nil
	}

	for i := 0; i < len(m.SwallowedObjects); i++ {
		if swag.IsZero(m.SwallowedObjects[i]) { // not required
			continue
		}

		if m.SwallowedObjects[i] != nil {
			if err := m.SwallowedObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SwallowedObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create e n t incident based on the context it is used
func (m *CreateENTIncident) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSwallowedObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateENTIncident) contextValidateSwallowedObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SwallowedObjects); i++ {

		if m.SwallowedObjects[i] != nil {
			if err := m.SwallowedObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SwallowedObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateENTIncident) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateENTIncident) UnmarshalBinary(b []byte) error {
	var res CreateENTIncident
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
