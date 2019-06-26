// Code generated by go-swagger; DO NOT EDIT.

package slide

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/capella/topcrush/models"
)

// PutSlideLikeHandlerFunc turns a function with the right signature into a put slide like handler
type PutSlideLikeHandlerFunc func(PutSlideLikeParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PutSlideLikeHandlerFunc) Handle(params PutSlideLikeParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PutSlideLikeHandler interface for that can handle valid put slide like params
type PutSlideLikeHandler interface {
	Handle(PutSlideLikeParams, *models.Principal) middleware.Responder
}

// NewPutSlideLike creates a new http.Handler for the put slide like operation
func NewPutSlideLike(ctx *middleware.Context, handler PutSlideLikeHandler) *PutSlideLike {
	return &PutSlideLike{Context: ctx, Handler: handler}
}

/*PutSlideLike swagger:route PUT /slide/like slide putSlideLike

make a like

*/
type PutSlideLike struct {
	Context *middleware.Context
	Handler PutSlideLikeHandler
}

func (o *PutSlideLike) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutSlideLikeParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
