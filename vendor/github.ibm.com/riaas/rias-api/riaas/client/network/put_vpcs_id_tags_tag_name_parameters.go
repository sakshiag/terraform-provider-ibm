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

// NewPutVpcsIDTagsTagNameParams creates a new PutVpcsIDTagsTagNameParams object
// with the default values initialized.
func NewPutVpcsIDTagsTagNameParams() *PutVpcsIDTagsTagNameParams {
	var ()
	return &PutVpcsIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutVpcsIDTagsTagNameParamsWithTimeout creates a new PutVpcsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutVpcsIDTagsTagNameParamsWithTimeout(timeout time.Duration) *PutVpcsIDTagsTagNameParams {
	var ()
	return &PutVpcsIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewPutVpcsIDTagsTagNameParamsWithContext creates a new PutVpcsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutVpcsIDTagsTagNameParamsWithContext(ctx context.Context) *PutVpcsIDTagsTagNameParams {
	var ()
	return &PutVpcsIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewPutVpcsIDTagsTagNameParamsWithHTTPClient creates a new PutVpcsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutVpcsIDTagsTagNameParamsWithHTTPClient(client *http.Client) *PutVpcsIDTagsTagNameParams {
	var ()
	return &PutVpcsIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*PutVpcsIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the put vpcs ID tags tag name operation typically these are written to a http.Request
*/
type PutVpcsIDTagsTagNameParams struct {

	/*ID
	  The resource identifier

	*/
	ID string
	/*TagName
	  The name of the tag

	*/
	TagName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) WithTimeout(timeout time.Duration) *PutVpcsIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) WithContext(ctx context.Context) *PutVpcsIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) WithHTTPClient(client *http.Client) *PutVpcsIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) WithID(id string) *PutVpcsIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) WithTagName(tagName string) *PutVpcsIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the put vpcs ID tags tag name params
func (o *PutVpcsIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WriteToRequest writes these params to a swagger request
func (o *PutVpcsIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
