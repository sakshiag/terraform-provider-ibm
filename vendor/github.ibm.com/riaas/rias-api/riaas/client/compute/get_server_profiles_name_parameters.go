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

// NewGetServerProfilesNameParams creates a new GetServerProfilesNameParams object
// with the default values initialized.
func NewGetServerProfilesNameParams() *GetServerProfilesNameParams {
	var ()
	return &GetServerProfilesNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerProfilesNameParamsWithTimeout creates a new GetServerProfilesNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServerProfilesNameParamsWithTimeout(timeout time.Duration) *GetServerProfilesNameParams {
	var ()
	return &GetServerProfilesNameParams{

		timeout: timeout,
	}
}

// NewGetServerProfilesNameParamsWithContext creates a new GetServerProfilesNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServerProfilesNameParamsWithContext(ctx context.Context) *GetServerProfilesNameParams {
	var ()
	return &GetServerProfilesNameParams{

		Context: ctx,
	}
}

// NewGetServerProfilesNameParamsWithHTTPClient creates a new GetServerProfilesNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServerProfilesNameParamsWithHTTPClient(client *http.Client) *GetServerProfilesNameParams {
	var ()
	return &GetServerProfilesNameParams{
		HTTPClient: client,
	}
}

/*GetServerProfilesNameParams contains all the parameters to send to the API endpoint
for the get server profiles name operation typically these are written to a http.Request
*/
type GetServerProfilesNameParams struct {

	/*Name
	  The server profile name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get server profiles name params
func (o *GetServerProfilesNameParams) WithTimeout(timeout time.Duration) *GetServerProfilesNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server profiles name params
func (o *GetServerProfilesNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server profiles name params
func (o *GetServerProfilesNameParams) WithContext(ctx context.Context) *GetServerProfilesNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server profiles name params
func (o *GetServerProfilesNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server profiles name params
func (o *GetServerProfilesNameParams) WithHTTPClient(client *http.Client) *GetServerProfilesNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server profiles name params
func (o *GetServerProfilesNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get server profiles name params
func (o *GetServerProfilesNameParams) WithName(name string) *GetServerProfilesNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get server profiles name params
func (o *GetServerProfilesNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerProfilesNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
