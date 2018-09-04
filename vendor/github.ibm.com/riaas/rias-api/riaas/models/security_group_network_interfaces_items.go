// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityGroupNetworkInterfacesItems NetworkInterfaceReference
// swagger:model securityGroupNetworkInterfacesItems
type SecurityGroupNetworkInterfacesItems struct {

	// The CRN for this network interface
	Crn string `json:"crn,omitempty"`

	// The URL for this network interface
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The unique identifier for this network interface
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this network interface
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// The primary IPv4 address
	PrimaryIPV4Address string `json:"primary_ipv4_address,omitempty"`
}

// Validate validates this security group network interfaces items
func (m *SecurityGroupNetworkInterfacesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupNetworkInterfacesItems) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *SecurityGroupNetworkInterfacesItems) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroupNetworkInterfacesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroupNetworkInterfacesItems) UnmarshalBinary(b []byte) error {
	var res SecurityGroupNetworkInterfacesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
