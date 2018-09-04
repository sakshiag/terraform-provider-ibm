// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostServersParamsBodyResourceGroup idreference
// swagger:model postServersParamsBodyResourceGroup
type PostServersParamsBodyResourceGroup struct {

	// The unique identifier for this resource
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this post servers params body resource group
func (m *PostServersParamsBodyResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostServersParamsBodyResourceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostServersParamsBodyResourceGroup) UnmarshalBinary(b []byte) error {
	var res PostServersParamsBodyResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
