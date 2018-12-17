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

// NewPutInstancesIDTagsTagNameParams creates a new PutInstancesIDTagsTagNameParams object
// with the default values initialized.
func NewPutInstancesIDTagsTagNameParams() *PutInstancesIDTagsTagNameParams {
	var ()
	return &PutInstancesIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutInstancesIDTagsTagNameParamsWithTimeout creates a new PutInstancesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutInstancesIDTagsTagNameParamsWithTimeout(timeout time.Duration) *PutInstancesIDTagsTagNameParams {
	var ()
	return &PutInstancesIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewPutInstancesIDTagsTagNameParamsWithContext creates a new PutInstancesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutInstancesIDTagsTagNameParamsWithContext(ctx context.Context) *PutInstancesIDTagsTagNameParams {
	var ()
	return &PutInstancesIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewPutInstancesIDTagsTagNameParamsWithHTTPClient creates a new PutInstancesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutInstancesIDTagsTagNameParamsWithHTTPClient(client *http.Client) *PutInstancesIDTagsTagNameParams {
	var ()
	return &PutInstancesIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*PutInstancesIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the put instances ID tags tag name operation typically these are written to a http.Request
*/
type PutInstancesIDTagsTagNameParams struct {

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

// WithTimeout adds the timeout to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) WithTimeout(timeout time.Duration) *PutInstancesIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) WithContext(ctx context.Context) *PutInstancesIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) WithHTTPClient(client *http.Client) *PutInstancesIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) WithID(id string) *PutInstancesIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) WithTagName(tagName string) *PutInstancesIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) WithVersion(version string) *PutInstancesIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the put instances ID tags tag name params
func (o *PutInstancesIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PutInstancesIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
