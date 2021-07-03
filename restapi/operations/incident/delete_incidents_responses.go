// Code generated by go-swagger; DO NOT EDIT.

package incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// DeleteIncidentsOKCode is the HTTP code returned for type DeleteIncidentsOK
const DeleteIncidentsOKCode int = 200

/*DeleteIncidentsOK successful operation

swagger:response deleteIncidentsOK
*/
type DeleteIncidentsOK struct {

	/*
	  In: Body
	*/
	Payload *models.GoodResponse `json:"body,omitempty"`
}

// NewDeleteIncidentsOK creates DeleteIncidentsOK with default headers values
func NewDeleteIncidentsOK() *DeleteIncidentsOK {

	return &DeleteIncidentsOK{}
}

// WithPayload adds the payload to the delete incidents o k response
func (o *DeleteIncidentsOK) WithPayload(payload *models.GoodResponse) *DeleteIncidentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete incidents o k response
func (o *DeleteIncidentsOK) SetPayload(payload *models.GoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIncidentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteIncidentsUnauthorizedCode is the HTTP code returned for type DeleteIncidentsUnauthorized
const DeleteIncidentsUnauthorizedCode int = 401

/*DeleteIncidentsUnauthorized bad authorization token

swagger:response deleteIncidentsUnauthorized
*/
type DeleteIncidentsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewDeleteIncidentsUnauthorized creates DeleteIncidentsUnauthorized with default headers values
func NewDeleteIncidentsUnauthorized() *DeleteIncidentsUnauthorized {

	return &DeleteIncidentsUnauthorized{}
}

// WithPayload adds the payload to the delete incidents unauthorized response
func (o *DeleteIncidentsUnauthorized) WithPayload(payload *models.BadResponse) *DeleteIncidentsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete incidents unauthorized response
func (o *DeleteIncidentsUnauthorized) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIncidentsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteIncidentsNotFoundCode is the HTTP code returned for type DeleteIncidentsNotFound
const DeleteIncidentsNotFoundCode int = 404

/*DeleteIncidentsNotFound incident not found

swagger:response deleteIncidentsNotFound
*/
type DeleteIncidentsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewDeleteIncidentsNotFound creates DeleteIncidentsNotFound with default headers values
func NewDeleteIncidentsNotFound() *DeleteIncidentsNotFound {

	return &DeleteIncidentsNotFound{}
}

// WithPayload adds the payload to the delete incidents not found response
func (o *DeleteIncidentsNotFound) WithPayload(payload *models.BadResponse) *DeleteIncidentsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete incidents not found response
func (o *DeleteIncidentsNotFound) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteIncidentsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
