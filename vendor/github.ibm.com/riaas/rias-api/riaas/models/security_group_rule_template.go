// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityGroupRuleTemplate SecurityGroupRuleTemplate
//
// When 'protocol' is either of 'tcp' or 'udp', then the rule may also contain fields to specify a port range.
// swagger:model SecurityGroupRuleTemplate
type SecurityGroupRuleTemplate struct {

	// The ICMP code to to allow. Valid values from 0 to 255. If missing, allow all codes.
	Code *int64 `json:"code,omitempty"`

	// The direction of traffic to enforce (ingress, egress)
	Direction string `json:"direction,omitempty"`

	// The IP version to enforce (ipv4, ipv6). The format of 'remote.address' or 'remote.cidr_block' must match this field, if they are used. Also, if 'remote' references another security group (ie. using remote.id, remote.name, remote.crn) then this rule will only apply to IP addresses (network interfaces) in that group which match this ip_version.
	IPVersion string `json:"ip_version,omitempty"`

	// The highest port in the range of ports to be matched; if unspecified, `65535` is used.
	PortMax *int64 `json:"port_max,omitempty"`

	// The lowest port in the range of ports to be matched; if unspecified, `1` is used.
	PortMin *int64 `json:"port_min,omitempty"`

	// The protocol to enforce. Must be one of (icmp, tcp, udp, all). Defaults to 'all' if omitted.
	Protocol *string `json:"protocol,omitempty"`

	// remote
	Remote *SecurityGroupRuleTemplateRemote `json:"remote,omitempty"`

	// The ICMP type to allow. Valid values from 0 to 255. If missing, allow all types.
	Type *int64 `json:"type,omitempty"`
}

// Validate validates this security group rule template
func (m *SecurityGroupRuleTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRemote(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var securityGroupRuleTemplateTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ingress","egress"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityGroupRuleTemplateTypeDirectionPropEnum = append(securityGroupRuleTemplateTypeDirectionPropEnum, v)
	}
}

const (
	// SecurityGroupRuleTemplateDirectionIngress captures enum value "ingress"
	SecurityGroupRuleTemplateDirectionIngress string = "ingress"
	// SecurityGroupRuleTemplateDirectionEgress captures enum value "egress"
	SecurityGroupRuleTemplateDirectionEgress string = "egress"
)

// prop value enum
func (m *SecurityGroupRuleTemplate) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, securityGroupRuleTemplateTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SecurityGroupRuleTemplate) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var securityGroupRuleTemplateTypeIPVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityGroupRuleTemplateTypeIPVersionPropEnum = append(securityGroupRuleTemplateTypeIPVersionPropEnum, v)
	}
}

const (
	// SecurityGroupRuleTemplateIPVersionIPV4 captures enum value "ipv4"
	SecurityGroupRuleTemplateIPVersionIPV4 string = "ipv4"
	// SecurityGroupRuleTemplateIPVersionIPV6 captures enum value "ipv6"
	SecurityGroupRuleTemplateIPVersionIPV6 string = "ipv6"
)

// prop value enum
func (m *SecurityGroupRuleTemplate) validateIPVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, securityGroupRuleTemplateTypeIPVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SecurityGroupRuleTemplate) validateIPVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.IPVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateIPVersionEnum("ip_version", "body", m.IPVersion); err != nil {
		return err
	}

	return nil
}

var securityGroupRuleTemplateTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","icmp","tcp","udp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityGroupRuleTemplateTypeProtocolPropEnum = append(securityGroupRuleTemplateTypeProtocolPropEnum, v)
	}
}

const (
	// SecurityGroupRuleTemplateProtocolAll captures enum value "all"
	SecurityGroupRuleTemplateProtocolAll string = "all"
	// SecurityGroupRuleTemplateProtocolIcmp captures enum value "icmp"
	SecurityGroupRuleTemplateProtocolIcmp string = "icmp"
	// SecurityGroupRuleTemplateProtocolTCP captures enum value "tcp"
	SecurityGroupRuleTemplateProtocolTCP string = "tcp"
	// SecurityGroupRuleTemplateProtocolUDP captures enum value "udp"
	SecurityGroupRuleTemplateProtocolUDP string = "udp"
)

// prop value enum
func (m *SecurityGroupRuleTemplate) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, securityGroupRuleTemplateTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SecurityGroupRuleTemplate) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *SecurityGroupRuleTemplate) validateRemote(formats strfmt.Registry) error {

	if swag.IsZero(m.Remote) { // not required
		return nil
	}

	if m.Remote != nil {

		if err := m.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityGroupRuleTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityGroupRuleTemplate) UnmarshalBinary(b []byte) error {
	var res SecurityGroupRuleTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
