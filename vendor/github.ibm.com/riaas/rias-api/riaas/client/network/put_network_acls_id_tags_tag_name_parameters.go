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

// NewPutNetworkAclsIDTagsTagNameParams creates a new PutNetworkAclsIDTagsTagNameParams object
// with the default values initialized.
func NewPutNetworkAclsIDTagsTagNameParams() *PutNetworkAclsIDTagsTagNameParams {
	var ()
	return &PutNetworkAclsIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworkAclsIDTagsTagNameParamsWithTimeout creates a new PutNetworkAclsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNetworkAclsIDTagsTagNameParamsWithTimeout(timeout time.Duration) *PutNetworkAclsIDTagsTagNameParams {
	var ()
	return &PutNetworkAclsIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewPutNetworkAclsIDTagsTagNameParamsWithContext creates a new PutNetworkAclsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNetworkAclsIDTagsTagNameParamsWithContext(ctx context.Context) *PutNetworkAclsIDTagsTagNameParams {
	var ()
	return &PutNetworkAclsIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewPutNetworkAclsIDTagsTagNameParamsWithHTTPClient creates a new PutNetworkAclsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNetworkAclsIDTagsTagNameParamsWithHTTPClient(client *http.Client) *PutNetworkAclsIDTagsTagNameParams {
	var ()
	return &PutNetworkAclsIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*PutNetworkAclsIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the put network acls ID tags tag name operation typically these are written to a http.Request
*/
type PutNetworkAclsIDTagsTagNameParams struct {

	/*ID
	  The resource identifier

	*/
	ID string
	/*TagName
	  The name of the tag

	*/
	TagName string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) WithTimeout(timeout time.Duration) *PutNetworkAclsIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) WithContext(ctx context.Context) *PutNetworkAclsIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) WithHTTPClient(client *http.Client) *PutNetworkAclsIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) WithID(id string) *PutNetworkAclsIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) WithTagName(tagName string) *PutNetworkAclsIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) WithVersion(version string) *PutNetworkAclsIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the put network acls ID tags tag name params
func (o *PutNetworkAclsIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworkAclsIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param tag_name
	if err := r.SetPathParam("tag_name", o.TagName); err != nil {
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
