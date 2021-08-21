// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// VerifyExpertOKCode is the HTTP code returned for type VerifyExpertOK
const VerifyExpertOKCode int = 200

/*VerifyExpertOK successful operation

swagger:response verifyExpertOK
*/
type VerifyExpertOK struct {

	/*
	  In: Body
	*/
	Payload *models.GoodResponse `json:"body,omitempty"`
}

// NewVerifyExpertOK creates VerifyExpertOK with default headers values
func NewVerifyExpertOK() *VerifyExpertOK {

	return &VerifyExpertOK{}
}

// WithPayload adds the payload to the verify expert o k response
func (o *VerifyExpertOK) WithPayload(payload *models.GoodResponse) *VerifyExpertOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verify expert o k response
func (o *VerifyExpertOK) SetPayload(payload *models.GoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerifyExpertOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VerifyExpertUnauthorizedCode is the HTTP code returned for type VerifyExpertUnauthorized
const VerifyExpertUnauthorizedCode int = 401

/*VerifyExpertUnauthorized bad authorization token

swagger:response verifyExpertUnauthorized
*/
type VerifyExpertUnauthorized struct {
}

// NewVerifyExpertUnauthorized creates VerifyExpertUnauthorized with default headers values
func NewVerifyExpertUnauthorized() *VerifyExpertUnauthorized {

	return &VerifyExpertUnauthorized{}
}

// WriteResponse to the client
func (o *VerifyExpertUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// VerifyExpertNotFoundCode is the HTTP code returned for type VerifyExpertNotFound
const VerifyExpertNotFoundCode int = 404

/*VerifyExpertNotFound user not found

swagger:response verifyExpertNotFound
*/
type VerifyExpertNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewVerifyExpertNotFound creates VerifyExpertNotFound with default headers values
func NewVerifyExpertNotFound() *VerifyExpertNotFound {

	return &VerifyExpertNotFound{}
}

// WithPayload adds the payload to the verify expert not found response
func (o *VerifyExpertNotFound) WithPayload(payload *models.BadResponse) *VerifyExpertNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the verify expert not found response
func (o *VerifyExpertNotFound) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VerifyExpertNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
