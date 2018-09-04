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

// NewGetServersIDParams creates a new GetServersIDParams object
// with the default values initialized.
func NewGetServersIDParams() *GetServersIDParams {
	var ()
	return &GetServersIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServersIDParamsWithTimeout creates a new GetServersIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServersIDParamsWithTimeout(timeout time.Duration) *GetServersIDParams {
	var ()
	return &GetServersIDParams{

		timeout: timeout,
	}
}

// NewGetServersIDParamsWithContext creates a new GetServersIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServersIDParamsWithContext(ctx context.Context) *GetServersIDParams {
	var ()
	return &GetServersIDParams{

		Context: ctx,
	}
}

// NewGetServersIDParamsWithHTTPClient creates a new GetServersIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServersIDParamsWithHTTPClient(client *http.Client) *GetServersIDParams {
	var ()
	return &GetServersIDParams{
		HTTPClient: client,
	}
}

/*GetServersIDParams contains all the parameters to send to the API endpoint
for the get servers ID operation typically these are written to a http.Request
*/
type GetServersIDParams struct {

	/*ID
	  The server identifier

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get servers ID params
func (o *GetServersIDParams) WithTimeout(timeout time.Duration) *GetServersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get servers ID params
func (o *GetServersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get servers ID params
func (o *GetServersIDParams) WithContext(ctx context.Context) *GetServersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get servers ID params
func (o *GetServersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get servers ID params
func (o *GetServersIDParams) WithHTTPClient(client *http.Client) *GetServersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get servers ID params
func (o *GetServersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get servers ID params
func (o *GetServersIDParams) WithID(id string) *GetServersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get servers ID params
func (o *GetServersIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetServersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
