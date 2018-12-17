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

// NewGetNetworkAclsIDTagsParams creates a new GetNetworkAclsIDTagsParams object
// with the default values initialized.
func NewGetNetworkAclsIDTagsParams() *GetNetworkAclsIDTagsParams {
	var ()
	return &GetNetworkAclsIDTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkAclsIDTagsParamsWithTimeout creates a new GetNetworkAclsIDTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkAclsIDTagsParamsWithTimeout(timeout time.Duration) *GetNetworkAclsIDTagsParams {
	var ()
	return &GetNetworkAclsIDTagsParams{

		timeout: timeout,
	}
}

// NewGetNetworkAclsIDTagsParamsWithContext creates a new GetNetworkAclsIDTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkAclsIDTagsParamsWithContext(ctx context.Context) *GetNetworkAclsIDTagsParams {
	var ()
	return &GetNetworkAclsIDTagsParams{

		Context: ctx,
	}
}

// NewGetNetworkAclsIDTagsParamsWithHTTPClient creates a new GetNetworkAclsIDTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkAclsIDTagsParamsWithHTTPClient(client *http.Client) *GetNetworkAclsIDTagsParams {
	var ()
	return &GetNetworkAclsIDTagsParams{
		HTTPClient: client,
	}
}

/*GetNetworkAclsIDTagsParams contains all the parameters to send to the API endpoint
for the get network acls ID tags operation typically these are written to a http.Request
*/
type GetNetworkAclsIDTagsParams struct {

	/*ID
	  The resource identifier

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

// WithTimeout adds the timeout to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) WithTimeout(timeout time.Duration) *GetNetworkAclsIDTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) WithContext(ctx context.Context) *GetNetworkAclsIDTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) WithHTTPClient(client *http.Client) *GetNetworkAclsIDTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) WithID(id string) *GetNetworkAclsIDTagsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) WithVersion(version string) *GetNetworkAclsIDTagsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get network acls ID tags params
func (o *GetNetworkAclsIDTagsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkAclsIDTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
