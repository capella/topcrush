// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/capella/topcrush/models"
)

// GetChatMessagesIDMessageIndexOKCode is the HTTP code returned for type GetChatMessagesIDMessageIndexOK
const GetChatMessagesIDMessageIndexOKCode int = 200

/*GetChatMessagesIDMessageIndexOK list of messages

swagger:response getChatMessagesIdMessageIndexOK
*/
type GetChatMessagesIDMessageIndexOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Message `json:"body,omitempty"`
}

// NewGetChatMessagesIDMessageIndexOK creates GetChatMessagesIDMessageIndexOK with default headers values
func NewGetChatMessagesIDMessageIndexOK() *GetChatMessagesIDMessageIndexOK {

	return &GetChatMessagesIDMessageIndexOK{}
}

// WithPayload adds the payload to the get chat messages Id message index o k response
func (o *GetChatMessagesIDMessageIndexOK) WithPayload(payload []*models.Message) *GetChatMessagesIDMessageIndexOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chat messages Id message index o k response
func (o *GetChatMessagesIDMessageIndexOK) SetPayload(payload []*models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChatMessagesIDMessageIndexOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Message, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetChatMessagesIDMessageIndexDefault generic error response

swagger:response getChatMessagesIdMessageIndexDefault
*/
type GetChatMessagesIDMessageIndexDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetChatMessagesIDMessageIndexDefault creates GetChatMessagesIDMessageIndexDefault with default headers values
func NewGetChatMessagesIDMessageIndexDefault(code int) *GetChatMessagesIDMessageIndexDefault {
	if code <= 0 {
		code = 500
	}

	return &GetChatMessagesIDMessageIndexDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get chat messages ID message index default response
func (o *GetChatMessagesIDMessageIndexDefault) WithStatusCode(code int) *GetChatMessagesIDMessageIndexDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get chat messages ID message index default response
func (o *GetChatMessagesIDMessageIndexDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get chat messages ID message index default response
func (o *GetChatMessagesIDMessageIndexDefault) WithPayload(payload *models.Error) *GetChatMessagesIDMessageIndexDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chat messages ID message index default response
func (o *GetChatMessagesIDMessageIndexDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChatMessagesIDMessageIndexDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
