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

// NewGetInstancesIDParams creates a new GetInstancesIDParams object
// with the default values initialized.
func NewGetInstancesIDParams() *GetInstancesIDParams {
	var ()
	return &GetInstancesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesIDParamsWithTimeout creates a new GetInstancesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesIDParamsWithTimeout(timeout time.Duration) *GetInstancesIDParams {
	var ()
	return &GetInstancesIDParams{

		timeout: timeout,
	}
}

// NewGetInstancesIDParamsWithContext creates a new GetInstancesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesIDParamsWithContext(ctx context.Context) *GetInstancesIDParams {
	var ()
	return &GetInstancesIDParams{

		Context: ctx,
	}
}

// NewGetInstancesIDParamsWithHTTPClient creates a new GetInstancesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesIDParamsWithHTTPClient(client *http.Client) *GetInstancesIDParams {
	var ()
	return &GetInstancesIDParams{
		HTTPClient: client,
	}
}

/*GetInstancesIDParams contains all the parameters to send to the API endpoint
for the get instances ID operation typically these are written to a http.Request
*/
type GetInstancesIDParams struct {

	/*ID
	  The instance identifier

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

// WithTimeout adds the timeout to the get instances ID params
func (o *GetInstancesIDParams) WithTimeout(timeout time.Duration) *GetInstancesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances ID params
func (o *GetInstancesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances ID params
func (o *GetInstancesIDParams) WithContext(ctx context.Context) *GetInstancesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances ID params
func (o *GetInstancesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances ID params
func (o *GetInstancesIDParams) WithHTTPClient(client *http.Client) *GetInstancesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances ID params
func (o *GetInstancesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get instances ID params
func (o *GetInstancesIDParams) WithID(id string) *GetInstancesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get instances ID params
func (o *GetInstancesIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get instances ID params
func (o *GetInstancesIDParams) WithVersion(version string) *GetInstancesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get instances ID params
func (o *GetInstancesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
