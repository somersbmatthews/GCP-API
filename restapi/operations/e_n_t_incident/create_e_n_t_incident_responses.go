// Code generated by go-swagger; DO NOT EDIT.

package e_n_t_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// CreateENTIncidentOKCode is the HTTP code returned for type CreateENTIncidentOK
const CreateENTIncidentOKCode int = 200

/*CreateENTIncidentOK successful operation

swagger:response createENTIncidentOK
*/
type CreateENTIncidentOK struct {

	/*
	  In: Body
	*/
	Payload *models.GoodResponse `json:"body,omitempty"`
}

// NewCreateENTIncidentOK creates CreateENTIncidentOK with default headers values
func NewCreateENTIncidentOK() *CreateENTIncidentOK {

	return &CreateENTIncidentOK{}
}

// WithPayload adds the payload to the create e n t incident o k response
func (o *CreateENTIncidentOK) WithPayload(payload *models.GoodResponse) *CreateENTIncidentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create e n t incident o k response
func (o *CreateENTIncidentOK) SetPayload(payload *models.GoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateENTIncidentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateENTIncidentUnauthorizedCode is the HTTP code returned for type CreateENTIncidentUnauthorized
const CreateENTIncidentUnauthorizedCode int = 401

/*CreateENTIncidentUnauthorized bad authorization token

swagger:response createENTIncidentUnauthorized
*/
type CreateENTIncidentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewCreateENTIncidentUnauthorized creates CreateENTIncidentUnauthorized with default headers values
func NewCreateENTIncidentUnauthorized() *CreateENTIncidentUnauthorized {

	return &CreateENTIncidentUnauthorized{}
}

// WithPayload adds the payload to the create e n t incident unauthorized response
func (o *CreateENTIncidentUnauthorized) WithPayload(payload *models.BadResponse) *CreateENTIncidentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create e n t incident unauthorized response
func (o *CreateENTIncidentUnauthorized) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateENTIncidentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateENTIncidentNotFoundCode is the HTTP code returned for type CreateENTIncidentNotFound
const CreateENTIncidentNotFoundCode int = 404

/*CreateENTIncidentNotFound incident not found

swagger:response createENTIncidentNotFound
*/
type CreateENTIncidentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewCreateENTIncidentNotFound creates CreateENTIncidentNotFound with default headers values
func NewCreateENTIncidentNotFound() *CreateENTIncidentNotFound {

	return &CreateENTIncidentNotFound{}
}

// WithPayload adds the payload to the create e n t incident not found response
func (o *CreateENTIncidentNotFound) WithPayload(payload *models.BadResponse) *CreateENTIncidentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create e n t incident not found response
func (o *CreateENTIncidentNotFound) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateENTIncidentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
