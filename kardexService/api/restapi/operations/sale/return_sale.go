// Code generated by go-swagger; DO NOT EDIT.

package sale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReturnSaleHandlerFunc turns a function with the right signature into a return sale handler
type ReturnSaleHandlerFunc func(ReturnSaleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReturnSaleHandlerFunc) Handle(params ReturnSaleParams) middleware.Responder {
	return fn(params)
}

// ReturnSaleHandler interface for that can handle valid return sale params
type ReturnSaleHandler interface {
	Handle(ReturnSaleParams) middleware.Responder
}

// NewReturnSale creates a new http.Handler for the return sale operation
func NewReturnSale(ctx *middleware.Context, handler ReturnSaleHandler) *ReturnSale {
	return &ReturnSale{Context: ctx, Handler: handler}
}

/*ReturnSale swagger:route POST /inventory/sale/return sale returnSale

Return on sale

*/
type ReturnSale struct {
	Context *middleware.Context
	Handler ReturnSaleHandler
}

func (o *ReturnSale) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReturnSaleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
