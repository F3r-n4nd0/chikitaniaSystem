// Code generated by go-swagger; DO NOT EDIT.

package sale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ReturnSaleCreatedCode is the HTTP code returned for type ReturnSaleCreated
const ReturnSaleCreatedCode int = 201

/*ReturnSaleCreated Created

swagger:response returnSaleCreated
*/
type ReturnSaleCreated struct {
}

// NewReturnSaleCreated creates ReturnSaleCreated with default headers values
func NewReturnSaleCreated() *ReturnSaleCreated {

	return &ReturnSaleCreated{}
}

// WriteResponse to the client
func (o *ReturnSaleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// ReturnSaleBadRequestCode is the HTTP code returned for type ReturnSaleBadRequest
const ReturnSaleBadRequestCode int = 400

/*ReturnSaleBadRequest Invalid Order

swagger:response returnSaleBadRequest
*/
type ReturnSaleBadRequest struct {
}

// NewReturnSaleBadRequest creates ReturnSaleBadRequest with default headers values
func NewReturnSaleBadRequest() *ReturnSaleBadRequest {

	return &ReturnSaleBadRequest{}
}

// WriteResponse to the client
func (o *ReturnSaleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}