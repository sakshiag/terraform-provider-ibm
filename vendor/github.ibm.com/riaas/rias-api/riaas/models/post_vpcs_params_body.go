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

// PostVpcsParamsBody VPCTemplate
// swagger:model postVpcsParamsBody
type PostVpcsParamsBody struct {

	// Indicates whether this VPC is connected to Classic Infrastructure. If true, this VPC's
	// resources have private network connectivity to the account's Classic Infrastructure
	// resources. Only one VPC on an account may be connected in this way. This value is set at
	// creation and subsequently immutable.
	//
	ClassicAccess bool `json:"classic_access,omitempty"`

	// default network acl
	DefaultNetworkACL *PostVpcsParamsBodyDefaultNetworkACL `json:"default_network_acl,omitempty"`

	// Indicates whether this is the default VPC for the account
	IsDefault bool `json:"is_default,omitempty"`

	// The user-defined name for this VPC
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// resource group
	ResourceGroup *PostVpcsParamsBodyResourceGroup `json:"resource_group,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`
}

// Validate validates this post vpcs params body
func (m *PostVpcsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultNetworkACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostVpcsParamsBody) validateDefaultNetworkACL(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultNetworkACL) { // not required
		return nil
	}

	if m.DefaultNetworkACL != nil {
		if err := m.DefaultNetworkACL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_network_acl")
			}
			return err
		}
	}

	return nil
}

func (m *PostVpcsParamsBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PostVpcsParamsBody) validateResourceGroup(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PostVpcsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostVpcsParamsBody) UnmarshalBinary(b []byte) error {
	var res PostVpcsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
