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

// ServerCPU CPU
//
// The CPU for this server
// swagger:model serverCpu
type ServerCPU struct {

	// The CPU architecture
	Architecture string `json:"architecture,omitempty"`

	// The number of logical CPU cores per CPU
	// Minimum: 1
	Cores int64 `json:"cores,omitempty"`

	// The speed of the CPU in megahertz
	// Minimum: 500
	// Multiple Of: 500
	Frequency int64 `json:"frequency,omitempty"`
}

// Validate validates this server Cpu
func (m *ServerCPU) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchitecture(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCores(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serverCpuTypeArchitecturePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["amd64","powerpc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverCpuTypeArchitecturePropEnum = append(serverCpuTypeArchitecturePropEnum, v)
	}
}

const (
	// ServerCPUArchitectureAmd64 captures enum value "amd64"
	ServerCPUArchitectureAmd64 string = "amd64"
	// ServerCPUArchitecturePowerpc captures enum value "powerpc"
	ServerCPUArchitecturePowerpc string = "powerpc"
)

// prop value enum
func (m *ServerCPU) validateArchitectureEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serverCpuTypeArchitecturePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServerCPU) validateArchitecture(formats strfmt.Registry) error {

	if swag.IsZero(m.Architecture) { // not required
		return nil
	}

	// value enum
	if err := m.validateArchitectureEnum("architecture", "body", m.Architecture); err != nil {
		return err
	}

	return nil
}

func (m *ServerCPU) validateCores(formats strfmt.Registry) error {

	if swag.IsZero(m.Cores) { // not required
		return nil
	}

	if err := validate.MinimumInt("cores", "body", int64(m.Cores), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ServerCPU) validateFrequency(formats strfmt.Registry) error {

	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	if err := validate.MinimumInt("frequency", "body", int64(m.Frequency), 500, false); err != nil {
		return err
	}

	if err := validate.MultipleOf("frequency", "body", float64(m.Frequency), 500); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerCPU) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerCPU) UnmarshalBinary(b []byte) error {
	var res ServerCPU
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
