// Code generated by go-swagger; DO NOT EDIT.

package fire_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// CreateFireIncidentOKCode is the HTTP code returned for type CreateFireIncidentOK
const CreateFireIncidentOKCode int = 200

/*CreateFireIncidentOK successful operation

swagger:response createFireIncidentOK
*/
type CreateFireIncidentOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteIncidentGoodResponse `json:"body,omitempty"`
}

// NewCreateFireIncidentOK creates CreateFireIncidentOK with default headers values
func NewCreateFireIncidentOK() *CreateFireIncidentOK {

	return &CreateFireIncidentOK{}
}

// WithPayload adds the payload to the create fire incident o k response
func (o *CreateFireIncidentOK) WithPayload(payload *models.DeleteIncidentGoodResponse) *CreateFireIncidentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create fire incident o k response
func (o *CreateFireIncidentOK) SetPayload(payload *models.DeleteIncidentGoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFireIncidentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateFireIncidentUnauthorizedCode is the HTTP code returned for type CreateFireIncidentUnauthorized
const CreateFireIncidentUnauthorizedCode int = 401

/*CreateFireIncidentUnauthorized bad authorization token

swagger:response createFireIncidentUnauthorized
*/
type CreateFireIncidentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.UpdateIncidentIncidentIDNotFoundResponse `json:"body,omitempty"`
}

// NewCreateFireIncidentUnauthorized creates CreateFireIncidentUnauthorized with default headers values
func NewCreateFireIncidentUnauthorized() *CreateFireIncidentUnauthorized {

	return &CreateFireIncidentUnauthorized{}
}

// WithPayload adds the payload to the create fire incident unauthorized response
func (o *CreateFireIncidentUnauthorized) WithPayload(payload *models.UpdateIncidentIncidentIDNotFoundResponse) *CreateFireIncidentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create fire incident unauthorized response
func (o *CreateFireIncidentUnauthorized) SetPayload(payload *models.UpdateIncidentIncidentIDNotFoundResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFireIncidentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateFireIncidentNotFoundCode is the HTTP code returned for type CreateFireIncidentNotFound
const CreateFireIncidentNotFoundCode int = 404

/*CreateFireIncidentNotFound incident not found

swagger:response createFireIncidentNotFound
*/
type CreateFireIncidentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteIncidentIncidentIDNotFoundResponse `json:"body,omitempty"`
}

// NewCreateFireIncidentNotFound creates CreateFireIncidentNotFound with default headers values
func NewCreateFireIncidentNotFound() *CreateFireIncidentNotFound {

	return &CreateFireIncidentNotFound{}
}

// WithPayload adds the payload to the create fire incident not found response
func (o *CreateFireIncidentNotFound) WithPayload(payload *models.DeleteIncidentIncidentIDNotFoundResponse) *CreateFireIncidentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create fire incident not found response
func (o *CreateFireIncidentNotFound) SetPayload(payload *models.DeleteIncidentIncidentIDNotFoundResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateFireIncidentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
