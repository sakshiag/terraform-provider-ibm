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

// NewGetLoadBalancersIDStatisticsParams creates a new GetLoadBalancersIDStatisticsParams object
// with the default values initialized.
func NewGetLoadBalancersIDStatisticsParams() *GetLoadBalancersIDStatisticsParams {
	var ()
	return &GetLoadBalancersIDStatisticsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoadBalancersIDStatisticsParamsWithTimeout creates a new GetLoadBalancersIDStatisticsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoadBalancersIDStatisticsParamsWithTimeout(timeout time.Duration) *GetLoadBalancersIDStatisticsParams {
	var ()
	return &GetLoadBalancersIDStatisticsParams{

		timeout: timeout,
	}
}

// NewGetLoadBalancersIDStatisticsParamsWithContext creates a new GetLoadBalancersIDStatisticsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoadBalancersIDStatisticsParamsWithContext(ctx context.Context) *GetLoadBalancersIDStatisticsParams {
	var ()
	return &GetLoadBalancersIDStatisticsParams{

		Context: ctx,
	}
}

// NewGetLoadBalancersIDStatisticsParamsWithHTTPClient creates a new GetLoadBalancersIDStatisticsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoadBalancersIDStatisticsParamsWithHTTPClient(client *http.Client) *GetLoadBalancersIDStatisticsParams {
	var ()
	return &GetLoadBalancersIDStatisticsParams{
		HTTPClient: client,
	}
}

/*GetLoadBalancersIDStatisticsParams contains all the parameters to send to the API endpoint
for the get load balancers ID statistics operation typically these are written to a http.Request
*/
type GetLoadBalancersIDStatisticsParams struct {

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

// WithTimeout adds the timeout to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) WithTimeout(timeout time.Duration) *GetLoadBalancersIDStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) WithContext(ctx context.Context) *GetLoadBalancersIDStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) WithHTTPClient(client *http.Client) *GetLoadBalancersIDStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) WithID(id string) *GetLoadBalancersIDStatisticsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) WithVersion(version string) *GetLoadBalancersIDStatisticsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get load balancers ID statistics params
func (o *GetLoadBalancersIDStatisticsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoadBalancersIDStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
