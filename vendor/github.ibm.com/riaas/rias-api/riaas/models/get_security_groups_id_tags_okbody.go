// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetSecurityGroupsIDTagsOKBody get security groups Id tags o k body
// swagger:model getSecurityGroupsIdTagsOKBody
type GetSecurityGroupsIDTagsOKBody struct {

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`
}

// Validate validates this get security groups Id tags o k body
func (m *GetSecurityGroupsIDTagsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSecurityGroupsIDTagsOKBody) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSecurityGroupsIDTagsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSecurityGroupsIDTagsOKBody) UnmarshalBinary(b []byte) error {
	var res GetSecurityGroupsIDTagsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
