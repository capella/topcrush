// Code generated by go-swagger; DO NOT EDIT.

package slide

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/capella/topcrush/models"
)

// PutSlideBoastCreatedCode is the HTTP code returned for type PutSlideBoastCreated
const PutSlideBoastCreatedCode int = 201

/*PutSlideBoastCreated boast worked

swagger:response putSlideBoastCreated
*/
type PutSlideBoastCreated struct {

	/*
	  In: Body
	*/
	Payload models.Success `json:"body,omitempty"`
}

// NewPutSlideBoastCreated creates PutSlideBoastCreated with default headers values
func NewPutSlideBoastCreated() *PutSlideBoastCreated {

	return &PutSlideBoastCreated{}
}

// WithPayload adds the payload to the put slide boast created response
func (o *PutSlideBoastCreated) WithPayload(payload models.Success) *PutSlideBoastCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put slide boast created response
func (o *PutSlideBoastCreated) SetPayload(payload models.Success) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSlideBoastCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutSlideBoastPaymentRequiredCode is the HTTP code returned for type PutSlideBoastPaymentRequired
const PutSlideBoastPaymentRequiredCode int = 402

/*PutSlideBoastPaymentRequired no more boasts

swagger:response putSlideBoastPaymentRequired
*/
type PutSlideBoastPaymentRequired struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutSlideBoastPaymentRequired creates PutSlideBoastPaymentRequired with default headers values
func NewPutSlideBoastPaymentRequired() *PutSlideBoastPaymentRequired {

	return &PutSlideBoastPaymentRequired{}
}

// WithPayload adds the payload to the put slide boast payment required response
func (o *PutSlideBoastPaymentRequired) WithPayload(payload *models.Error) *PutSlideBoastPaymentRequired {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put slide boast payment required response
func (o *PutSlideBoastPaymentRequired) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSlideBoastPaymentRequired) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(402)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutSlideBoastDefault generic error response

swagger:response putSlideBoastDefault
*/
type PutSlideBoastDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutSlideBoastDefault creates PutSlideBoastDefault with default headers values
func NewPutSlideBoastDefault(code int) *PutSlideBoastDefault {
	if code <= 0 {
		code = 500
	}

	return &PutSlideBoastDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put slide boast default response
func (o *PutSlideBoastDefault) WithStatusCode(code int) *PutSlideBoastDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put slide boast default response
func (o *PutSlideBoastDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put slide boast default response
func (o *PutSlideBoastDefault) WithPayload(payload *models.Error) *PutSlideBoastDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put slide boast default response
func (o *PutSlideBoastDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSlideBoastDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}