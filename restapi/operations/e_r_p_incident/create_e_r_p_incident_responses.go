// Code generated by go-swagger; DO NOT EDIT.

package e_r_p_incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// CreateERPIncidentOKCode is the HTTP code returned for type CreateERPIncidentOK
const CreateERPIncidentOKCode int = 200

/*CreateERPIncidentOK successful operation

swagger:response createERPIncidentOK
*/
type CreateERPIncidentOK struct {

	/*
	  In: Body
	*/
	Payload *models.DeleteIncidentGoodResponse `json:"body,omitempty"`
}

// NewCreateERPIncidentOK creates CreateERPIncidentOK with default headers values
func NewCreateERPIncidentOK() *CreateERPIncidentOK {

	return &CreateERPIncidentOK{}
}

// WithPayload adds the payload to the create e r p incident o k response
func (o *CreateERPIncidentOK) WithPayload(payload *models.DeleteIncidentGoodResponse) *CreateERPIncidentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create e r p incident o k response
func (o *CreateERPIncidentOK) SetPayload(payload *models.DeleteIncidentGoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateERPIncidentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateERPIncidentUnauthorizedCode is the HTTP code returned for type CreateERPIncidentUnauthorized
const CreateERPIncidentUnauthorizedCode int = 401

/*CreateERPIncidentUnauthorized bad authorization token

swagger:response createERPIncidentUnauthorized
*/
type CreateERPIncidentUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewCreateERPIncidentUnauthorized creates CreateERPIncidentUnauthorized with default headers values
func NewCreateERPIncidentUnauthorized() *CreateERPIncidentUnauthorized {

	return &CreateERPIncidentUnauthorized{}
}

// WithPayload adds the payload to the create e r p incident unauthorized response
func (o *CreateERPIncidentUnauthorized) WithPayload(payload *models.BadResponse) *CreateERPIncidentUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create e r p incident unauthorized response
func (o *CreateERPIncidentUnauthorized) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateERPIncidentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateERPIncidentNotFoundCode is the HTTP code returned for type CreateERPIncidentNotFound
const CreateERPIncidentNotFoundCode int = 404

/*CreateERPIncidentNotFound incident not found

swagger:response createERPIncidentNotFound
*/
type CreateERPIncidentNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.BadResponse `json:"body,omitempty"`
}

// NewCreateERPIncidentNotFound creates CreateERPIncidentNotFound with default headers values
func NewCreateERPIncidentNotFound() *CreateERPIncidentNotFound {

	return &CreateERPIncidentNotFound{}
}

// WithPayload adds the payload to the create e r p incident not found response
func (o *CreateERPIncidentNotFound) WithPayload(payload *models.BadResponse) *CreateERPIncidentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create e r p incident not found response
func (o *CreateERPIncidentNotFound) SetPayload(payload *models.BadResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateERPIncidentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
