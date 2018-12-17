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

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// NewPatchInstancesInstanceIDVolumeAttachmentsIDParams creates a new PatchInstancesInstanceIDVolumeAttachmentsIDParams object
// with the default values initialized.
func NewPatchInstancesInstanceIDVolumeAttachmentsIDParams() *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	var ()
	return &PatchInstancesInstanceIDVolumeAttachmentsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchInstancesInstanceIDVolumeAttachmentsIDParamsWithTimeout creates a new PatchInstancesInstanceIDVolumeAttachmentsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchInstancesInstanceIDVolumeAttachmentsIDParamsWithTimeout(timeout time.Duration) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	var ()
	return &PatchInstancesInstanceIDVolumeAttachmentsIDParams{

		timeout: timeout,
	}
}

// NewPatchInstancesInstanceIDVolumeAttachmentsIDParamsWithContext creates a new PatchInstancesInstanceIDVolumeAttachmentsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchInstancesInstanceIDVolumeAttachmentsIDParamsWithContext(ctx context.Context) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	var ()
	return &PatchInstancesInstanceIDVolumeAttachmentsIDParams{

		Context: ctx,
	}
}

// NewPatchInstancesInstanceIDVolumeAttachmentsIDParamsWithHTTPClient creates a new PatchInstancesInstanceIDVolumeAttachmentsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchInstancesInstanceIDVolumeAttachmentsIDParamsWithHTTPClient(client *http.Client) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	var ()
	return &PatchInstancesInstanceIDVolumeAttachmentsIDParams{
		HTTPClient: client,
	}
}

/*PatchInstancesInstanceIDVolumeAttachmentsIDParams contains all the parameters to send to the API endpoint
for the patch instances instance ID volume attachments ID operation typically these are written to a http.Request
*/
type PatchInstancesInstanceIDVolumeAttachmentsIDParams struct {

	/*Body*/
	Body *models.PatchInstancesInstanceIDVolumeAttachmentsIDParamsBody
	/*ID
	  The volume interface identifier

	*/
	ID string
	/*InstanceID
	  The instance identifier

	*/
	InstanceID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) WithTimeout(timeout time.Duration) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) WithContext(ctx context.Context) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) WithHTTPClient(client *http.Client) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) WithBody(body *models.PatchInstancesInstanceIDVolumeAttachmentsIDParamsBody) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) SetBody(body *models.PatchInstancesInstanceIDVolumeAttachmentsIDParamsBody) {
	o.Body = body
}

// WithID adds the id to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) WithID(id string) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) SetID(id string) {
	o.ID = id
}

// WithInstanceID adds the instanceID to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) WithInstanceID(instanceID string) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithVersion adds the version to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) WithVersion(version string) *PatchInstancesInstanceIDVolumeAttachmentsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch instances instance ID volume attachments ID params
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchInstancesInstanceIDVolumeAttachmentsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
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
