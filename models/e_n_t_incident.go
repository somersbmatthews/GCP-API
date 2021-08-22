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

	// ID
	// Required: true
	ID *string `json:"ID"`

	// age months
	// Required: true
	AgeMonths *string `json:"ageMonths"`

	// age years
	// Required: true
	AgeYears *string `json:"ageYears"`

	// anesthesia
	// Required: true
	Anesthesia *string `json:"anesthesia"`

	// complications
	// Required: true
	Complications []string `json:"complications"`

	// country
	// Required: true
	Country *string `json:"country"`

	// custom complications
	// Required: true
	CustomComplications []string `json:"customComplications"`

	// custom symptoms
	// Required: true
	CustomSymptoms []string `json:"customSymptoms"`

	// days until removal
	// Required: true
	DaysUntilRemoval *int64 `json:"daysUntilRemoval"`

	// device type
	// Required: true
	DeviceType *string `json:"deviceType"`

	// ease of removal
	// Required: true
	EaseOfRemoval *string `json:"easeOfRemoval"`

	// gender
	// Required: true
	Gender *string `json:"gender"`

	// hospital stay
	// Required: true
	HospitalStay *string `json:"hospitalStay"`

	// hours until removal
	// Required: true
	HoursUntilRemoval *int64 `json:"hoursUntilRemoval"`

	// incident description
	// Required: true
	IncidentDescription *string `json:"incidentDescription"`

	// minutes until removal
	// Required: true
	MinutesUntilRemoval *int64 `json:"minutesUntilRemoval"`

	// open surgery
	// Required: true
	OpenSurgery *string `json:"openSurgery"`

	// prognosis
	// Required: true
	Prognosis *string `json:"prognosis"`

	// removal strategy
	// Required: true
	RemovalStrategy *string `json:"removalStrategy"`

	// submitted
	Submitted bool `json:"submitted,omitempty"`

	// swallowed objects
	// Required: true
	SwallowedObjects []*SwallowedObject `json:"swallowedObjects"`

	// symptom severity
	// Required: true
	SymptomSeverity *string `json:"symptomSeverity"`

	// symptoms
	// Required: true
	Symptoms []string `json:"symptoms"`

	// was incident life threatening
	// Required: true
	WasIncidentLifeThreatening *string `json:"wasIncidentLifeThreatening"`

	// year
	// Required: true
	Year *string `json:"year"`
}

// Validate validates this e n t incident
func (m *ENTIncident) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgeMonths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAgeYears(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnesthesia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomComplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomSymptoms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaysUntilRemoval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEaseOfRemoval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHospitalStay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoursUntilRemoval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidentDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinutesUntilRemoval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenSurgery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrognosis(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemovalStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwallowedObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSymptomSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSymptoms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWasIncidentLifeThreatening(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYear(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ENTIncident) validateID(formats strfmt.Registry) error {

	if err := validate.Required("ID", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateAgeMonths(formats strfmt.Registry) error {

	if err := validate.Required("ageMonths", "body", m.AgeMonths); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateAgeYears(formats strfmt.Registry) error {

	if err := validate.Required("ageYears", "body", m.AgeYears); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateAnesthesia(formats strfmt.Registry) error {

	if err := validate.Required("anesthesia", "body", m.Anesthesia); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateComplications(formats strfmt.Registry) error {

	if err := validate.Required("complications", "body", m.Complications); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateCustomComplications(formats strfmt.Registry) error {

	if err := validate.Required("customComplications", "body", m.CustomComplications); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateCustomSymptoms(formats strfmt.Registry) error {

	if err := validate.Required("customSymptoms", "body", m.CustomSymptoms); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateDaysUntilRemoval(formats strfmt.Registry) error {

	if err := validate.Required("daysUntilRemoval", "body", m.DaysUntilRemoval); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("deviceType", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateEaseOfRemoval(formats strfmt.Registry) error {

	if err := validate.Required("easeOfRemoval", "body", m.EaseOfRemoval); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateGender(formats strfmt.Registry) error {

	if err := validate.Required("gender", "body", m.Gender); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateHospitalStay(formats strfmt.Registry) error {

	if err := validate.Required("hospitalStay", "body", m.HospitalStay); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateHoursUntilRemoval(formats strfmt.Registry) error {

	if err := validate.Required("hoursUntilRemoval", "body", m.HoursUntilRemoval); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateIncidentDescription(formats strfmt.Registry) error {

	if err := validate.Required("incidentDescription", "body", m.IncidentDescription); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateMinutesUntilRemoval(formats strfmt.Registry) error {

	if err := validate.Required("minutesUntilRemoval", "body", m.MinutesUntilRemoval); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateOpenSurgery(formats strfmt.Registry) error {

	if err := validate.Required("openSurgery", "body", m.OpenSurgery); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validatePrognosis(formats strfmt.Registry) error {

	if err := validate.Required("prognosis", "body", m.Prognosis); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateRemovalStrategy(formats strfmt.Registry) error {

	if err := validate.Required("removalStrategy", "body", m.RemovalStrategy); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateSwallowedObjects(formats strfmt.Registry) error {

	if err := validate.Required("swallowedObjects", "body", m.SwallowedObjects); err != nil {
		return err
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

func (m *ENTIncident) validateSymptomSeverity(formats strfmt.Registry) error {

	if err := validate.Required("symptomSeverity", "body", m.SymptomSeverity); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateSymptoms(formats strfmt.Registry) error {

	if err := validate.Required("symptoms", "body", m.Symptoms); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateWasIncidentLifeThreatening(formats strfmt.Registry) error {

	if err := validate.Required("wasIncidentLifeThreatening", "body", m.WasIncidentLifeThreatening); err != nil {
		return err
	}

	return nil
}

func (m *ENTIncident) validateYear(formats strfmt.Registry) error {

	if err := validate.Required("year", "body", m.Year); err != nil {
		return err
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
