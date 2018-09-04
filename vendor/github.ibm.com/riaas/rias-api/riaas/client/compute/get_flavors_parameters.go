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

// NewGetFlavorsParams creates a new GetFlavorsParams object
// with the default values initialized.
func NewGetFlavorsParams() *GetFlavorsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetFlavorsParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFlavorsParamsWithTimeout creates a new GetFlavorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFlavorsParamsWithTimeout(timeout time.Duration) *GetFlavorsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetFlavorsParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetFlavorsParamsWithContext creates a new GetFlavorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFlavorsParamsWithContext(ctx context.Context) *GetFlavorsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetFlavorsParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetFlavorsParamsWithHTTPClient creates a new GetFlavorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFlavorsParamsWithHTTPClient(client *http.Client) *GetFlavorsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetFlavorsParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetFlavorsParams contains all the parameters to send to the API endpoint
for the get flavors operation typically these are written to a http.Request
*/
type GetFlavorsParams struct {

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

// WithTimeout adds the timeout to the get flavors params
func (o *GetFlavorsParams) WithTimeout(timeout time.Duration) *GetFlavorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get flavors params
func (o *GetFlavorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get flavors params
func (o *GetFlavorsParams) WithContext(ctx context.Context) *GetFlavorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get flavors params
func (o *GetFlavorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get flavors params
func (o *GetFlavorsParams) WithHTTPClient(client *http.Client) *GetFlavorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get flavors params
func (o *GetFlavorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get flavors params
func (o *GetFlavorsParams) WithLimit(limit *int32) *GetFlavorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get flavors params
func (o *GetFlavorsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithStart adds the start to the get flavors params
func (o *GetFlavorsParams) WithStart(start *string) *GetFlavorsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get flavors params
func (o *GetFlavorsParams) SetStart(start *string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetFlavorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
