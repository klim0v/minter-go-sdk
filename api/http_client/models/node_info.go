// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeInfo node info
//
// swagger:model NodeInfo
type NodeInfo struct {

	// channels
	Channels string `json:"channels,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// listen addr
	ListenAddr string `json:"listen_addr,omitempty"`

	// moniker
	Moniker string `json:"moniker,omitempty"`

	// network
	Network string `json:"network,omitempty"`

	// other
	Other *NodeInfoOther `json:"other,omitempty"`

	// protocol version
	ProtocolVersion *NodeInfoProtocolVersion `json:"protocol_version,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this node info
func (m *NodeInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOther(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeInfo) validateOther(formats strfmt.Registry) error {

	if swag.IsZero(m.Other) { // not required
		return nil
	}

	if m.Other != nil {
		if err := m.Other.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("other")
			}
			return err
		}
	}

	return nil
}

func (m *NodeInfo) validateProtocolVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.ProtocolVersion) { // not required
		return nil
	}

	if m.ProtocolVersion != nil {
		if err := m.ProtocolVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol_version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeInfo) UnmarshalBinary(b []byte) error {
	var res NodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}