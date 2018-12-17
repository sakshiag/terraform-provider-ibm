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

// NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams creates a new GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams object
// with the default values initialized.
func NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams() *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	var ()
	return &GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithTimeout creates a new GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithTimeout(timeout time.Duration) *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	var ()
	return &GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams{

		timeout: timeout,
	}
}

// NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithContext creates a new GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithContext(ctx context.Context) *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	var ()
	return &GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams{

		Context: ctx,
	}
}

// NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithHTTPClient creates a new GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithHTTPClient(client *http.Client) *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	var ()
	return &GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams{
		HTTPClient: client,
	}
}

/*GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams contains all the parameters to send to the API endpoint
for the get security groups security group ID network interfaces ID operation typically these are written to a http.Request
*/
type GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams struct {

	/*ID
	  The network interface identifier

	*/
	ID string
	/*SecurityGroupID
	  The security group identifier

	*/
	SecurityGroupID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithTimeout(timeout time.Duration) *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithContext(ctx context.Context) *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithHTTPClient(client *http.Client) *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithID(id string) *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetID(id string) {
	o.ID = id
}

// WithSecurityGroupID adds the securityGroupID to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithSecurityGroupID(securityGroupID string) *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetSecurityGroupID(securityGroupID)
	return o
}

// SetSecurityGroupID adds the securityGroupId to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetSecurityGroupID(securityGroupID string) {
	o.SecurityGroupID = securityGroupID
}

// WithVersion adds the version to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithVersion(version string) *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get security groups security group ID network interfaces ID params
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
