// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostInstancesParamsBodyBootVolumeAttachmentVolumeResourceGroup idreference
// swagger:model postInstancesParamsBodyBootVolumeAttachmentVolumeResourceGroup
type PostInstancesParamsBodyBootVolumeAttachmentVolumeResourceGroup struct {

	// The unique identifier for this resource
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this post instances params body boot volume attachment volume resource group
func (m *PostInstancesParamsBodyBootVolumeAttachmentVolumeResourceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostInstancesParamsBodyBootVolumeAttachmentVolumeResourceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostInstancesParamsBodyBootVolumeAttachmentVolumeResourceGroup) UnmarshalBinary(b []byte) error {
	var res PostInstancesParamsBodyBootVolumeAttachmentVolumeResourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
