// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostVolumesParamsBodyResourceGroup idreference
// swagger:model postVolumesParamsBodyResourceGroup
type PostVolumesParamsBodyResourceGroup struct {

	// The unique identifier for this resource
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this post volumes params body resource group
func (m *PostVolumesParamsBodyResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostVolumesParamsBodyResourceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostVolumesParamsBodyResourceGroup) UnmarshalBinary(b []byte) error {
	var res PostVolumesParamsBodyResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
