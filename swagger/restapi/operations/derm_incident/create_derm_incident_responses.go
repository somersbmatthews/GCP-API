// Code generated by go-swagger; DO NOT EDIT.

package derm_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/swagger/models"
)

// CreateDermIncidentOKCode is the HTTP code returned for type CreateDermIncidentOK
const CreateDermIncidentOKCode int = 200

/*CreateDermIncidentOK successful operation

swagger:response createDermIncidentOK
*/
type CreateDermIncidentOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteIncidentGoodResponse `json:"body,omitempty"`
}

// NewCreateDermIncidentOK creates CreateDermIncidentOK with default headers values
func NewCreateDermIncidentOK() *CreateDermIncidentOK {

	return &CreateDermIncidentOK{}
}

// WithPayload adds the payload to the create derm incident o k response
func (o *CreateDermIncidentOK) WithPayload(payload *models.DeleteIncidentGoodResponse) *CreateDermIncidentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create derm incident o k response
func (o *CreateDermIncidentOK) SetPayload(payload *models.DeleteIncidentGoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDermIncidentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDermIncidentUnauthorizedCode is the HTTP code returned for type CreateDermIncidentUnauthorized
const CreateDermIncidentUnauthorizedCode int = 401

/*CreateDermIncidentUnauthorized bad authorization token

swagger:response createDermIncidentUnauthorized
*/
type CreateDermIncidentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.UpdateIncidentIncidentIDNotFoundResponse `json:"body,omitempty"`
}

// NewCreateDermIncidentUnauthorized creates CreateDermIncidentUnauthorized with default headers values
func NewCreateDermIncidentUnauthorized() *CreateDermIncidentUnauthorized {

	return &CreateDermIncidentUnauthorized{}
}

// WithPayload adds the payload to the create derm incident unauthorized response
func (o *CreateDermIncidentUnauthorized) WithPayload(payload *models.UpdateIncidentIncidentIDNotFoundResponse) *CreateDermIncidentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create derm incident unauthorized response
func (o *CreateDermIncidentUnauthorized) SetPayload(payload *models.UpdateIncidentIncidentIDNotFoundResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDermIncidentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDermIncidentNotFoundCode is the HTTP code returned for type CreateDermIncidentNotFound
const CreateDermIncidentNotFoundCode int = 404

/*CreateDermIncidentNotFound incident not found

swagger:response createDermIncidentNotFound
*/
type CreateDermIncidentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteIncidentIncidentIDNotFoundResponse `json:"body,omitempty"`
}

// NewCreateDermIncidentNotFound creates CreateDermIncidentNotFound with default headers values
func NewCreateDermIncidentNotFound() *CreateDermIncidentNotFound {

	return &CreateDermIncidentNotFound{}
}

// WithPayload adds the payload to the create derm incident not found response
func (o *CreateDermIncidentNotFound) WithPayload(payload *models.DeleteIncidentIncidentIDNotFoundResponse) *CreateDermIncidentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create derm incident not found response
func (o *CreateDermIncidentNotFound) SetPayload(payload *models.DeleteIncidentIncidentIDNotFoundResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDermIncidentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
