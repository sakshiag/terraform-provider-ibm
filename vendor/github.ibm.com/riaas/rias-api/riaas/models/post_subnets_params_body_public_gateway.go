// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostSubnetsParamsBodyPublicGateway PublicGatewayIdentity
//
// The public gateway to handle internet bound traffic for this subnet.
// swagger:model postSubnetsParamsBodyPublicGateway
type PostSubnetsParamsBodyPublicGateway struct {

	// The unique identifier for this public gateway
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this post subnets params body public gateway
func (m *PostSubnetsParamsBodyPublicGateway) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostSubnetsParamsBodyPublicGateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostSubnetsParamsBodyPublicGateway) UnmarshalBinary(b []byte) error {
	var res PostSubnetsParamsBodyPublicGateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
