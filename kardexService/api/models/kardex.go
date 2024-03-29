// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Kardex Kardex
// swagger:model Kardex
type Kardex struct {

	// id
	ID string `json:"id,omitempty"`

	// product Id
	ProductID string `json:"productId,omitempty"`
}

// Validate validates this kardex
func (m *Kardex) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Kardex) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Kardex) UnmarshalBinary(b []byte) error {
	var res Kardex
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
