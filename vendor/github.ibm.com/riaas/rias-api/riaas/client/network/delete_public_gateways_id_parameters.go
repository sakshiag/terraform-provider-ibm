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

// NewDeletePublicGatewaysIDParams creates a new DeletePublicGatewaysIDParams object
// with the default values initialized.
func NewDeletePublicGatewaysIDParams() *DeletePublicGatewaysIDParams {
	var ()
	return &DeletePublicGatewaysIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePublicGatewaysIDParamsWithTimeout creates a new DeletePublicGatewaysIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePublicGatewaysIDParamsWithTimeout(timeout time.Duration) *DeletePublicGatewaysIDParams {
	var ()
	return &DeletePublicGatewaysIDParams{

		timeout: timeout,
	}
}

// NewDeletePublicGatewaysIDParamsWithContext creates a new DeletePublicGatewaysIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePublicGatewaysIDParamsWithContext(ctx context.Context) *DeletePublicGatewaysIDParams {
	var ()
	return &DeletePublicGatewaysIDParams{

		Context: ctx,
	}
}

// NewDeletePublicGatewaysIDParamsWithHTTPClient creates a new DeletePublicGatewaysIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePublicGatewaysIDParamsWithHTTPClient(client *http.Client) *DeletePublicGatewaysIDParams {
	var ()
	return &DeletePublicGatewaysIDParams{
		HTTPClient: client,
	}
}

/*DeletePublicGatewaysIDParams contains all the parameters to send to the API endpoint
for the delete public gateways ID operation typically these are written to a http.Request
*/
type DeletePublicGatewaysIDParams struct {

	/*ID
	  The public gateway identifier

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

// WithTimeout adds the timeout to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) WithTimeout(timeout time.Duration) *DeletePublicGatewaysIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) WithContext(ctx context.Context) *DeletePublicGatewaysIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) WithHTTPClient(client *http.Client) *DeletePublicGatewaysIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) WithID(id string) *DeletePublicGatewaysIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) WithVersion(version string) *DeletePublicGatewaysIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete public gateways ID params
func (o *DeletePublicGatewaysIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePublicGatewaysIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
