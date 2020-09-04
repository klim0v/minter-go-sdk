// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPbErrorBodyError api pb error body error
//
// swagger:model api_pbErrorBodyError
type APIPbErrorBodyError struct {

	// code
	Code string `json:"code,omitempty"`

	// data
	Data map[string]string `json:"data,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this api pb error body error
func (m *APIPbErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPbErrorBodyError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPbErrorBodyError) UnmarshalBinary(b []byte) error {
	var res APIPbErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
