// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"th3-sh0p-api/models"
)

// GetGoogleProfileHandlerFunc turns a function with the right signature into a get google profile handler
type GetGoogleProfileHandlerFunc func(GetGoogleProfileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGoogleProfileHandlerFunc) Handle(params GetGoogleProfileParams) middleware.Responder {
	return fn(params)
}

// GetGoogleProfileHandler interface for that can handle valid get google profile params
type GetGoogleProfileHandler interface {
	Handle(GetGoogleProfileParams) middleware.Responder
}

// NewGetGoogleProfile creates a new http.Handler for the get google profile operation
func NewGetGoogleProfile(ctx *middleware.Context, handler GetGoogleProfileHandler) *GetGoogleProfile {
	return &GetGoogleProfile{Context: ctx, Handler: handler}
}

/* GetGoogleProfile swagger:route GET /google-profile getGoogleProfile

GetGoogleProfile get google profile API

*/
type GetGoogleProfile struct {
	Context *middleware.Context
	Handler GetGoogleProfileHandler
}

func (o *GetGoogleProfile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetGoogleProfileParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetGoogleProfileOKBody get google profile o k body
//
// swagger:model GetGoogleProfileOKBody
type GetGoogleProfileOKBody struct {

	// image credit
	ImageCredit int64 `json:"imageCredit,omitempty"`

	// profile
	Profile *models.Profile `json:"profile,omitempty"`
}

// Validate validates this get google profile o k body
func (o *GetGoogleProfileOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGoogleProfileOKBody) validateProfile(formats strfmt.Registry) error {
	if swag.IsZero(o.Profile) { // not required
		return nil
	}

	if o.Profile != nil {
		if err := o.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getGoogleProfileOK" + "." + "profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getGoogleProfileOK" + "." + "profile")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get google profile o k body based on the context it is used
func (o *GetGoogleProfileOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGoogleProfileOKBody) contextValidateProfile(ctx context.Context, formats strfmt.Registry) error {

	if o.Profile != nil {
		if err := o.Profile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getGoogleProfileOK" + "." + "profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getGoogleProfileOK" + "." + "profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGoogleProfileOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGoogleProfileOKBody) UnmarshalBinary(b []byte) error {
	var res GetGoogleProfileOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
