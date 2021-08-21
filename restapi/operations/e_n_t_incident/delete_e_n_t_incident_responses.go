// Code generated by go-swagger; DO NOT EDIT.

package e_n_t_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// DeleteENTIncidentOKCode is the HTTP code returned for type DeleteENTIncidentOK
const DeleteENTIncidentOKCode int = 200

/*DeleteENTIncidentOK successful operation

swagger:response deleteENTIncidentOK
*/
type DeleteENTIncidentOK struct {

	/*
	  In: Body
	*/
	Payload *models.GoodResponse `json:"body,omitempty"`
}

// NewDeleteENTIncidentOK creates DeleteENTIncidentOK with default headers values
func NewDeleteENTIncidentOK() *DeleteENTIncidentOK {

	return &DeleteENTIncidentOK{}
}

// WithPayload adds the payload to the delete e n t incident o k response
func (o *DeleteENTIncidentOK) WithPayload(payload *models.GoodResponse) *DeleteENTIncidentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete e n t incident o k response
func (o *DeleteENTIncidentOK) SetPayload(payload *models.GoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteENTIncidentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteENTIncidentUnauthorizedCode is the HTTP code returned for type DeleteENTIncidentUnauthorized
const DeleteENTIncidentUnauthorizedCode int = 401

/*DeleteENTIncidentUnauthorized bad authorization token

swagger:response deleteENTIncidentUnauthorized
*/
type DeleteENTIncidentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewDeleteENTIncidentUnauthorized creates DeleteENTIncidentUnauthorized with default headers values
func NewDeleteENTIncidentUnauthorized() *DeleteENTIncidentUnauthorized {

	return &DeleteENTIncidentUnauthorized{}
}

// WithPayload adds the payload to the delete e n t incident unauthorized response
func (o *DeleteENTIncidentUnauthorized) WithPayload(payload *models.BadResponse) *DeleteENTIncidentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete e n t incident unauthorized response
func (o *DeleteENTIncidentUnauthorized) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteENTIncidentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteENTIncidentNotFoundCode is the HTTP code returned for type DeleteENTIncidentNotFound
const DeleteENTIncidentNotFoundCode int = 404

/*DeleteENTIncidentNotFound incident not found

swagger:response deleteENTIncidentNotFound
*/
type DeleteENTIncidentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewDeleteENTIncidentNotFound creates DeleteENTIncidentNotFound with default headers values
func NewDeleteENTIncidentNotFound() *DeleteENTIncidentNotFound {

	return &DeleteENTIncidentNotFound{}
}

// WithPayload adds the payload to the delete e n t incident not found response
func (o *DeleteENTIncidentNotFound) WithPayload(payload *models.BadResponse) *DeleteENTIncidentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete e n t incident not found response
func (o *DeleteENTIncidentNotFound) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteENTIncidentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}