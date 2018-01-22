// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1Haproxy v1 haproxy
// swagger:model V1Haproxy
type V1Haproxy struct {

	// Array of backends.
	Backends []*V1Backend `json:"backends"`

	// Array of frontends.
	Frontends []*V1Frontend `json:"frontends"`

	// stats
	Stats *V1Stats `json:"stats,omitempty"`
}

func (m *V1Haproxy) UnmarshalJSON(b []byte) error {
	type V1HaproxyAlias V1Haproxy
	var t V1HaproxyAlias
	if err := json.Unmarshal([]byte("{\"backends\":[],\"frontends\":[],\"stats\":{}}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = V1Haproxy(t)
	return nil
}

// Validate validates this v1 haproxy
func (m *V1Haproxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackends(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFrontends(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Haproxy) validateBackends(formats strfmt.Registry) error {

	if swag.IsZero(m.Backends) { // not required
		return nil
	}

	for i := 0; i < len(m.Backends); i++ {

		if swag.IsZero(m.Backends[i]) { // not required
			continue
		}

		if m.Backends[i] != nil {

			if err := m.Backends[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backends" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Haproxy) validateFrontends(formats strfmt.Registry) error {

	if swag.IsZero(m.Frontends) { // not required
		return nil
	}

	for i := 0; i < len(m.Frontends); i++ {

		if swag.IsZero(m.Frontends[i]) { // not required
			continue
		}

		if m.Frontends[i] != nil {

			if err := m.Frontends[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("frontends" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Haproxy) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {

		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Haproxy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Haproxy) UnmarshalBinary(b []byte) error {
	var res V1Haproxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
