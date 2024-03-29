// Code generated by go-swagger; DO NOT EDIT.

package slide

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/capella/topcrush/models"
)

// PutSlideSuperlikeCreatedCode is the HTTP code returned for type PutSlideSuperlikeCreated
const PutSlideSuperlikeCreatedCode int = 201

/*PutSlideSuperlikeCreated superlike worked

swagger:response putSlideSuperlikeCreated
*/
type PutSlideSuperlikeCreated struct {

	/*
	  In: Body
	*/
	Payload models.Success `json:"body,omitempty"`
}

// NewPutSlideSuperlikeCreated creates PutSlideSuperlikeCreated with default headers values
func NewPutSlideSuperlikeCreated() *PutSlideSuperlikeCreated {

	return &PutSlideSuperlikeCreated{}
}

// WithPayload adds the payload to the put slide superlike created response
func (o *PutSlideSuperlikeCreated) WithPayload(payload models.Success) *PutSlideSuperlikeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put slide superlike created response
func (o *PutSlideSuperlikeCreated) SetPayload(payload models.Success) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSlideSuperlikeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutSlideSuperlikePaymentRequiredCode is the HTTP code returned for type PutSlideSuperlikePaymentRequired
const PutSlideSuperlikePaymentRequiredCode int = 402

/*PutSlideSuperlikePaymentRequired no more superlikes avaliable

swagger:response putSlideSuperlikePaymentRequired
*/
type PutSlideSuperlikePaymentRequired struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutSlideSuperlikePaymentRequired creates PutSlideSuperlikePaymentRequired with default headers values
func NewPutSlideSuperlikePaymentRequired() *PutSlideSuperlikePaymentRequired {

	return &PutSlideSuperlikePaymentRequired{}
}

// WithPayload adds the payload to the put slide superlike payment required response
func (o *PutSlideSuperlikePaymentRequired) WithPayload(payload *models.Error) *PutSlideSuperlikePaymentRequired {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put slide superlike payment required response
func (o *PutSlideSuperlikePaymentRequired) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSlideSuperlikePaymentRequired) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(402)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutSlideSuperlikeDefault generic error response

swagger:response putSlideSuperlikeDefault
*/
type PutSlideSuperlikeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutSlideSuperlikeDefault creates PutSlideSuperlikeDefault with default headers values
func NewPutSlideSuperlikeDefault(code int) *PutSlideSuperlikeDefault {
	if code <= 0 {
		code = 500
	}

	return &PutSlideSuperlikeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put slide superlike default response
func (o *PutSlideSuperlikeDefault) WithStatusCode(code int) *PutSlideSuperlikeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put slide superlike default response
func (o *PutSlideSuperlikeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put slide superlike default response
func (o *PutSlideSuperlikeDefault) WithPayload(payload *models.Error) *PutSlideSuperlikeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put slide superlike default response
func (o *PutSlideSuperlikeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSlideSuperlikeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
