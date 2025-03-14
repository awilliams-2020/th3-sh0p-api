// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostImagePackHandlerFunc turns a function with the right signature into a post image pack handler
type PostImagePackHandlerFunc func(PostImagePackParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostImagePackHandlerFunc) Handle(params PostImagePackParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostImagePackHandler interface for that can handle valid post image pack params
type PostImagePackHandler interface {
	Handle(PostImagePackParams, interface{}) middleware.Responder
}

// NewPostImagePack creates a new http.Handler for the post image pack operation
func NewPostImagePack(ctx *middleware.Context, handler PostImagePackHandler) *PostImagePack {
	return &PostImagePack{Context: ctx, Handler: handler}
}

/* PostImagePack swagger:route POST /image-pack postImagePack

PostImagePack post image pack API

*/
type PostImagePack struct {
	Context *middleware.Context
	Handler PostImagePackHandler
}

func (o *PostImagePack) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostImagePackParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostImagePackBody post image pack body
//
// swagger:model PostImagePackBody
type PostImagePackBody struct {

	// image pack
	// Required: true
	// Enum: [pack_1 pack_2 pack_3]
	ImagePack *string `json:"imagePack"`
}

// Validate validates this post image pack body
func (o *PostImagePackBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateImagePack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postImagePackBodyTypeImagePackPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pack_1","pack_2","pack_3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postImagePackBodyTypeImagePackPropEnum = append(postImagePackBodyTypeImagePackPropEnum, v)
	}
}

const (

	// PostImagePackBodyImagePackPack1 captures enum value "pack_1"
	PostImagePackBodyImagePackPack1 string = "pack_1"

	// PostImagePackBodyImagePackPack2 captures enum value "pack_2"
	PostImagePackBodyImagePackPack2 string = "pack_2"

	// PostImagePackBodyImagePackPack3 captures enum value "pack_3"
	PostImagePackBodyImagePackPack3 string = "pack_3"
)

// prop value enum
func (o *PostImagePackBody) validateImagePackEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postImagePackBodyTypeImagePackPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostImagePackBody) validateImagePack(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"imagePack", "body", o.ImagePack); err != nil {
		return err
	}

	// value enum
	if err := o.validateImagePackEnum("body"+"."+"imagePack", "body", *o.ImagePack); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post image pack body based on context it is used
func (o *PostImagePackBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostImagePackBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostImagePackBody) UnmarshalBinary(b []byte) error {
	var res PostImagePackBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostImagePackOKBody post image pack o k body
//
// swagger:model PostImagePackOKBody
type PostImagePackOKBody struct {

	// payment intent
	PaymentIntent string `json:"paymentIntent,omitempty"`
}

// Validate validates this post image pack o k body
func (o *PostImagePackOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post image pack o k body based on context it is used
func (o *PostImagePackOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostImagePackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostImagePackOKBody) UnmarshalBinary(b []byte) error {
	var res PostImagePackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
