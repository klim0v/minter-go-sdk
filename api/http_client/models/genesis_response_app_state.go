// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenesisResponseAppState genesis response app state
//
// swagger:model GenesisResponseAppState
type GenesisResponseAppState struct {

	// accounts
	Accounts []*AppStateAccount `json:"accounts"`

	// candidates
	Candidates []*AppStateCandidate `json:"candidates"`

	// coins
	Coins []*GenesisResponseAppStateCoin `json:"coins"`

	// frozen funds
	FrozenFunds []*AppStateFrozenFund `json:"frozen_funds"`

	// halt blocks
	HaltBlocks []*AppStateHaltBlock `json:"halt_blocks"`

	// max gas
	MaxGas uint64 `json:"max_gas,omitempty,string"`

	// note
	Note string `json:"note,omitempty"`

	// start height
	StartHeight uint64 `json:"start_height,omitempty,string"`

	// total slashed
	TotalSlashed string `json:"total_slashed,omitempty"`

	// used checks
	UsedChecks []string `json:"used_checks"`

	// validators
	Validators []*AppStateValidators `json:"validators"`

	// waitlist
	Waitlist []*AppStateWaitlist `json:"waitlist"`
}

// Validate validates this genesis response app state
func (m *GenesisResponseAppState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCandidates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrozenFunds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHaltBlocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWaitlist(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenesisResponseAppState) validateAccounts(formats strfmt.Registry) error {

	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Accounts); i++ {
		if swag.IsZero(m.Accounts[i]) { // not required
			continue
		}

		if m.Accounts[i] != nil {
			if err := m.Accounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GenesisResponseAppState) validateCandidates(formats strfmt.Registry) error {

	if swag.IsZero(m.Candidates) { // not required
		return nil
	}

	for i := 0; i < len(m.Candidates); i++ {
		if swag.IsZero(m.Candidates[i]) { // not required
			continue
		}

		if m.Candidates[i] != nil {
			if err := m.Candidates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("candidates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GenesisResponseAppState) validateCoins(formats strfmt.Registry) error {

	if swag.IsZero(m.Coins) { // not required
		return nil
	}

	for i := 0; i < len(m.Coins); i++ {
		if swag.IsZero(m.Coins[i]) { // not required
			continue
		}

		if m.Coins[i] != nil {
			if err := m.Coins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GenesisResponseAppState) validateFrozenFunds(formats strfmt.Registry) error {

	if swag.IsZero(m.FrozenFunds) { // not required
		return nil
	}

	for i := 0; i < len(m.FrozenFunds); i++ {
		if swag.IsZero(m.FrozenFunds[i]) { // not required
			continue
		}

		if m.FrozenFunds[i] != nil {
			if err := m.FrozenFunds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("frozen_funds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GenesisResponseAppState) validateHaltBlocks(formats strfmt.Registry) error {

	if swag.IsZero(m.HaltBlocks) { // not required
		return nil
	}

	for i := 0; i < len(m.HaltBlocks); i++ {
		if swag.IsZero(m.HaltBlocks[i]) { // not required
			continue
		}

		if m.HaltBlocks[i] != nil {
			if err := m.HaltBlocks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("halt_blocks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GenesisResponseAppState) validateValidators(formats strfmt.Registry) error {

	if swag.IsZero(m.Validators) { // not required
		return nil
	}

	for i := 0; i < len(m.Validators); i++ {
		if swag.IsZero(m.Validators[i]) { // not required
			continue
		}

		if m.Validators[i] != nil {
			if err := m.Validators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GenesisResponseAppState) validateWaitlist(formats strfmt.Registry) error {

	if swag.IsZero(m.Waitlist) { // not required
		return nil
	}

	for i := 0; i < len(m.Waitlist); i++ {
		if swag.IsZero(m.Waitlist[i]) { // not required
			continue
		}

		if m.Waitlist[i] != nil {
			if err := m.Waitlist[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("waitlist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenesisResponseAppState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenesisResponseAppState) UnmarshalBinary(b []byte) error {
	var res GenesisResponseAppState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}