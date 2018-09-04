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

// SecurityGroupRuleTemplateRemote SecurityGroupIdentityByName
//
// Uniquely identifies a security group using any one of ID, CRN, or name.
// swagger:model securityGroupRuleTemplateRemote
type SecurityGroupRuleTemplateRemote struct {

	// A single IPv4 or IPv6 address.
	Address string `json:"address,omitempty"`

	// A range of IPv4 or IPv6 addresses, in CIDR format.
	CidrBlock string `json:"cidr_block,omitempty"`

	// The security group's CRN
	Crn string `json:"crn,omitempty"`

	// The URL for the first page of resources
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The security group's unique identifier.
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this security group rule template remote
func (m *SecurityGroupRuleTemplateRemote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityGroupRuleTemplateRemote) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroupRuleTemplateRemote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroupRuleTemplateRemote) UnmarshalBinary(b []byte) error {
	var res SecurityGroupRuleTemplateRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
