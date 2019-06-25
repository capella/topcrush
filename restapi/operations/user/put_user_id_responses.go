// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/capella/topcrush/models"
)

// PutUserIDCreatedCode is the HTTP code returned for type PutUserIDCreated
const PutUserIDCreatedCode int = 201

/*PutUserIDCreated Uploaded

swagger:response putUserIdCreated
*/
type PutUserIDCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPutUserIDCreated creates PutUserIDCreated with default headers values
func NewPutUserIDCreated() *PutUserIDCreated {

	return &PutUserIDCreated{}
}

// WithPayload adds the payload to the put user Id created response
func (o *PutUserIDCreated) WithPayload(payload *models.User) *PutUserIDCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put user Id created response
func (o *PutUserIDCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUserIDCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutUserIDForbiddenCode is the HTTP code returned for type PutUserIDForbidden
const PutUserIDForbiddenCode int = 403

/*PutUserIDForbidden Forbidde

swagger:response putUserIdForbidden
*/
type PutUserIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutUserIDForbidden creates PutUserIDForbidden with default headers values
func NewPutUserIDForbidden() *PutUserIDForbidden {

	return &PutUserIDForbidden{}
}

// WithPayload adds the payload to the put user Id forbidden response
func (o *PutUserIDForbidden) WithPayload(payload *models.Error) *PutUserIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put user Id forbidden response
func (o *PutUserIDForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUserIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutUserIDDefault Generic error

swagger:response putUserIdDefault
*/
type PutUserIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutUserIDDefault creates PutUserIDDefault with default headers values
func NewPutUserIDDefault(code int) *PutUserIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PutUserIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put user ID default response
func (o *PutUserIDDefault) WithStatusCode(code int) *PutUserIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put user ID default response
func (o *PutUserIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put user ID default response
func (o *PutUserIDDefault) WithPayload(payload *models.Error) *PutUserIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put user ID default response
func (o *PutUserIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUserIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
