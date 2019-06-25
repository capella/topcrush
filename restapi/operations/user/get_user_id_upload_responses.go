// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/capella/topcrush/models"
)

// GetUserIDUploadOKCode is the HTTP code returned for type GetUserIDUploadOK
const GetUserIDUploadOKCode int = 200

/*GetUserIDUploadOK produce a pre-signed link for upload

swagger:response getUserIdUploadOK
*/
type GetUserIDUploadOK struct {

	/*
	  In: Body
	*/
	Payload *models.PreSignedLink `json:"body,omitempty"`
}

// NewGetUserIDUploadOK creates GetUserIDUploadOK with default headers values
func NewGetUserIDUploadOK() *GetUserIDUploadOK {

	return &GetUserIDUploadOK{}
}

// WithPayload adds the payload to the get user Id upload o k response
func (o *GetUserIDUploadOK) WithPayload(payload *models.PreSignedLink) *GetUserIDUploadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user Id upload o k response
func (o *GetUserIDUploadOK) SetPayload(payload *models.PreSignedLink) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserIDUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUserIDUploadDefault generic error response

swagger:response getUserIdUploadDefault
*/
type GetUserIDUploadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserIDUploadDefault creates GetUserIDUploadDefault with default headers values
func NewGetUserIDUploadDefault(code int) *GetUserIDUploadDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserIDUploadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user ID upload default response
func (o *GetUserIDUploadDefault) WithStatusCode(code int) *GetUserIDUploadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user ID upload default response
func (o *GetUserIDUploadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get user ID upload default response
func (o *GetUserIDUploadDefault) WithPayload(payload *models.Error) *GetUserIDUploadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user ID upload default response
func (o *GetUserIDUploadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserIDUploadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}