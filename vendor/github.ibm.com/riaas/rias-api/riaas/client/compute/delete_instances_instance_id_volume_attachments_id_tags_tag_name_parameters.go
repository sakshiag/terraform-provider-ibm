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

// NewDeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams creates a new DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams object
// with the default values initialized.
func NewDeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams() *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	var ()
	return &DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParamsWithTimeout creates a new DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParamsWithTimeout(timeout time.Duration) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	var ()
	return &DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewDeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParamsWithContext creates a new DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParamsWithContext(ctx context.Context) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	var ()
	return &DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewDeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParamsWithHTTPClient creates a new DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParamsWithHTTPClient(client *http.Client) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	var ()
	return &DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the delete instances instance ID volume attachments ID tags tag name operation typically these are written to a http.Request
*/
type DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams struct {

	/*ID
	  The resource identifier

	*/
	ID string
	/*InstanceID
	  The instance identifier

	*/
	InstanceID string
	/*TagName
	  The name of the tag

	*/
	TagName string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) WithTimeout(timeout time.Duration) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) WithContext(ctx context.Context) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) WithHTTPClient(client *http.Client) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) WithID(id string) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithInstanceID adds the instanceID to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) WithInstanceID(instanceID string) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithTagName adds the tagName to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) WithTagName(tagName string) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) WithVersion(version string) *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete instances instance ID volume attachments ID tags tag name params
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesInstanceIDVolumeAttachmentsIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tag_name
	if err := r.SetPathParam("tag_name", o.TagName); err != nil {
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
