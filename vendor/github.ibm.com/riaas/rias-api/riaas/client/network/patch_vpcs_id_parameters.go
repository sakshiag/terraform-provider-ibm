// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// NewPatchVpcsIDParams creates a new PatchVpcsIDParams object
// with the default values initialized.
func NewPatchVpcsIDParams() *PatchVpcsIDParams {
	var ()
	return &PatchVpcsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchVpcsIDParamsWithTimeout creates a new PatchVpcsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchVpcsIDParamsWithTimeout(timeout time.Duration) *PatchVpcsIDParams {
	var ()
	return &PatchVpcsIDParams{

		timeout: timeout,
	}
}

// NewPatchVpcsIDParamsWithContext creates a new PatchVpcsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchVpcsIDParamsWithContext(ctx context.Context) *PatchVpcsIDParams {
	var ()
	return &PatchVpcsIDParams{

		Context: ctx,
	}
}

// NewPatchVpcsIDParamsWithHTTPClient creates a new PatchVpcsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchVpcsIDParamsWithHTTPClient(client *http.Client) *PatchVpcsIDParams {
	var ()
	return &PatchVpcsIDParams{
		HTTPClient: client,
	}
}

/*PatchVpcsIDParams contains all the parameters to send to the API endpoint
for the patch vpcs ID operation typically these are written to a http.Request
*/
type PatchVpcsIDParams struct {

	/*Body*/
	Body *models.PatchVpcsIDParamsBody
	/*ID
	  The VPC identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch vpcs ID params
func (o *PatchVpcsIDParams) WithTimeout(timeout time.Duration) *PatchVpcsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch vpcs ID params
func (o *PatchVpcsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch vpcs ID params
func (o *PatchVpcsIDParams) WithContext(ctx context.Context) *PatchVpcsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch vpcs ID params
func (o *PatchVpcsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch vpcs ID params
func (o *PatchVpcsIDParams) WithHTTPClient(client *http.Client) *PatchVpcsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch vpcs ID params
func (o *PatchVpcsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch vpcs ID params
func (o *PatchVpcsIDParams) WithBody(body *models.PatchVpcsIDParamsBody) *PatchVpcsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch vpcs ID params
func (o *PatchVpcsIDParams) SetBody(body *models.PatchVpcsIDParamsBody) {
	o.Body = body
}

// WithID adds the id to the patch vpcs ID params
func (o *PatchVpcsIDParams) WithID(id string) *PatchVpcsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch vpcs ID params
func (o *PatchVpcsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the patch vpcs ID params
func (o *PatchVpcsIDParams) WithVersion(version string) *PatchVpcsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch vpcs ID params
func (o *PatchVpcsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchVpcsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
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
