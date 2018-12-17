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

// NewPatchImagesIDParams creates a new PatchImagesIDParams object
// with the default values initialized.
func NewPatchImagesIDParams() *PatchImagesIDParams {
	var ()
	return &PatchImagesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchImagesIDParamsWithTimeout creates a new PatchImagesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchImagesIDParamsWithTimeout(timeout time.Duration) *PatchImagesIDParams {
	var ()
	return &PatchImagesIDParams{

		timeout: timeout,
	}
}

// NewPatchImagesIDParamsWithContext creates a new PatchImagesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchImagesIDParamsWithContext(ctx context.Context) *PatchImagesIDParams {
	var ()
	return &PatchImagesIDParams{

		Context: ctx,
	}
}

// NewPatchImagesIDParamsWithHTTPClient creates a new PatchImagesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchImagesIDParamsWithHTTPClient(client *http.Client) *PatchImagesIDParams {
	var ()
	return &PatchImagesIDParams{
		HTTPClient: client,
	}
}

/*PatchImagesIDParams contains all the parameters to send to the API endpoint
for the patch images ID operation typically these are written to a http.Request
*/
type PatchImagesIDParams struct {

	/*ID
	  The image identifier

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

// WithTimeout adds the timeout to the patch images ID params
func (o *PatchImagesIDParams) WithTimeout(timeout time.Duration) *PatchImagesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch images ID params
func (o *PatchImagesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch images ID params
func (o *PatchImagesIDParams) WithContext(ctx context.Context) *PatchImagesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch images ID params
func (o *PatchImagesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch images ID params
func (o *PatchImagesIDParams) WithHTTPClient(client *http.Client) *PatchImagesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch images ID params
func (o *PatchImagesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch images ID params
func (o *PatchImagesIDParams) WithID(id string) *PatchImagesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch images ID params
func (o *PatchImagesIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the patch images ID params
func (o *PatchImagesIDParams) WithVersion(version string) *PatchImagesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch images ID params
func (o *PatchImagesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchImagesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
