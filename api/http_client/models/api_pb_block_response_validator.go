// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPbBlockResponseValidator api pb block response validator
//
// swagger:model api_pbBlockResponseValidator
type APIPbBlockResponseValidator struct {

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// signed
	Signed bool `json:"signed,omitempty"`
}

// Validate validates this api pb block response validator
func (m *APIPbBlockResponseValidator) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPbBlockResponseValidator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPbBlockResponseValidator) UnmarshalBinary(b []byte) error {
	var res APIPbBlockResponseValidator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}