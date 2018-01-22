// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1ServerPort v1 server port
// swagger:model V1ServerPort
type V1ServerPort struct {

	// For AUTO_IP, AGENT_IP and CONTAINER_IP this can be used to expose all defined ports.
	// This should only be used if a name is not defined for the port and there is a single port defined for the service.
	All bool `json:"all,omitempty"`

	// The name of the port. This is used for AUTO_IP, AGENT_IP and CONTAINER_IP.
	Name string `json:"name,omitempty"`

	// Set the VIP definition directly (e.g. "/myvip:1234").
	Vip string `json:"vip,omitempty"`
}

// Validate validates this v1 server port
func (m *V1ServerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1ServerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ServerPort) UnmarshalBinary(b []byte) error {
	var res V1ServerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
