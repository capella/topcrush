// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/capella/topcrush/models"
)

// PutChatMessagesIDCreatedCode is the HTTP code returned for type PutChatMessagesIDCreated
const PutChatMessagesIDCreatedCode int = 201

/*PutChatMessagesIDCreated message submited

swagger:response putChatMessagesIdCreated
*/
type PutChatMessagesIDCreated struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPutChatMessagesIDCreated creates PutChatMessagesIDCreated with default headers values
func NewPutChatMessagesIDCreated() *PutChatMessagesIDCreated {

	return &PutChatMessagesIDCreated{}
}

// WithPayload adds the payload to the put chat messages Id created response
func (o *PutChatMessagesIDCreated) WithPayload(payload string) *PutChatMessagesIDCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put chat messages Id created response
func (o *PutChatMessagesIDCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChatMessagesIDCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PutChatMessagesIDDefault generic error response

swagger:response putChatMessagesIdDefault
*/
type PutChatMessagesIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutChatMessagesIDDefault creates PutChatMessagesIDDefault with default headers values
func NewPutChatMessagesIDDefault(code int) *PutChatMessagesIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PutChatMessagesIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put chat messages ID default response
func (o *PutChatMessagesIDDefault) WithStatusCode(code int) *PutChatMessagesIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put chat messages ID default response
func (o *PutChatMessagesIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put chat messages ID default response
func (o *PutChatMessagesIDDefault) WithPayload(payload *models.Error) *PutChatMessagesIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put chat messages ID default response
func (o *PutChatMessagesIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChatMessagesIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
