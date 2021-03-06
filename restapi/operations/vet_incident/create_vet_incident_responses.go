// Code generated by go-swagger; DO NOT EDIT.

package vet_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// CreateVetIncidentOKCode is the HTTP code returned for type CreateVetIncidentOK
const CreateVetIncidentOKCode int = 200

/*CreateVetIncidentOK successful operation

swagger:response createVetIncidentOK
*/
type CreateVetIncidentOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteIncidentGoodResponse `json:"body,omitempty"`
}

// NewCreateVetIncidentOK creates CreateVetIncidentOK with default headers values
func NewCreateVetIncidentOK() *CreateVetIncidentOK {

	return &CreateVetIncidentOK{}
}

// WithPayload adds the payload to the create vet incident o k response
func (o *CreateVetIncidentOK) WithPayload(payload *models.DeleteIncidentGoodResponse) *CreateVetIncidentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create vet incident o k response
func (o *CreateVetIncidentOK) SetPayload(payload *models.DeleteIncidentGoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVetIncidentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateVetIncidentUnauthorizedCode is the HTTP code returned for type CreateVetIncidentUnauthorized
const CreateVetIncidentUnauthorizedCode int = 401

/*CreateVetIncidentUnauthorized bad authorization token

swagger:response createVetIncidentUnauthorized
*/
type CreateVetIncidentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewCreateVetIncidentUnauthorized creates CreateVetIncidentUnauthorized with default headers values
func NewCreateVetIncidentUnauthorized() *CreateVetIncidentUnauthorized {

	return &CreateVetIncidentUnauthorized{}
}

// WithPayload adds the payload to the create vet incident unauthorized response
func (o *CreateVetIncidentUnauthorized) WithPayload(payload *models.BadResponse) *CreateVetIncidentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create vet incident unauthorized response
func (o *CreateVetIncidentUnauthorized) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVetIncidentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateVetIncidentNotFoundCode is the HTTP code returned for type CreateVetIncidentNotFound
const CreateVetIncidentNotFoundCode int = 404

/*CreateVetIncidentNotFound incident not found

swagger:response createVetIncidentNotFound
*/
type CreateVetIncidentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewCreateVetIncidentNotFound creates CreateVetIncidentNotFound with default headers values
func NewCreateVetIncidentNotFound() *CreateVetIncidentNotFound {

	return &CreateVetIncidentNotFound{}
}

// WithPayload adds the payload to the create vet incident not found response
func (o *CreateVetIncidentNotFound) WithPayload(payload *models.BadResponse) *CreateVetIncidentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create vet incident not found response
func (o *CreateVetIncidentNotFound) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateVetIncidentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
