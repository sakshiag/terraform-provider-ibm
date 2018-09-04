// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkInterfaceReference NetworkInterfaceReference
// swagger:model network_interface_reference
type NetworkInterfaceReference struct {
	ResourceReference

	NetworkInterfaceReferenceAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NetworkInterfaceReference) UnmarshalJSON(raw []byte) error {

	var aO0 ResourceReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceReference = aO0

	var aO1 NetworkInterfaceReferenceAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NetworkInterfaceReferenceAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NetworkInterfaceReference) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.ResourceReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.NetworkInterfaceReferenceAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this network interface reference
func (m *NetworkInterfaceReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.ResourceReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.NetworkInterfaceReferenceAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInterfaceReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterfaceReference) UnmarshalBinary(b []byte) error {
	var res NetworkInterfaceReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
