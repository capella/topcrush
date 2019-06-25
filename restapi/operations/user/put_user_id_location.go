// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutUserIDLocationHandlerFunc turns a function with the right signature into a put user ID location handler
type PutUserIDLocationHandlerFunc func(PutUserIDLocationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutUserIDLocationHandlerFunc) Handle(params PutUserIDLocationParams) middleware.Responder {
	return fn(params)
}

// PutUserIDLocationHandler interface for that can handle valid put user ID location params
type PutUserIDLocationHandler interface {
	Handle(PutUserIDLocationParams) middleware.Responder
}

// NewPutUserIDLocation creates a new http.Handler for the put user ID location operation
func NewPutUserIDLocation(ctx *middleware.Context, handler PutUserIDLocationHandler) *PutUserIDLocation {
	return &PutUserIDLocation{Context: ctx, Handler: handler}
}

/*PutUserIDLocation swagger:route PUT /user/{id}/location user putUserIdLocation

update user position

*/
type PutUserIDLocation struct {
	Context *middleware.Context
	Handler PutUserIDLocationHandler
}

func (o *PutUserIDLocation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutUserIDLocationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}