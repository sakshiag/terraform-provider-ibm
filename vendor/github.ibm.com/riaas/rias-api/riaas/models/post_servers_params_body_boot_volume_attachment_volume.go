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

// PostServersParamsBodyBootVolumeAttachmentVolume BootVolumeTemplateServerContext
//
// The identity of the volume to attach to the server or a template for a new volume
// swagger:model postServersParamsBodyBootVolumeAttachmentVolume
type PostServersParamsBodyBootVolumeAttachmentVolume struct {

	// If set to true, this volume will be automatically deleted if the only server it is attached to is deleted
	AutoDelete *bool `json:"auto_delete,omitempty"`

	// The billing term for the volume
	BillingTerm *string `json:"billing_term,omitempty"`

	// The capacity of the volume in gigabytes
	// Maximum: 64000
	// Minimum: 10
	Capacity int64 `json:"capacity,omitempty"`

	// The CRN for this volume
	Crn string `json:"crn,omitempty"`

	// encryption key
	EncryptionKey *PostServersParamsBodyBootVolumeAttachmentVolumeEncryptionKey `json:"encryption_key,omitempty"`

	// The unique identifier for this volume
	ID strfmt.UUID `json:"id,omitempty"`

	// The bandwidth for the volume
	Iops int64 `json:"iops,omitempty"`

	// The user-defined name for this volume
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// profile
	Profile *PostServersParamsBodyBootVolumeAttachmentVolumeProfile `json:"profile,omitempty"`

	// resource group
	ResourceGroup *PostServersParamsBodyBootVolumeAttachmentVolumeResourceGroup `json:"resource_group,omitempty"`

	// source snapshot
	SourceSnapshot *PostServersParamsBodyBootVolumeAttachmentVolumeSourceSnapshot `json:"source_snapshot,omitempty"`

	// A collection of tags for this resource
	Tags []string `json:"tags,omitempty"`

	// The volume type
	Type string `json:"type,omitempty"`
}

// Validate validates this post servers params body boot volume attachment volume
func (m *PostServersParamsBodyBootVolumeAttachmentVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingTerm(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCapacity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEncryptionKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourceSnapshot(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postServersParamsBodyBootVolumeAttachmentVolumeTypeBillingTermPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hourly","monthly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postServersParamsBodyBootVolumeAttachmentVolumeTypeBillingTermPropEnum = append(postServersParamsBodyBootVolumeAttachmentVolumeTypeBillingTermPropEnum, v)
	}
}

const (
	// PostServersParamsBodyBootVolumeAttachmentVolumeBillingTermHourly captures enum value "hourly"
	PostServersParamsBodyBootVolumeAttachmentVolumeBillingTermHourly string = "hourly"
	// PostServersParamsBodyBootVolumeAttachmentVolumeBillingTermMonthly captures enum value "monthly"
	PostServersParamsBodyBootVolumeAttachmentVolumeBillingTermMonthly string = "monthly"
)

// prop value enum
func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateBillingTermEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postServersParamsBodyBootVolumeAttachmentVolumeTypeBillingTermPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateBillingTerm(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingTerm) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingTermEnum("billing_term", "body", *m.BillingTerm); err != nil {
		return err
	}

	return nil
}

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateCapacity(formats strfmt.Registry) error {

	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if err := validate.MinimumInt("capacity", "body", int64(m.Capacity), 10, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("capacity", "body", int64(m.Capacity), 64000, false); err != nil {
		return err
	}

	return nil
}

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateEncryptionKey(formats strfmt.Registry) error {

	if swag.IsZero(m.EncryptionKey) { // not required
		return nil
	}

	if m.EncryptionKey != nil {

		if err := m.EncryptionKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption_key")
			}
			return err
		}
	}

	return nil
}

var postServersParamsBodyBootVolumeAttachmentVolumeTypeIopsPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1000,10000,100000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postServersParamsBodyBootVolumeAttachmentVolumeTypeIopsPropEnum = append(postServersParamsBodyBootVolumeAttachmentVolumeTypeIopsPropEnum, v)
	}
}

// prop value enum
func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateIopsEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, postServersParamsBodyBootVolumeAttachmentVolumeTypeIopsPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateIops(formats strfmt.Registry) error {

	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	// value enum
	if err := m.validateIopsEnum("iops", "body", m.Iops); err != nil {
		return err
	}

	return nil
}

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {

		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateResourceGroup(formats strfmt.Registry) error {

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

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateSourceSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceSnapshot) { // not required
		return nil
	}

	if m.SourceSnapshot != nil {

		if err := m.SourceSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_snapshot")
			}
			return err
		}
	}

	return nil
}

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	return nil
}

var postServersParamsBodyBootVolumeAttachmentVolumeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["boot","data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postServersParamsBodyBootVolumeAttachmentVolumeTypeTypePropEnum = append(postServersParamsBodyBootVolumeAttachmentVolumeTypeTypePropEnum, v)
	}
}

const (
	// PostServersParamsBodyBootVolumeAttachmentVolumeTypeBoot captures enum value "boot"
	PostServersParamsBodyBootVolumeAttachmentVolumeTypeBoot string = "boot"
	// PostServersParamsBodyBootVolumeAttachmentVolumeTypeData captures enum value "data"
	PostServersParamsBodyBootVolumeAttachmentVolumeTypeData string = "data"
)

// prop value enum
func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postServersParamsBodyBootVolumeAttachmentVolumeTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostServersParamsBodyBootVolumeAttachmentVolume) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostServersParamsBodyBootVolumeAttachmentVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostServersParamsBodyBootVolumeAttachmentVolume) UnmarshalBinary(b []byte) error {
	var res PostServersParamsBodyBootVolumeAttachmentVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
