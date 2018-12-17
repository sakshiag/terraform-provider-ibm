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

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// NewPatchSecurityGroupsIDParams creates a new PatchSecurityGroupsIDParams object
// with the default values initialized.
func NewPatchSecurityGroupsIDParams() *PatchSecurityGroupsIDParams {
	var ()
	return &PatchSecurityGroupsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchSecurityGroupsIDParamsWithTimeout creates a new PatchSecurityGroupsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchSecurityGroupsIDParamsWithTimeout(timeout time.Duration) *PatchSecurityGroupsIDParams {
	var ()
	return &PatchSecurityGroupsIDParams{

		timeout: timeout,
	}
}

// NewPatchSecurityGroupsIDParamsWithContext creates a new PatchSecurityGroupsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchSecurityGroupsIDParamsWithContext(ctx context.Context) *PatchSecurityGroupsIDParams {
	var ()
	return &PatchSecurityGroupsIDParams{

		Context: ctx,
	}
}

// NewPatchSecurityGroupsIDParamsWithHTTPClient creates a new PatchSecurityGroupsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchSecurityGroupsIDParamsWithHTTPClient(client *http.Client) *PatchSecurityGroupsIDParams {
	var ()
	return &PatchSecurityGroupsIDParams{
		HTTPClient: client,
	}
}

/*PatchSecurityGroupsIDParams contains all the parameters to send to the API endpoint
for the patch security groups ID operation typically these are written to a http.Request
*/
type PatchSecurityGroupsIDParams struct {

	/*Body*/
	Body *models.PatchSecurityGroupsIDParamsBody
	/*ID
	  The security group identifier

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

// WithTimeout adds the timeout to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) WithTimeout(timeout time.Duration) *PatchSecurityGroupsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) WithContext(ctx context.Context) *PatchSecurityGroupsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) WithHTTPClient(client *http.Client) *PatchSecurityGroupsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) WithBody(body *models.PatchSecurityGroupsIDParamsBody) *PatchSecurityGroupsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) SetBody(body *models.PatchSecurityGroupsIDParamsBody) {
	o.Body = body
}

// WithID adds the id to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) WithID(id string) *PatchSecurityGroupsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) WithVersion(version string) *PatchSecurityGroupsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch security groups ID params
func (o *PatchSecurityGroupsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchSecurityGroupsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
