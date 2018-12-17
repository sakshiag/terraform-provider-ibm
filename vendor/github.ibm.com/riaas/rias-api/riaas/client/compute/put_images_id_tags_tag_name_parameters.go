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

// NewPutImagesIDTagsTagNameParams creates a new PutImagesIDTagsTagNameParams object
// with the default values initialized.
func NewPutImagesIDTagsTagNameParams() *PutImagesIDTagsTagNameParams {
	var ()
	return &PutImagesIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutImagesIDTagsTagNameParamsWithTimeout creates a new PutImagesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutImagesIDTagsTagNameParamsWithTimeout(timeout time.Duration) *PutImagesIDTagsTagNameParams {
	var ()
	return &PutImagesIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewPutImagesIDTagsTagNameParamsWithContext creates a new PutImagesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutImagesIDTagsTagNameParamsWithContext(ctx context.Context) *PutImagesIDTagsTagNameParams {
	var ()
	return &PutImagesIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewPutImagesIDTagsTagNameParamsWithHTTPClient creates a new PutImagesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutImagesIDTagsTagNameParamsWithHTTPClient(client *http.Client) *PutImagesIDTagsTagNameParams {
	var ()
	return &PutImagesIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*PutImagesIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the put images ID tags tag name operation typically these are written to a http.Request
*/
type PutImagesIDTagsTagNameParams struct {

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

// WithTimeout adds the timeout to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) WithTimeout(timeout time.Duration) *PutImagesIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) WithContext(ctx context.Context) *PutImagesIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) WithHTTPClient(client *http.Client) *PutImagesIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) WithID(id string) *PutImagesIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) WithTagName(tagName string) *PutImagesIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) WithVersion(version string) *PutImagesIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the put images ID tags tag name params
func (o *PutImagesIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PutImagesIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
