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

// ServerVolumeAttachment VolumeAttachment
// swagger:model server_volume_attachment
type ServerVolumeAttachment struct {

	// The date and time that the volume was attached
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The CRN for this volume
	Crn string `json:"crn,omitempty"`

	// The URL for this volume interface
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href string `json:"href,omitempty"`

	// The unique identifier for this volume interface
	ID strfmt.UUID `json:"id,omitempty"`

	// The user-defined name for this volume interface
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// resource group
	ResourceGroup *ResourceReference `json:"resource_group,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`

	// volume
	Volume *ResourceReference `json:"volume,omitempty"`
}

// Validate validates this server volume attachment
func (m *ServerVolumeAttachment) Validate(formats strfmt.Registry) error {
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

	if err := m.validateVolume(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerVolumeAttachment) validateHref(formats strfmt.Registry) error {

	if swag.IsZero(m.Href) { // not required
		return nil
	}

	if err := validate.Pattern("href", "body", string(m.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

func (m *ServerVolumeAttachment) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ServerVolumeAttachment) validateResourceGroup(formats strfmt.Registry) error {

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

var serverVolumeAttachmentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["attaching","attached","detaching"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverVolumeAttachmentTypeStatusPropEnum = append(serverVolumeAttachmentTypeStatusPropEnum, v)
	}
}

const (
	// ServerVolumeAttachmentStatusAttaching captures enum value "attaching"
	ServerVolumeAttachmentStatusAttaching string = "attaching"
	// ServerVolumeAttachmentStatusAttached captures enum value "attached"
	ServerVolumeAttachmentStatusAttached string = "attached"
	// ServerVolumeAttachmentStatusDetaching captures enum value "detaching"
	ServerVolumeAttachmentStatusDetaching string = "detaching"
)

// prop value enum
func (m *ServerVolumeAttachment) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverVolumeAttachmentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServerVolumeAttachment) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ServerVolumeAttachment) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

func (m *ServerVolumeAttachment) validateVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {

		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerVolumeAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerVolumeAttachment) UnmarshalBinary(b []byte) error {
	var res ServerVolumeAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
