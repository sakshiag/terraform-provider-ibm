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

// NewDeleteNetworkAclsIDTagsTagNameParams creates a new DeleteNetworkAclsIDTagsTagNameParams object
// with the default values initialized.
func NewDeleteNetworkAclsIDTagsTagNameParams() *DeleteNetworkAclsIDTagsTagNameParams {
	var ()
	return &DeleteNetworkAclsIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkAclsIDTagsTagNameParamsWithTimeout creates a new DeleteNetworkAclsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNetworkAclsIDTagsTagNameParamsWithTimeout(timeout time.Duration) *DeleteNetworkAclsIDTagsTagNameParams {
	var ()
	return &DeleteNetworkAclsIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewDeleteNetworkAclsIDTagsTagNameParamsWithContext creates a new DeleteNetworkAclsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNetworkAclsIDTagsTagNameParamsWithContext(ctx context.Context) *DeleteNetworkAclsIDTagsTagNameParams {
	var ()
	return &DeleteNetworkAclsIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewDeleteNetworkAclsIDTagsTagNameParamsWithHTTPClient creates a new DeleteNetworkAclsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNetworkAclsIDTagsTagNameParamsWithHTTPClient(client *http.Client) *DeleteNetworkAclsIDTagsTagNameParams {
	var ()
	return &DeleteNetworkAclsIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*DeleteNetworkAclsIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the delete network acls ID tags tag name operation typically these are written to a http.Request
*/
type DeleteNetworkAclsIDTagsTagNameParams struct {

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

// WithTimeout adds the timeout to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) WithTimeout(timeout time.Duration) *DeleteNetworkAclsIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) WithContext(ctx context.Context) *DeleteNetworkAclsIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) WithHTTPClient(client *http.Client) *DeleteNetworkAclsIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) WithID(id string) *DeleteNetworkAclsIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) WithTagName(tagName string) *DeleteNetworkAclsIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) WithVersion(version string) *DeleteNetworkAclsIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete network acls ID tags tag name params
func (o *DeleteNetworkAclsIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkAclsIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
