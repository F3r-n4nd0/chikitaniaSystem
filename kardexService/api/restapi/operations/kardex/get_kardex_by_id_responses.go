// Code generated by go-swagger; DO NOT EDIT.

package kardex

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "kardexService/api/models"
)

// GetKardexByIDOKCode is the HTTP code returned for type GetKardexByIDOK
const GetKardexByIDOKCode int = 200

/*GetKardexByIDOK Successful operation

swagger:response getKardexByIdOK
*/
type GetKardexByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Kardex `json:"body,omitempty"`
}

// NewGetKardexByIDOK creates GetKardexByIDOK with default headers values
func NewGetKardexByIDOK() *GetKardexByIDOK {

	return &GetKardexByIDOK{}
}

// WithPayload adds the payload to the get kardex by Id o k response
func (o *GetKardexByIDOK) WithPayload(payload *models.Kardex) *GetKardexByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get kardex by Id o k response
func (o *GetKardexByIDOK) SetPayload(payload *models.Kardex) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetKardexByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetKardexByIDBadRequestCode is the HTTP code returned for type GetKardexByIDBadRequest
const GetKardexByIDBadRequestCode int = 400

/*GetKardexByIDBadRequest Invalid ID supplied

swagger:response getKardexByIdBadRequest
*/
type GetKardexByIDBadRequest struct {
}

// NewGetKardexByIDBadRequest creates GetKardexByIDBadRequest with default headers values
func NewGetKardexByIDBadRequest() *GetKardexByIDBadRequest {

	return &GetKardexByIDBadRequest{}
}

// WriteResponse to the client
func (o *GetKardexByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetKardexByIDNotFoundCode is the HTTP code returned for type GetKardexByIDNotFound
const GetKardexByIDNotFoundCode int = 404

/*GetKardexByIDNotFound Kardex not found

swagger:response getKardexByIdNotFound
*/
type GetKardexByIDNotFound struct {
}

// NewGetKardexByIDNotFound creates GetKardexByIDNotFound with default headers values
func NewGetKardexByIDNotFound() *GetKardexByIDNotFound {

	return &GetKardexByIDNotFound{}
}

// WriteResponse to the client
func (o *GetKardexByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}