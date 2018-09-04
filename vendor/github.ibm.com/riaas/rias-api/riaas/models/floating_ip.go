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

// FloatingIP FloatingIP
// swagger:model floating_ip
type FloatingIP struct {

	// The globally unique IP address
	Address string `json:"address,omitempty"`

	// The date and time that the floating IP was created
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The CRN for this floating IP
	Crn string `json:"crn,omitempty"`

	// The URL for this floating IPs
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The unique identifier for this floating ip
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this floating IP
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// resource group
	ResourceGroup *ResourceReference `json:"resource_group,omitempty"`

	// The status of the floating IP
	Status string `json:"status,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`

	// target
	Target *FloatingIPTarget `json:"target,omitempty"`

	// zone
	Zone *FloatingIPZone `json:"zone,omitempty"`
}

// Validate validates this floating ip
func (m *FloatingIP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FloatingIP) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *FloatingIP) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *FloatingIP) validateResourceGroup(formats strfmt.Registry) error {

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

var floatingIpTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","available"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		floatingIpTypeStatusPropEnum = append(floatingIpTypeStatusPropEnum, v)
	}
}

const (
	// FloatingIPStatusPending captures enum value "pending"
	FloatingIPStatusPending string = "pending"
	// FloatingIPStatusAvailable captures enum value "available"
	FloatingIPStatusAvailable string = "available"
)

// prop value enum
func (m *FloatingIP) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, floatingIpTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FloatingIP) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *FloatingIP) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

func (m *FloatingIP) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {

		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *FloatingIP) validateZone(formats strfmt.Registry) error {

	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if m.Zone != nil {

		if err := m.Zone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FloatingIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FloatingIP) UnmarshalBinary(b []byte) error {
	var res FloatingIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
