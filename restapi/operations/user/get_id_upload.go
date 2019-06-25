// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetIDUploadHandlerFunc turns a function with the right signature into a get ID upload handler
type GetIDUploadHandlerFunc func(GetIDUploadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetIDUploadHandlerFunc) Handle(params GetIDUploadParams) middleware.Responder {
	return fn(params)
}

// GetIDUploadHandler interface for that can handle valid get ID upload params
type GetIDUploadHandler interface {
	Handle(GetIDUploadParams) middleware.Responder
}

// NewGetIDUpload creates a new http.Handler for the get ID upload operation
func NewGetIDUpload(ctx *middleware.Context, handler GetIDUploadHandler) *GetIDUpload {
	return &GetIDUpload{Context: ctx, Handler: handler}
}

/*GetIDUpload swagger:route GET /{id}/upload user getIdUpload

GetIDUpload get ID upload API

*/
type GetIDUpload struct {
	Context *middleware.Context
	Handler GetIDUploadHandler
}

func (o *GetIDUpload) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetIDUploadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
