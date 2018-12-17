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

// NewGetInstancesInstanceIDVolumeAttachmentsIDTagsParams creates a new GetInstancesInstanceIDVolumeAttachmentsIDTagsParams object
// with the default values initialized.
func NewGetInstancesInstanceIDVolumeAttachmentsIDTagsParams() *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	var ()
	return &GetInstancesInstanceIDVolumeAttachmentsIDTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesInstanceIDVolumeAttachmentsIDTagsParamsWithTimeout creates a new GetInstancesInstanceIDVolumeAttachmentsIDTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesInstanceIDVolumeAttachmentsIDTagsParamsWithTimeout(timeout time.Duration) *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	var ()
	return &GetInstancesInstanceIDVolumeAttachmentsIDTagsParams{

		timeout: timeout,
	}
}

// NewGetInstancesInstanceIDVolumeAttachmentsIDTagsParamsWithContext creates a new GetInstancesInstanceIDVolumeAttachmentsIDTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesInstanceIDVolumeAttachmentsIDTagsParamsWithContext(ctx context.Context) *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	var ()
	return &GetInstancesInstanceIDVolumeAttachmentsIDTagsParams{

		Context: ctx,
	}
}

// NewGetInstancesInstanceIDVolumeAttachmentsIDTagsParamsWithHTTPClient creates a new GetInstancesInstanceIDVolumeAttachmentsIDTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesInstanceIDVolumeAttachmentsIDTagsParamsWithHTTPClient(client *http.Client) *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	var ()
	return &GetInstancesInstanceIDVolumeAttachmentsIDTagsParams{
		HTTPClient: client,
	}
}

/*GetInstancesInstanceIDVolumeAttachmentsIDTagsParams contains all the parameters to send to the API endpoint
for the get instances instance ID volume attachments ID tags operation typically these are written to a http.Request
*/
type GetInstancesInstanceIDVolumeAttachmentsIDTagsParams struct {

	/*ID
	  The resource identifier

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

// WithTimeout adds the timeout to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) WithTimeout(timeout time.Duration) *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) WithContext(ctx context.Context) *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) WithHTTPClient(client *http.Client) *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) WithID(id string) *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) SetID(id string) {
	o.ID = id
}

// WithInstanceID adds the instanceID to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) WithInstanceID(instanceID string) *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithVersion adds the version to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) WithVersion(version string) *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get instances instance ID volume attachments ID tags params
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesInstanceIDVolumeAttachmentsIDTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
