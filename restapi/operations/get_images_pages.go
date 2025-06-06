// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetImagesPagesHandlerFunc turns a function with the right signature into a get images pages handler
type GetImagesPagesHandlerFunc func(GetImagesPagesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetImagesPagesHandlerFunc) Handle(params GetImagesPagesParams) middleware.Responder {
	return fn(params)
}

// GetImagesPagesHandler interface for that can handle valid get images pages params
type GetImagesPagesHandler interface {
	Handle(GetImagesPagesParams) middleware.Responder
}

// NewGetImagesPages creates a new http.Handler for the get images pages operation
func NewGetImagesPages(ctx *middleware.Context, handler GetImagesPagesHandler) *GetImagesPages {
	return &GetImagesPages{Context: ctx, Handler: handler}
}

/* GetImagesPages swagger:route GET /images/pages getImagesPages

GetImagesPages get images pages API

*/
type GetImagesPages struct {
	Context *middleware.Context
	Handler GetImagesPagesHandler
}

func (o *GetImagesPages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetImagesPagesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
