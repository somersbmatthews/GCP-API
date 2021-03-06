// Code generated by go-swagger; DO NOT EDIT.

package medical_expert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// GetExpertOKCode is the HTTP code returned for type GetExpertOK
const GetExpertOKCode int = 200

/*GetExpertOK successful operation

swagger:response getExpertOK
*/
type GetExpertOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetExpertResponse `json:"body,omitempty"`
}

// NewGetExpertOK creates GetExpertOK with default headers values
func NewGetExpertOK() *GetExpertOK {

	return &GetExpertOK{}
}

// WithPayload adds the payload to the get expert o k response
func (o *GetExpertOK) WithPayload(payload *models.GetExpertResponse) *GetExpertOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get expert o k response
func (o *GetExpertOK) SetPayload(payload *models.GetExpertResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetExpertOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetExpertUnauthorizedCode is the HTTP code returned for type GetExpertUnauthorized
const GetExpertUnauthorizedCode int = 401

/*GetExpertUnauthorized bad authorization token

swagger:response getExpertUnauthorized
*/
type GetExpertUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewGetExpertUnauthorized creates GetExpertUnauthorized with default headers values
func NewGetExpertUnauthorized() *GetExpertUnauthorized {

	return &GetExpertUnauthorized{}
}

// WithPayload adds the payload to the get expert unauthorized response
func (o *GetExpertUnauthorized) WithPayload(payload *models.BadResponse) *GetExpertUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get expert unauthorized response
func (o *GetExpertUnauthorized) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetExpertUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetExpertNotFoundCode is the HTTP code returned for type GetExpertNotFound
const GetExpertNotFoundCode int = 404

/*GetExpertNotFound incident not found

swagger:response getExpertNotFound
*/
type GetExpertNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewGetExpertNotFound creates GetExpertNotFound with default headers values
func NewGetExpertNotFound() *GetExpertNotFound {

	return &GetExpertNotFound{}
}

// WithPayload adds the payload to the get expert not found response
func (o *GetExpertNotFound) WithPayload(payload *models.BadResponse) *GetExpertNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get expert not found response
func (o *GetExpertNotFound) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetExpertNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
