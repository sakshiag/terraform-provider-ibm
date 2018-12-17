// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

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

// NewGetIpsecPoliciesParams creates a new GetIpsecPoliciesParams object
// with the default values initialized.
func NewGetIpsecPoliciesParams() *GetIpsecPoliciesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetIpsecPoliciesParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIpsecPoliciesParamsWithTimeout creates a new GetIpsecPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIpsecPoliciesParamsWithTimeout(timeout time.Duration) *GetIpsecPoliciesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetIpsecPoliciesParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetIpsecPoliciesParamsWithContext creates a new GetIpsecPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIpsecPoliciesParamsWithContext(ctx context.Context) *GetIpsecPoliciesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetIpsecPoliciesParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetIpsecPoliciesParamsWithHTTPClient creates a new GetIpsecPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIpsecPoliciesParamsWithHTTPClient(client *http.Client) *GetIpsecPoliciesParams {
	var (
		limitDefault = int32(50)
	)
	return &GetIpsecPoliciesParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetIpsecPoliciesParams contains all the parameters to send to the API endpoint
for the get ipsec policies operation typically these are written to a http.Request
*/
type GetIpsecPoliciesParams struct {

	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ipsec policies params
func (o *GetIpsecPoliciesParams) WithTimeout(timeout time.Duration) *GetIpsecPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ipsec policies params
func (o *GetIpsecPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ipsec policies params
func (o *GetIpsecPoliciesParams) WithContext(ctx context.Context) *GetIpsecPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ipsec policies params
func (o *GetIpsecPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ipsec policies params
func (o *GetIpsecPoliciesParams) WithHTTPClient(client *http.Client) *GetIpsecPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ipsec policies params
func (o *GetIpsecPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get ipsec policies params
func (o *GetIpsecPoliciesParams) WithLimit(limit *int32) *GetIpsecPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get ipsec policies params
func (o *GetIpsecPoliciesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithStart adds the start to the get ipsec policies params
func (o *GetIpsecPoliciesParams) WithStart(start *string) *GetIpsecPoliciesParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get ipsec policies params
func (o *GetIpsecPoliciesParams) SetStart(start *string) {
	o.Start = start
}

// WithVersion adds the version to the get ipsec policies params
func (o *GetIpsecPoliciesParams) WithVersion(version string) *GetIpsecPoliciesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get ipsec policies params
func (o *GetIpsecPoliciesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetIpsecPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
