// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountMultisigData account multisig data
//
// swagger:model AccountMultisigData
type AccountMultisigData struct {

	// addresses
	Addresses []string `json:"addresses"`

	// threshold
	Threshold uint64 `json:"threshold,omitempty,string"`

	// weights
	Weights []uint64 `json:"weights"`
}

// Validate validates this account multisig data
func (m *AccountMultisigData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountMultisigData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountMultisigData) UnmarshalBinary(b []byte) error {
	var res AccountMultisigData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}