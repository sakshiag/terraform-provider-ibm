// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewPostVolumesVolumeIDSnapshotsParams creates a new PostVolumesVolumeIDSnapshotsParams object
// with the default values initialized.
func NewPostVolumesVolumeIDSnapshotsParams() *PostVolumesVolumeIDSnapshotsParams {
	var ()
	return &PostVolumesVolumeIDSnapshotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostVolumesVolumeIDSnapshotsParamsWithTimeout creates a new PostVolumesVolumeIDSnapshotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostVolumesVolumeIDSnapshotsParamsWithTimeout(timeout time.Duration) *PostVolumesVolumeIDSnapshotsParams {
	var ()
	return &PostVolumesVolumeIDSnapshotsParams{

		timeout: timeout,
	}
}

// NewPostVolumesVolumeIDSnapshotsParamsWithContext creates a new PostVolumesVolumeIDSnapshotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostVolumesVolumeIDSnapshotsParamsWithContext(ctx context.Context) *PostVolumesVolumeIDSnapshotsParams {
	var ()
	return &PostVolumesVolumeIDSnapshotsParams{

		Context: ctx,
	}
}

// NewPostVolumesVolumeIDSnapshotsParamsWithHTTPClient creates a new PostVolumesVolumeIDSnapshotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostVolumesVolumeIDSnapshotsParamsWithHTTPClient(client *http.Client) *PostVolumesVolumeIDSnapshotsParams {
	var ()
	return &PostVolumesVolumeIDSnapshotsParams{
		HTTPClient: client,
	}
}

/*PostVolumesVolumeIDSnapshotsParams contains all the parameters to send to the API endpoint
for the post volumes volume ID snapshots operation typically these are written to a http.Request
*/
type PostVolumesVolumeIDSnapshotsParams struct {

	/*Body*/
	Body *models.PostVolumesVolumeIDSnapshotsParamsBody
	/*VolumeID
	  The volume identifier

	*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) WithTimeout(timeout time.Duration) *PostVolumesVolumeIDSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) WithContext(ctx context.Context) *PostVolumesVolumeIDSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) WithHTTPClient(client *http.Client) *PostVolumesVolumeIDSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) WithBody(body *models.PostVolumesVolumeIDSnapshotsParamsBody) *PostVolumesVolumeIDSnapshotsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) SetBody(body *models.PostVolumesVolumeIDSnapshotsParamsBody) {
	o.Body = body
}

// WithVolumeID adds the volumeID to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) WithVolumeID(volumeID string) *PostVolumesVolumeIDSnapshotsParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the post volumes volume ID snapshots params
func (o *PostVolumesVolumeIDSnapshotsParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *PostVolumesVolumeIDSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param volume_id
	if err := r.SetPathParam("volume_id", o.VolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
