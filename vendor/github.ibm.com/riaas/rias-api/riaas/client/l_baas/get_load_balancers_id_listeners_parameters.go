// Code generated by go-swagger; DO NOT EDIT.

package l_baas

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

// NewGetLoadBalancersIDListenersParams creates a new GetLoadBalancersIDListenersParams object
// with the default values initialized.
func NewGetLoadBalancersIDListenersParams() *GetLoadBalancersIDListenersParams {
	var ()
	return &GetLoadBalancersIDListenersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoadBalancersIDListenersParamsWithTimeout creates a new GetLoadBalancersIDListenersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoadBalancersIDListenersParamsWithTimeout(timeout time.Duration) *GetLoadBalancersIDListenersParams {
	var ()
	return &GetLoadBalancersIDListenersParams{

		timeout: timeout,
	}
}

// NewGetLoadBalancersIDListenersParamsWithContext creates a new GetLoadBalancersIDListenersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoadBalancersIDListenersParamsWithContext(ctx context.Context) *GetLoadBalancersIDListenersParams {
	var ()
	return &GetLoadBalancersIDListenersParams{

		Context: ctx,
	}
}

// NewGetLoadBalancersIDListenersParamsWithHTTPClient creates a new GetLoadBalancersIDListenersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoadBalancersIDListenersParamsWithHTTPClient(client *http.Client) *GetLoadBalancersIDListenersParams {
	var ()
	return &GetLoadBalancersIDListenersParams{
		HTTPClient: client,
	}
}

/*GetLoadBalancersIDListenersParams contains all the parameters to send to the API endpoint
for the get load balancers ID listeners operation typically these are written to a http.Request
*/
type GetLoadBalancersIDListenersParams struct {

	/*ID
	  The load balancer identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) WithTimeout(timeout time.Duration) *GetLoadBalancersIDListenersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) WithContext(ctx context.Context) *GetLoadBalancersIDListenersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) WithHTTPClient(client *http.Client) *GetLoadBalancersIDListenersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) WithID(id string) *GetLoadBalancersIDListenersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) WithVersion(version string) *GetLoadBalancersIDListenersParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get load balancers ID listeners params
func (o *GetLoadBalancersIDListenersParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoadBalancersIDListenersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
