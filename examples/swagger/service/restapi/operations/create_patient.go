// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreatePatientHandlerFunc turns a function with the right signature into a create patient handler
type CreatePatientHandlerFunc func(CreatePatientParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreatePatientHandlerFunc) Handle(params CreatePatientParams) middleware.Responder {
	return fn(params)
}

// CreatePatientHandler interface for that can handle valid create patient params
type CreatePatientHandler interface {
	Handle(CreatePatientParams) middleware.Responder
}

// NewCreatePatient creates a new http.Handler for the create patient operation
func NewCreatePatient(ctx *middleware.Context, handler CreatePatientHandler) *CreatePatient {
	return &CreatePatient{Context: ctx, Handler: handler}
}

/* CreatePatient swagger:route POST /patients createPatient

CreatePatient create patient API

*/
type CreatePatient struct {
	Context *middleware.Context
	Handler CreatePatientHandler
}

func (o *CreatePatient) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreatePatientParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}