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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServerProfilesParams creates a new GetServerProfilesParams object
// with the default values initialized.
func NewGetServerProfilesParams() *GetServerProfilesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetServerProfilesParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerProfilesParamsWithTimeout creates a new GetServerProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServerProfilesParamsWithTimeout(timeout time.Duration) *GetServerProfilesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetServerProfilesParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetServerProfilesParamsWithContext creates a new GetServerProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServerProfilesParamsWithContext(ctx context.Context) *GetServerProfilesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetServerProfilesParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetServerProfilesParamsWithHTTPClient creates a new GetServerProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServerProfilesParamsWithHTTPClient(client *http.Client) *GetServerProfilesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetServerProfilesParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetServerProfilesParams contains all the parameters to send to the API endpoint
for the get server profiles operation typically these are written to a http.Request
*/
type GetServerProfilesParams struct {

	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get server profiles params
func (o *GetServerProfilesParams) WithTimeout(timeout time.Duration) *GetServerProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server profiles params
func (o *GetServerProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server profiles params
func (o *GetServerProfilesParams) WithContext(ctx context.Context) *GetServerProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server profiles params
func (o *GetServerProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server profiles params
func (o *GetServerProfilesParams) WithHTTPClient(client *http.Client) *GetServerProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server profiles params
func (o *GetServerProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get server profiles params
func (o *GetServerProfilesParams) WithLimit(limit *int32) *GetServerProfilesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get server profiles params
func (o *GetServerProfilesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithStart adds the start to the get server profiles params
func (o *GetServerProfilesParams) WithStart(start *string) *GetServerProfilesParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get server profiles params
func (o *GetServerProfilesParams) SetStart(start *string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
