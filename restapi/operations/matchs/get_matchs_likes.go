// Code generated by go-swagger; DO NOT EDIT.

package matchs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/capella/topcrush/models"
)

// GetMatchsLikesHandlerFunc turns a function with the right signature into a get matchs likes handler
type GetMatchsLikesHandlerFunc func(GetMatchsLikesParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMatchsLikesHandlerFunc) Handle(params GetMatchsLikesParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetMatchsLikesHandler interface for that can handle valid get matchs likes params
type GetMatchsLikesHandler interface {
	Handle(GetMatchsLikesParams, *models.Principal) middleware.Responder
}

// NewGetMatchsLikes creates a new http.Handler for the get matchs likes operation
func NewGetMatchsLikes(ctx *middleware.Context, handler GetMatchsLikesHandler) *GetMatchsLikes {
	return &GetMatchsLikes{Context: ctx, Handler: handler}
}

/*GetMatchsLikes swagger:route GET /matchs/likes matchs getMatchsLikes

all persons who like me

*/
type GetMatchsLikes struct {
	Context *middleware.Context
	Handler GetMatchsLikesHandler
}

func (o *GetMatchsLikes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMatchsLikesParams()

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
