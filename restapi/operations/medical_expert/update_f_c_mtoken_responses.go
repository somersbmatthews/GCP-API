// Code generated by go-swagger; DO NOT EDIT.

package medical_expert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// UpdateFCMtokenOKCode is the HTTP code returned for type UpdateFCMtokenOK
const UpdateFCMtokenOKCode int = 200

/*UpdateFCMtokenOK successful operation

swagger:response updateFCMtokenOK
*/
type UpdateFCMtokenOK struct {

	/*
	  In: Body
	*/
	Payload *models.GoodResponse `json:"body,omitempty"`
}

// NewUpdateFCMtokenOK creates UpdateFCMtokenOK with default headers values
func NewUpdateFCMtokenOK() *UpdateFCMtokenOK {

	return &UpdateFCMtokenOK{}
}

// WithPayload adds the payload to the update f c mtoken o k response
func (o *UpdateFCMtokenOK) WithPayload(payload *models.GoodResponse) *UpdateFCMtokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update f c mtoken o k response
func (o *UpdateFCMtokenOK) SetPayload(payload *models.GoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateFCMtokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateFCMtokenUnauthorizedCode is the HTTP code returned for type UpdateFCMtokenUnauthorized
const UpdateFCMtokenUnauthorizedCode int = 401

/*UpdateFCMtokenUnauthorized bad authorization token

swagger:response updateFCMtokenUnauthorized
*/
type UpdateFCMtokenUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewUpdateFCMtokenUnauthorized creates UpdateFCMtokenUnauthorized with default headers values
func NewUpdateFCMtokenUnauthorized() *UpdateFCMtokenUnauthorized {

	return &UpdateFCMtokenUnauthorized{}
}

// WithPayload adds the payload to the update f c mtoken unauthorized response
func (o *UpdateFCMtokenUnauthorized) WithPayload(payload *models.BadResponse) *UpdateFCMtokenUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update f c mtoken unauthorized response
func (o *UpdateFCMtokenUnauthorized) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateFCMtokenUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateFCMtokenNotFoundCode is the HTTP code returned for type UpdateFCMtokenNotFound
const UpdateFCMtokenNotFoundCode int = 404

/*UpdateFCMtokenNotFound user not found

swagger:response updateFCMtokenNotFound
*/
type UpdateFCMtokenNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewUpdateFCMtokenNotFound creates UpdateFCMtokenNotFound with default headers values
func NewUpdateFCMtokenNotFound() *UpdateFCMtokenNotFound {

	return &UpdateFCMtokenNotFound{}
}

// WithPayload adds the payload to the update f c mtoken not found response
func (o *UpdateFCMtokenNotFound) WithPayload(payload *models.BadResponse) *UpdateFCMtokenNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update f c mtoken not found response
func (o *UpdateFCMtokenNotFound) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateFCMtokenNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}