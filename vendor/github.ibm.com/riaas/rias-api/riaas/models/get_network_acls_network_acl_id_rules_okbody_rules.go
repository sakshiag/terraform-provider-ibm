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

// GetNetworkAclsNetworkACLIDRulesOKBodyRules Collection of rules
// swagger:model getNetworkAclsNetworkAclIdRulesOKBodyRules
type GetNetworkAclsNetworkACLIDRulesOKBodyRules []*NetworkACLRule

// Validate validates this get network acls network Acl Id rules o k body rules
func (m GetNetworkAclsNetworkACLIDRulesOKBodyRules) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {

			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
