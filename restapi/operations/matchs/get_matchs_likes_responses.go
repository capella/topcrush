// Code generated by go-swagger; DO NOT EDIT.

package matchs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/capella/topcrush/models"
)

// GetMatchsLikesCreatedCode is the HTTP code returned for type GetMatchsLikesCreated
const GetMatchsLikesCreatedCode int = 201

/*GetMatchsLikesCreated list of users

swagger:response getMatchsLikesCreated
*/
type GetMatchsLikesCreated struct {

	/*
	  In: Body
	*/
	Payload []*models.UserPublic `json:"body,omitempty"`
}

// NewGetMatchsLikesCreated creates GetMatchsLikesCreated with default headers values
func NewGetMatchsLikesCreated() *GetMatchsLikesCreated {

	return &GetMatchsLikesCreated{}
}

// WithPayload adds the payload to the get matchs likes created response
func (o *GetMatchsLikesCreated) WithPayload(payload []*models.UserPublic) *GetMatchsLikesCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get matchs likes created response
func (o *GetMatchsLikesCreated) SetPayload(payload []*models.UserPublic) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMatchsLikesCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.UserPublic, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetMatchsLikesPaymentRequiredCode is the HTTP code returned for type GetMatchsLikesPaymentRequired
const GetMatchsLikesPaymentRequiredCode int = 402

/*GetMatchsLikesPaymentRequired subscription required

swagger:response getMatchsLikesPaymentRequired
*/
type GetMatchsLikesPaymentRequired struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMatchsLikesPaymentRequired creates GetMatchsLikesPaymentRequired with default headers values
func NewGetMatchsLikesPaymentRequired() *GetMatchsLikesPaymentRequired {

	return &GetMatchsLikesPaymentRequired{}
}

// WithPayload adds the payload to the get matchs likes payment required response
func (o *GetMatchsLikesPaymentRequired) WithPayload(payload *models.Error) *GetMatchsLikesPaymentRequired {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get matchs likes payment required response
func (o *GetMatchsLikesPaymentRequired) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMatchsLikesPaymentRequired) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(402)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetMatchsLikesDefault generic error response

swagger:response getMatchsLikesDefault
*/
type GetMatchsLikesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMatchsLikesDefault creates GetMatchsLikesDefault with default headers values
func NewGetMatchsLikesDefault(code int) *GetMatchsLikesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMatchsLikesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get matchs likes default response
func (o *GetMatchsLikesDefault) WithStatusCode(code int) *GetMatchsLikesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get matchs likes default response
func (o *GetMatchsLikesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get matchs likes default response
func (o *GetMatchsLikesDefault) WithPayload(payload *models.Error) *GetMatchsLikesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get matchs likes default response
func (o *GetMatchsLikesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMatchsLikesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
