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

	"github.ibm.com/riaas/rias-api/riaas/models"
)

// NewPostVpcsParams creates a new PostVpcsParams object
// with the default values initialized.
func NewPostVpcsParams() *PostVpcsParams {
	var ()
	return &PostVpcsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVpcsParamsWithTimeout creates a new PostVpcsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVpcsParamsWithTimeout(timeout time.Duration) *PostVpcsParams {
	var ()
	return &PostVpcsParams{

		timeout: timeout,
	}
}

// NewPostVpcsParamsWithContext creates a new PostVpcsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVpcsParamsWithContext(ctx context.Context) *PostVpcsParams {
	var ()
	return &PostVpcsParams{

		Context: ctx,
	}
}

// NewPostVpcsParamsWithHTTPClient creates a new PostVpcsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVpcsParamsWithHTTPClient(client *http.Client) *PostVpcsParams {
	var ()
	return &PostVpcsParams{
		HTTPClient: client,
	}
}

/*PostVpcsParams contains all the parameters to send to the API endpoint
for the post vpcs operation typically these are written to a http.Request
*/
type PostVpcsParams struct {

	/*Body*/
	Body *models.PostVpcsParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post vpcs params
func (o *PostVpcsParams) WithTimeout(timeout time.Duration) *PostVpcsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post vpcs params
func (o *PostVpcsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post vpcs params
func (o *PostVpcsParams) WithContext(ctx context.Context) *PostVpcsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post vpcs params
func (o *PostVpcsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post vpcs params
func (o *PostVpcsParams) WithHTTPClient(client *http.Client) *PostVpcsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post vpcs params
func (o *PostVpcsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post vpcs params
func (o *PostVpcsParams) WithBody(body *models.PostVpcsParamsBody) *PostVpcsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post vpcs params
func (o *PostVpcsParams) SetBody(body *models.PostVpcsParamsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVpcsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
