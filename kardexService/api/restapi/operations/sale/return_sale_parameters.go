// Code generated by go-swagger; DO NOT EDIT.

package sale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "kardexService/api/models"
)

// NewReturnSaleParams creates a new ReturnSaleParams object
// no default values defined in spec.
func NewReturnSaleParams() ReturnSaleParams {

	return ReturnSaleParams{}
}

// ReturnSaleParams contains all the bound params for the return sale operation
// typically these are obtained from a http.Request
//
// swagger:parameters returnSale
type ReturnSaleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*placed a return on sale
	  Required: true
	  In: body
	*/
	Body *models.ReturnSale
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewReturnSaleParams() beforehand.
func (o *ReturnSaleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ReturnSale
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
