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

// ENTIncident e n t incident
//
// swagger:model ENTIncident
type ENTIncident struct {

	// age
	Age string `json:"age,omitempty"`

	// anesthesia
	Anesthesia int64 `json:"anesthesia,omitempty"`

	// complications
	Complications []string `json:"complications"`

	// country
	Country string `json:"country,omitempty"`

	// custom complications
	CustomComplications []string `json:"customComplications"`

	// custom procedures
	CustomProcedures []string `json:"customProcedures"`

	// custom symptoms
	CustomSymptoms []string `json:"customSymptoms"`

	// gender
	Gender string `json:"gender,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// minutes until removal
	MinutesUntilRemoval string `json:"minutesUntilRemoval,omitempty"`

	// procedures
	Procedures []string `json:"procedures"`

	// prognosis
	Prognosis string `json:"prognosis,omitempty"`

	// removal setting
	RemovalSetting string `json:"removalSetting,omitempty"`

	// swallowed objects
	SwallowedObjects []*SwallowedObject `json:"swallowedObjects"`

	// symptom severity
	SymptomSeverity string `json:"symptomSeverity,omitempty"`

	// symptoms
	Symptoms []string `json:"symptoms"`
}

// Validate validates this e n t incident
func (m *ENTIncident) Validate(formats strfmt.Registry) error {
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

func (m *ENTIncident) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateSwallowedObjects(formats strfmt.Registry) error {
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
					return ve.ValidateName("swallowedObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this e n t incident based on the context it is used
func (m *ENTIncident) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSwallowedObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ENTIncident) contextValidateSwallowedObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SwallowedObjects); i++ {

		if m.SwallowedObjects[i] != nil {
			if err := m.SwallowedObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("swallowedObjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ENTIncident) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ENTIncident) UnmarshalBinary(b []byte) error {
	var res ENTIncident
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}