// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Player player
//
// swagger:model Player
type Player string

const (

	// PlayerP1 captures enum value "P1"
	PlayerP1 Player = "P1"

	// PlayerP2 captures enum value "P2"
	PlayerP2 Player = "P2"
)

// for schema
var playerEnum []interface{}

func init() {
	var res []Player
	if err := json.Unmarshal([]byte(`["P1","P2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		playerEnum = append(playerEnum, v)
	}
}

func (m Player) validatePlayerEnum(path, location string, value Player) error {
	if err := validate.EnumCase(path, location, value, playerEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this player
func (m Player) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePlayerEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
