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

// PostInstancesInstanceIDActionsParamsBody instanceActionTemplate
// swagger:model postInstancesInstanceIdActionsParamsBody
type PostInstancesInstanceIDActionsParamsBody struct {

	// The type of action
	// Enum: [start stop reboot reset]
	Type string `json:"type,omitempty"`
}

// Validate validates this post instances instance Id actions params body
func (m *PostInstancesInstanceIDActionsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postInstancesInstanceIdActionsParamsBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["start","stop","reboot","reset"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postInstancesInstanceIdActionsParamsBodyTypeTypePropEnum = append(postInstancesInstanceIdActionsParamsBodyTypeTypePropEnum, v)
	}
}

const (

	// PostInstancesInstanceIDActionsParamsBodyTypeStart captures enum value "start"
	PostInstancesInstanceIDActionsParamsBodyTypeStart string = "start"

	// PostInstancesInstanceIDActionsParamsBodyTypeStop captures enum value "stop"
	PostInstancesInstanceIDActionsParamsBodyTypeStop string = "stop"

	// PostInstancesInstanceIDActionsParamsBodyTypeReboot captures enum value "reboot"
	PostInstancesInstanceIDActionsParamsBodyTypeReboot string = "reboot"

	// PostInstancesInstanceIDActionsParamsBodyTypeReset captures enum value "reset"
	PostInstancesInstanceIDActionsParamsBodyTypeReset string = "reset"
)

// prop value enum
func (m *PostInstancesInstanceIDActionsParamsBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postInstancesInstanceIdActionsParamsBodyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostInstancesInstanceIDActionsParamsBody) validateType(formats strfmt.Registry) error {

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
func (m *PostInstancesInstanceIDActionsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostInstancesInstanceIDActionsParamsBody) UnmarshalBinary(b []byte) error {
	var res PostInstancesInstanceIDActionsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
