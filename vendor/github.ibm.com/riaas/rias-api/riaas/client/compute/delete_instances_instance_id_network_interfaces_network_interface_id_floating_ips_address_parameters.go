// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams creates a new DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams object
// with the default values initialized.
func NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams() *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	var ()
	return &DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParamsWithTimeout creates a new DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParamsWithTimeout(timeout time.Duration) *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	var ()
	return &DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams{

		timeout: timeout,
	}
}

// NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParamsWithContext creates a new DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParamsWithContext(ctx context.Context) *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	var ()
	return &DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams{

		Context: ctx,
	}
}

// NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParamsWithHTTPClient creates a new DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParamsWithHTTPClient(client *http.Client) *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	var ()
	return &DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams{
		HTTPClient: client,
	}
}

/*DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams contains all the parameters to send to the API endpoint
for the delete instances instance ID network interfaces network interface ID floating ips address operation typically these are written to a http.Request
*/
type DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams struct {

	/*Address
	  The floating IP address

	*/
	Address string
	/*InstanceID
	  The instance identifier

	*/
	InstanceID string
	/*NetworkInterfaceID
	  The network interface identifier

	*/
	NetworkInterfaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) WithTimeout(timeout time.Duration) *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) WithContext(ctx context.Context) *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) WithHTTPClient(client *http.Client) *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) WithAddress(address string) *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) SetAddress(address string) {
	o.Address = address
}

// WithInstanceID adds the instanceID to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) WithInstanceID(instanceID string) *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithNetworkInterfaceID adds the networkInterfaceID to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) WithNetworkInterfaceID(networkInterfaceID string) *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams {
	o.SetNetworkInterfaceID(networkInterfaceID)
	return o
}

// SetNetworkInterfaceID adds the networkInterfaceId to the delete instances instance ID network interfaces network interface ID floating ips address params
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) SetNetworkInterfaceID(networkInterfaceID string) {
	o.NetworkInterfaceID = networkInterfaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
		return err
	}

	// path param network_interface_id
	if err := r.SetPathParam("network_interface_id", o.NetworkInterfaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
