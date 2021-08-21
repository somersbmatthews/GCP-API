// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// BanMedicalExpertOKCode is the HTTP code returned for type BanMedicalExpertOK
const BanMedicalExpertOKCode int = 200

/*BanMedicalExpertOK successful operation

swagger:response banMedicalExpertOK
*/
type BanMedicalExpertOK struct {

	/*
	  In: Body
	*/
	Payload *models.GoodResponse `json:"body,omitempty"`
}

// NewBanMedicalExpertOK creates BanMedicalExpertOK with default headers values
func NewBanMedicalExpertOK() *BanMedicalExpertOK {

	return &BanMedicalExpertOK{}
}

// WithPayload adds the payload to the ban medical expert o k response
func (o *BanMedicalExpertOK) WithPayload(payload *models.GoodResponse) *BanMedicalExpertOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the ban medical expert o k response
func (o *BanMedicalExpertOK) SetPayload(payload *models.GoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BanMedicalExpertOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BanMedicalExpertUnauthorizedCode is the HTTP code returned for type BanMedicalExpertUnauthorized
const BanMedicalExpertUnauthorizedCode int = 401

/*BanMedicalExpertUnauthorized bad authorization token

swagger:response banMedicalExpertUnauthorized
*/
type BanMedicalExpertUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewBanMedicalExpertUnauthorized creates BanMedicalExpertUnauthorized with default headers values
func NewBanMedicalExpertUnauthorized() *BanMedicalExpertUnauthorized {

	return &BanMedicalExpertUnauthorized{}
}

// WithPayload adds the payload to the ban medical expert unauthorized response
func (o *BanMedicalExpertUnauthorized) WithPayload(payload *models.BadResponse) *BanMedicalExpertUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the ban medical expert unauthorized response
func (o *BanMedicalExpertUnauthorized) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BanMedicalExpertUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BanMedicalExpertNotFoundCode is the HTTP code returned for type BanMedicalExpertNotFound
const BanMedicalExpertNotFoundCode int = 404

/*BanMedicalExpertNotFound user not found

swagger:response banMedicalExpertNotFound
*/
type BanMedicalExpertNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewBanMedicalExpertNotFound creates BanMedicalExpertNotFound with default headers values
func NewBanMedicalExpertNotFound() *BanMedicalExpertNotFound {

	return &BanMedicalExpertNotFound{}
}

// WithPayload adds the payload to the ban medical expert not found response
func (o *BanMedicalExpertNotFound) WithPayload(payload *models.BadResponse) *BanMedicalExpertNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the ban medical expert not found response
func (o *BanMedicalExpertNotFound) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BanMedicalExpertNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
