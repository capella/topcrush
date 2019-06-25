// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/capella/topcrush/models"
)

// PostUserIDLocationCreatedCode is the HTTP code returned for type PostUserIDLocationCreated
const PostUserIDLocationCreatedCode int = 201

/*PostUserIDLocationCreated Created

swagger:response postUserIdLocationCreated
*/
type PostUserIDLocationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPostUserIDLocationCreated creates PostUserIDLocationCreated with default headers values
func NewPostUserIDLocationCreated() *PostUserIDLocationCreated {

	return &PostUserIDLocationCreated{}
}

// WithPayload adds the payload to the post user Id location created response
func (o *PostUserIDLocationCreated) WithPayload(payload *models.User) *PostUserIDLocationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user Id location created response
func (o *PostUserIDLocationCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserIDLocationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostUserIDLocationDefault generic error response

swagger:response postUserIdLocationDefault
*/
type PostUserIDLocationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostUserIDLocationDefault creates PostUserIDLocationDefault with default headers values
func NewPostUserIDLocationDefault(code int) *PostUserIDLocationDefault {
	if code <= 0 {
		code = 500
	}

	return &PostUserIDLocationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post user ID location default response
func (o *PostUserIDLocationDefault) WithStatusCode(code int) *PostUserIDLocationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post user ID location default response
func (o *PostUserIDLocationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post user ID location default response
func (o *PostUserIDLocationDefault) WithPayload(payload *models.Error) *PostUserIDLocationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user ID location default response
func (o *PostUserIDLocationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserIDLocationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}