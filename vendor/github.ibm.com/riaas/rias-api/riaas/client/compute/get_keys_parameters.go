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

// NewGetKeysParams creates a new GetKeysParams object
// with the default values initialized.
func NewGetKeysParams() *GetKeysParams {
	var (
		limitDefault = int32(50)
	)
	return &GetKeysParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeysParamsWithTimeout creates a new GetKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeysParamsWithTimeout(timeout time.Duration) *GetKeysParams {
	var (
		limitDefault = int32(50)
	)
	return &GetKeysParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetKeysParamsWithContext creates a new GetKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeysParamsWithContext(ctx context.Context) *GetKeysParams {
	var (
		limitDefault = int32(50)
	)
	return &GetKeysParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetKeysParamsWithHTTPClient creates a new GetKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeysParamsWithHTTPClient(client *http.Client) *GetKeysParams {
	var (
		limitDefault = int32(50)
	)
	return &GetKeysParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetKeysParams contains all the parameters to send to the API endpoint
for the get keys operation typically these are written to a http.Request
*/
type GetKeysParams struct {

	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*ResourceGroupID
	  Filters the collection to resources within the resource group of the specified identifier

	*/
	ResourceGroupID *string
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string
	/*Tag
	  Filters the collection to resources containing the specified tag

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get keys params
func (o *GetKeysParams) WithTimeout(timeout time.Duration) *GetKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get keys params
func (o *GetKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get keys params
func (o *GetKeysParams) WithContext(ctx context.Context) *GetKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get keys params
func (o *GetKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get keys params
func (o *GetKeysParams) WithHTTPClient(client *http.Client) *GetKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get keys params
func (o *GetKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get keys params
func (o *GetKeysParams) WithLimit(limit *int32) *GetKeysParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get keys params
func (o *GetKeysParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithResourceGroupID adds the resourceGroupID to the get keys params
func (o *GetKeysParams) WithResourceGroupID(resourceGroupID *string) *GetKeysParams {
	o.SetResourceGroupID(resourceGroupID)
	return o
}

// SetResourceGroupID adds the resourceGroupId to the get keys params
func (o *GetKeysParams) SetResourceGroupID(resourceGroupID *string) {
	o.ResourceGroupID = resourceGroupID
}

// WithStart adds the start to the get keys params
func (o *GetKeysParams) WithStart(start *string) *GetKeysParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get keys params
func (o *GetKeysParams) SetStart(start *string) {
	o.Start = start
}

// WithTag adds the tag to the get keys params
func (o *GetKeysParams) WithTag(tag *string) *GetKeysParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get keys params
func (o *GetKeysParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ResourceGroupID != nil {

		// query param resource_group.id
		var qrResourceGroupID string
		if o.ResourceGroupID != nil {
			qrResourceGroupID = *o.ResourceGroupID
		}
		qResourceGroupID := qrResourceGroupID
		if qResourceGroupID != "" {
			if err := r.SetQueryParam("resource_group.id", qResourceGroupID); err != nil {
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

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
