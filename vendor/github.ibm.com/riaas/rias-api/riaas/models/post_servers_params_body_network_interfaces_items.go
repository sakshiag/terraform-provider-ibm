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

// PostServersParamsBodyNetworkInterfacesItems PrimaryNetworkInterfaceTemplate
//
// network interface
// swagger:model postServersParamsBodyNetworkInterfacesItems
type PostServersParamsBodyNetworkInterfacesItems struct {

	// The user-defined name for this network interface
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// The network interface port speed in Mbps
	PortSpeed int64 `json:"port_speed,omitempty"`

	// The primary IPv4 address
	PrimaryIPV4Address string `json:"primary_ipv4_address,omitempty"`

	// The primary IPv6 address in any valid notation as specified by RFC 4291
	PrimaryIPV6Address string `json:"primary_ipv6_address,omitempty"`

	// resource group
	ResourceGroup *PostServersParamsBodyNetworkInterfacesItemsResourceGroup `json:"resource_group,omitempty"`

	// Collection seconary IP addresses
	SecondaryAddresses []string `json:"secondary_addresses,omitempty"`

	// security groups
	SecurityGroups PostServersParamsBodyNetworkInterfacesItemsSecurityGroups `json:"security_groups,omitempty"`

	// subnet
	Subnet *PostServersParamsBodyNetworkInterfacesItemsSubnet `json:"subnet,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`
}

// Validate validates this post servers params body network interfaces items
func (m *PostServersParamsBodyNetworkInterfacesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSecondaryAddresses(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostServersParamsBodyNetworkInterfacesItems) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PostServersParamsBodyNetworkInterfacesItems) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {

		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

func (m *PostServersParamsBodyNetworkInterfacesItems) validateSecondaryAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryAddresses) { // not required
		return nil
	}

	return nil
}

func (m *PostServersParamsBodyNetworkInterfacesItems) validateSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {

		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

func (m *PostServersParamsBodyNetworkInterfacesItems) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostServersParamsBodyNetworkInterfacesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostServersParamsBodyNetworkInterfacesItems) UnmarshalBinary(b []byte) error {
	var res PostServersParamsBodyNetworkInterfacesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
