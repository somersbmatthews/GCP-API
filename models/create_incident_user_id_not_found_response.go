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

// CreateIncidentUserIDNotFoundResponse create incident user Id not found response
//
// swagger:model CreateIncidentUserIdNotFoundResponse
type CreateIncidentUserIDNotFoundResponse struct {

	// anterior
	Anterior string `json:"Anterior,omitempty"`

	// approximate patient age
	ApproximatePatientAge string `json:"Approximate_Patient_Age,omitempty"`

	// created
	Created *bool `json:"Created,omitempty"`

	// date of incident
	DateOfIncident string `json:"Date_of_Incident,omitempty"`

	// gender
	Gender string `json:"Gender,omitempty"`

	// ID
	// Required: true
	ID *string `json:"ID"`

	// incident description
	IncidentDescription string `json:"Incident_Description,omitempty"`

	// largest length
	LargestLength string `json:"Largest_Length,omitempty"`

	// location of object
	LocationOfObject string `json:"Location_of_object,omitempty"`

	// long term prognosis
	LongTermPrognosis string `json:"Long-term_prognosis,omitempty"`

	// object basic shape
	ObjectBasicShape string `json:"Object_Basic_Shape,omitempty"`

	// object consistency
	ObjectConsistency string `json:"Object_Consistency,omitempty"`

	// the object is
	TheObjectIs string `json:"The_object_is,omitempty"`

	// what material is the object made of
	WhatMaterialIsTheObjectMadeOf string `json:"What_material_is_the_object_made_of,omitempty"`
}

// Validate validates this create incident user Id not found response
func (m *CreateIncidentUserIDNotFoundResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateIncidentUserIDNotFoundResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create incident user Id not found response based on context it is used
func (m *CreateIncidentUserIDNotFoundResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateIncidentUserIDNotFoundResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateIncidentUserIDNotFoundResponse) UnmarshalBinary(b []byte) error {
	var res CreateIncidentUserIDNotFoundResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
