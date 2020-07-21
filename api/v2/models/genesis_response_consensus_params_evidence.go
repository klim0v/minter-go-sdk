// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenesisResponseConsensusParamsEvidence genesis response consensus params evidence
//
// swagger:model GenesisResponseConsensusParamsEvidence
type GenesisResponseConsensusParamsEvidence struct {

	// max age duration
	MaxAgeDuration string `json:"max_age_duration,omitempty"`

	// max age num blocks
	MaxAgeNumBlocks string `json:"max_age_num_blocks,omitempty"`
}

// Validate validates this genesis response consensus params evidence
func (m *GenesisResponseConsensusParamsEvidence) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenesisResponseConsensusParamsEvidence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenesisResponseConsensusParamsEvidence) UnmarshalBinary(b []byte) error {
	var res GenesisResponseConsensusParamsEvidence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}