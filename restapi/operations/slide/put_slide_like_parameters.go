// Code generated by go-swagger; DO NOT EDIT.

package slide

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutSlideLikeParams creates a new PutSlideLikeParams object
// no default values defined in spec.
func NewPutSlideLikeParams() PutSlideLikeParams {

	return PutSlideLikeParams{}
}

// PutSlideLikeParams contains all the bound params for the put slide like operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutSlideLike
type PutSlideLikeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body strfmt.ObjectId
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutSlideLikeParams() beforehand.
func (o *PutSlideLikeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body strfmt.ObjectId
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate inline body
			o.Body = body
			if err := o.validateBodyBody(route.Formats); err != nil {
				res = append(res, err)
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// validateBodyBody validates an inline body parameter
func (o *PutSlideLikeParams) validateBodyBody(formats strfmt.Registry) error {

	if err := validate.FormatOf("body", "body", "ObjectId", o.Body.String(), formats); err != nil {
		return err
	}
	return nil
}