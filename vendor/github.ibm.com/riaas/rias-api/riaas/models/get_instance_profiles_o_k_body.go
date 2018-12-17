// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetInstanceProfilesOKBody instanceProfileCollection
// swagger:model getInstanceProfilesOKBody
type GetInstanceProfilesOKBody struct {
	Pagination

	// Collection of instance profiels
	Profiles []*InstanceProfile `json:"profiles"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetInstanceProfilesOKBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Pagination
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Pagination = aO0

	// now for regular properties
	var propsGetInstanceProfilesOKBody struct {
		Profiles []*InstanceProfile `json:"profiles,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsGetInstanceProfilesOKBody); err != nil {
		return err
	}
	m.Profiles = propsGetInstanceProfilesOKBody.Profiles

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetInstanceProfilesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Pagination)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsGetInstanceProfilesOKBody struct {
		Profiles []*InstanceProfile `json:"profiles,omitempty"`
	}
	propsGetInstanceProfilesOKBody.Profiles = m.Profiles

	jsonDataPropsGetInstanceProfilesOKBody, errGetInstanceProfilesOKBody := swag.WriteJSON(propsGetInstanceProfilesOKBody)
	if errGetInstanceProfilesOKBody != nil {
		return nil, errGetInstanceProfilesOKBody
	}
	_parts = append(_parts, jsonDataPropsGetInstanceProfilesOKBody)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get instance profiles o k body
func (m *GetInstanceProfilesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Pagination
	if err := m.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetInstanceProfilesOKBody) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Profiles) { // not required
		return nil
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetInstanceProfilesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetInstanceProfilesOKBody) UnmarshalBinary(b []byte) error {
	var res GetInstanceProfilesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
