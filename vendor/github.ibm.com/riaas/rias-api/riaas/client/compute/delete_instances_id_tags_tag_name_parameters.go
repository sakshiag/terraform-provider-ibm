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

// NewDeleteInstancesIDTagsTagNameParams creates a new DeleteInstancesIDTagsTagNameParams object
// with the default values initialized.
func NewDeleteInstancesIDTagsTagNameParams() *DeleteInstancesIDTagsTagNameParams {
	var ()
	return &DeleteInstancesIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesIDTagsTagNameParamsWithTimeout creates a new DeleteInstancesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesIDTagsTagNameParamsWithTimeout(timeout time.Duration) *DeleteInstancesIDTagsTagNameParams {
	var ()
	return &DeleteInstancesIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewDeleteInstancesIDTagsTagNameParamsWithContext creates a new DeleteInstancesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesIDTagsTagNameParamsWithContext(ctx context.Context) *DeleteInstancesIDTagsTagNameParams {
	var ()
	return &DeleteInstancesIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewDeleteInstancesIDTagsTagNameParamsWithHTTPClient creates a new DeleteInstancesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesIDTagsTagNameParamsWithHTTPClient(client *http.Client) *DeleteInstancesIDTagsTagNameParams {
	var ()
	return &DeleteInstancesIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*DeleteInstancesIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the delete instances ID tags tag name operation typically these are written to a http.Request
*/
type DeleteInstancesIDTagsTagNameParams struct {

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

// WithTimeout adds the timeout to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) WithTimeout(timeout time.Duration) *DeleteInstancesIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) WithContext(ctx context.Context) *DeleteInstancesIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) WithHTTPClient(client *http.Client) *DeleteInstancesIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) WithID(id string) *DeleteInstancesIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) WithTagName(tagName string) *DeleteInstancesIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) WithVersion(version string) *DeleteInstancesIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete instances ID tags tag name params
func (o *DeleteInstancesIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
