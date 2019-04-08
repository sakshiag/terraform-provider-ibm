// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams creates a new DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized.
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams() *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithTimeout creates a new DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithTimeout(timeout time.Duration) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams{

		timeout: timeout,
	}
}

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithContext creates a new DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithContext(ctx context.Context) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams{

		Context: ctx,
	}
}

// NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithHTTPClient creates a new DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParamsWithHTTPClient(client *http.Client) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	var ()
	return &DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams{
		HTTPClient: client,
	}
}

/*DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams contains all the parameters to send to the API endpoint
for the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length operation typically these are written to a http.Request
*/
type DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams struct {

	/*ID
	  The VPN connection identifier

	*/
	ID string
	/*PrefixAddress
	  The prefix address part of the CIDR

	*/
	PrefixAddress string
	/*PrefixLength
	  The prefix length part of the CIDR

	*/
	PrefixLength string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpnGatewayID
	  The VPN gateway identifier

	*/
	VpnGatewayID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithTimeout(timeout time.Duration) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithContext(ctx context.Context) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithHTTPClient(client *http.Client) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithID(id string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetID(id string) {
	o.ID = id
}

// WithPrefixAddress adds the prefixAddress to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithPrefixAddress(prefixAddress string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetPrefixAddress(prefixAddress)
	return o
}

// SetPrefixAddress adds the prefixAddress to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetPrefixAddress(prefixAddress string) {
	o.PrefixAddress = prefixAddress
}

// WithPrefixLength adds the prefixLength to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithPrefixLength(prefixLength string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetPrefixLength(prefixLength)
	return o
}

// SetPrefixLength adds the prefixLength to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetPrefixLength(prefixLength string) {
	o.PrefixLength = prefixLength
}

// WithVersion adds the version to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithVersion(version string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetVersion(version string) {
	o.Version = version
}

// WithVpnGatewayID adds the vpnGatewayID to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WithVpnGatewayID(vpnGatewayID string) *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams {
	o.SetVpnGatewayID(vpnGatewayID)
	return o
}

// SetVpnGatewayID adds the vpnGatewayId to the delete vpn gateways vpn gateway ID connections ID local cidrs prefix address prefix length params
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) SetVpnGatewayID(vpnGatewayID string) {
	o.VpnGatewayID = vpnGatewayID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param prefix_address
	if err := r.SetPathParam("prefix_address", o.PrefixAddress); err != nil {
		return err
	}

	// path param prefix_length
	if err := r.SetPathParam("prefix_length", o.PrefixLength); err != nil {
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

	// path param vpn_gateway_id
	if err := r.SetPathParam("vpn_gateway_id", o.VpnGatewayID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
