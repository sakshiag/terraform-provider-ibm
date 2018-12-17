// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

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

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// NewPostIkePoliciesParams creates a new PostIkePoliciesParams object
// with the default values initialized.
func NewPostIkePoliciesParams() *PostIkePoliciesParams {
	var ()
	return &PostIkePoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostIkePoliciesParamsWithTimeout creates a new PostIkePoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostIkePoliciesParamsWithTimeout(timeout time.Duration) *PostIkePoliciesParams {
	var ()
	return &PostIkePoliciesParams{

		timeout: timeout,
	}
}

// NewPostIkePoliciesParamsWithContext creates a new PostIkePoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostIkePoliciesParamsWithContext(ctx context.Context) *PostIkePoliciesParams {
	var ()
	return &PostIkePoliciesParams{

		Context: ctx,
	}
}

// NewPostIkePoliciesParamsWithHTTPClient creates a new PostIkePoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostIkePoliciesParamsWithHTTPClient(client *http.Client) *PostIkePoliciesParams {
	var ()
	return &PostIkePoliciesParams{
		HTTPClient: client,
	}
}

/*PostIkePoliciesParams contains all the parameters to send to the API endpoint
for the post ike policies operation typically these are written to a http.Request
*/
type PostIkePoliciesParams struct {

	/*Body*/
	Body *models.IKEPolicyTemplate
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ike policies params
func (o *PostIkePoliciesParams) WithTimeout(timeout time.Duration) *PostIkePoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ike policies params
func (o *PostIkePoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ike policies params
func (o *PostIkePoliciesParams) WithContext(ctx context.Context) *PostIkePoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ike policies params
func (o *PostIkePoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ike policies params
func (o *PostIkePoliciesParams) WithHTTPClient(client *http.Client) *PostIkePoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ike policies params
func (o *PostIkePoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post ike policies params
func (o *PostIkePoliciesParams) WithBody(body *models.IKEPolicyTemplate) *PostIkePoliciesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post ike policies params
func (o *PostIkePoliciesParams) SetBody(body *models.IKEPolicyTemplate) {
	o.Body = body
}

// WithVersion adds the version to the post ike policies params
func (o *PostIkePoliciesParams) WithVersion(version string) *PostIkePoliciesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post ike policies params
func (o *PostIkePoliciesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostIkePoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
