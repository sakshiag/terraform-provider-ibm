// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetServersIDTagsOKBody get servers Id tags o k body
// swagger:model getServersIdTagsOKBody
type GetServersIDTagsOKBody struct {

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`
}

// Validate validates this get servers Id tags o k body
func (m *GetServersIDTagsOKBody) Validate(formats strfmt.Registry) error {
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

func (m *GetServersIDTagsOKBody) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetServersIDTagsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetServersIDTagsOKBody) UnmarshalBinary(b []byte) error {
	var res GetServersIDTagsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
