// Code generated by go-swagger; DO NOT EDIT.

package incident

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gircapp/api/models"
)

// UpdateIncidentsOKCode is the HTTP code returned for type UpdateIncidentsOK
const UpdateIncidentsOKCode int = 200

/*UpdateIncidentsOK successful operation

swagger:response updateIncidentsOK
*/
type UpdateIncidentsOK struct {

	/*
	  In: Body
	*/
	Payload *models.UpdateIncidentGoodResponse `json:"body,omitempty"`
}

// NewUpdateIncidentsOK creates UpdateIncidentsOK with default headers values
func NewUpdateIncidentsOK() *UpdateIncidentsOK {

	return &UpdateIncidentsOK{}
}

// WithPayload adds the payload to the update incidents o k response
func (o *UpdateIncidentsOK) WithPayload(payload *models.UpdateIncidentGoodResponse) *UpdateIncidentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update incidents o k response
func (o *UpdateIncidentsOK) SetPayload(payload *models.UpdateIncidentGoodResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateIncidentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateIncidentsBadRequestCode is the HTTP code returned for type UpdateIncidentsBadRequest
const UpdateIncidentsBadRequestCode int = 400

/*UpdateIncidentsBadRequest invalid incident data

swagger:response updateIncidentsBadRequest
*/
type UpdateIncidentsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.UpdateIncidentIncidentIDNotFoundResponse `json:"body,omitempty"`
}

// NewUpdateIncidentsBadRequest creates UpdateIncidentsBadRequest with default headers values
func NewUpdateIncidentsBadRequest() *UpdateIncidentsBadRequest {

	return &UpdateIncidentsBadRequest{}
}

// WithPayload adds the payload to the update incidents bad request response
func (o *UpdateIncidentsBadRequest) WithPayload(payload *models.UpdateIncidentIncidentIDNotFoundResponse) *UpdateIncidentsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update incidents bad request response
func (o *UpdateIncidentsBadRequest) SetPayload(payload *models.UpdateIncidentIncidentIDNotFoundResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateIncidentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateIncidentsNotFoundCode is the HTTP code returned for type UpdateIncidentsNotFound
const UpdateIncidentsNotFoundCode int = 404

/*UpdateIncidentsNotFound incident not found

swagger:response updateIncidentsNotFound
*/
type UpdateIncidentsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.UpdateIncidentIncidentIDNotFoundResponse `json:"body,omitempty"`
}

// NewUpdateIncidentsNotFound creates UpdateIncidentsNotFound with default headers values
func NewUpdateIncidentsNotFound() *UpdateIncidentsNotFound {

	return &UpdateIncidentsNotFound{}
}

// WithPayload adds the payload to the update incidents not found response
func (o *UpdateIncidentsNotFound) WithPayload(payload *models.UpdateIncidentIncidentIDNotFoundResponse) *UpdateIncidentsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update incidents not found response
func (o *UpdateIncidentsNotFound) SetPayload(payload *models.UpdateIncidentIncidentIDNotFoundResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateIncidentsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
