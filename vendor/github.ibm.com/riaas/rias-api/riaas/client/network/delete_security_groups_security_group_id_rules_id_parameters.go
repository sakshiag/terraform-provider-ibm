// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSecurityGroupsSecurityGroupIDRulesIDParams creates a new DeleteSecurityGroupsSecurityGroupIDRulesIDParams object
// with the default values initialized.
func NewDeleteSecurityGroupsSecurityGroupIDRulesIDParams() *DeleteSecurityGroupsSecurityGroupIDRulesIDParams {
	var ()
	return &DeleteSecurityGroupsSecurityGroupIDRulesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecurityGroupsSecurityGroupIDRulesIDParamsWithTimeout creates a new DeleteSecurityGroupsSecurityGroupIDRulesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSecurityGroupsSecurityGroupIDRulesIDParamsWithTimeout(timeout time.Duration) *DeleteSecurityGroupsSecurityGroupIDRulesIDParams {
	var ()
	return &DeleteSecurityGroupsSecurityGroupIDRulesIDParams{

		timeout: timeout,
	}
}

// NewDeleteSecurityGroupsSecurityGroupIDRulesIDParamsWithContext creates a new DeleteSecurityGroupsSecurityGroupIDRulesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSecurityGroupsSecurityGroupIDRulesIDParamsWithContext(ctx context.Context) *DeleteSecurityGroupsSecurityGroupIDRulesIDParams {
	var ()
	return &DeleteSecurityGroupsSecurityGroupIDRulesIDParams{

		Context: ctx,
	}
}

// NewDeleteSecurityGroupsSecurityGroupIDRulesIDParamsWithHTTPClient creates a new DeleteSecurityGroupsSecurityGroupIDRulesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSecurityGroupsSecurityGroupIDRulesIDParamsWithHTTPClient(client *http.Client) *DeleteSecurityGroupsSecurityGroupIDRulesIDParams {
	var ()
	return &DeleteSecurityGroupsSecurityGroupIDRulesIDParams{
		HTTPClient: client,
	}
}

/*DeleteSecurityGroupsSecurityGroupIDRulesIDParams contains all the parameters to send to the API endpoint
for the delete security groups security group ID rules ID operation typically these are written to a http.Request
*/
type DeleteSecurityGroupsSecurityGroupIDRulesIDParams struct {

	/*ID
	  The rule identifier

	*/
	ID string
	/*SecurityGroupID
	  The security group identifier

	*/
	SecurityGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) WithTimeout(timeout time.Duration) *DeleteSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) WithContext(ctx context.Context) *DeleteSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) WithHTTPClient(client *http.Client) *DeleteSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) WithID(id string) *DeleteSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) SetID(id string) {
	o.ID = id
}

// WithSecurityGroupID adds the securityGroupID to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) WithSecurityGroupID(securityGroupID string) *DeleteSecurityGroupsSecurityGroupIDRulesIDParams {
	o.SetSecurityGroupID(securityGroupID)
	return o
}

// SetSecurityGroupID adds the securityGroupId to the delete security groups security group ID rules ID params
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) SetSecurityGroupID(securityGroupID string) {
	o.SecurityGroupID = securityGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param security_group_id
	if err := r.SetPathParam("security_group_id", o.SecurityGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
