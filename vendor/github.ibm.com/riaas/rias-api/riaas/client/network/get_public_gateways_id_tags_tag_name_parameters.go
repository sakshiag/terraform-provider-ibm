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

// NewGetPublicGatewaysIDTagsTagNameParams creates a new GetPublicGatewaysIDTagsTagNameParams object
// with the default values initialized.
func NewGetPublicGatewaysIDTagsTagNameParams() *GetPublicGatewaysIDTagsTagNameParams {
	var ()
	return &GetPublicGatewaysIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicGatewaysIDTagsTagNameParamsWithTimeout creates a new GetPublicGatewaysIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicGatewaysIDTagsTagNameParamsWithTimeout(timeout time.Duration) *GetPublicGatewaysIDTagsTagNameParams {
	var ()
	return &GetPublicGatewaysIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewGetPublicGatewaysIDTagsTagNameParamsWithContext creates a new GetPublicGatewaysIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicGatewaysIDTagsTagNameParamsWithContext(ctx context.Context) *GetPublicGatewaysIDTagsTagNameParams {
	var ()
	return &GetPublicGatewaysIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewGetPublicGatewaysIDTagsTagNameParamsWithHTTPClient creates a new GetPublicGatewaysIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicGatewaysIDTagsTagNameParamsWithHTTPClient(client *http.Client) *GetPublicGatewaysIDTagsTagNameParams {
	var ()
	return &GetPublicGatewaysIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*GetPublicGatewaysIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the get public gateways ID tags tag name operation typically these are written to a http.Request
*/
type GetPublicGatewaysIDTagsTagNameParams struct {

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

// WithTimeout adds the timeout to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) WithTimeout(timeout time.Duration) *GetPublicGatewaysIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) WithContext(ctx context.Context) *GetPublicGatewaysIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) WithHTTPClient(client *http.Client) *GetPublicGatewaysIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) WithID(id string) *GetPublicGatewaysIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) WithTagName(tagName string) *GetPublicGatewaysIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the get public gateways ID tags tag name params
func (o *GetPublicGatewaysIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicGatewaysIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
