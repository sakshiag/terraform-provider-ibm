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

// NewDeleteServersServerIDVolumeAttachmentsIDParams creates a new DeleteServersServerIDVolumeAttachmentsIDParams object
// with the default values initialized.
func NewDeleteServersServerIDVolumeAttachmentsIDParams() *DeleteServersServerIDVolumeAttachmentsIDParams {
	var ()
	return &DeleteServersServerIDVolumeAttachmentsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServersServerIDVolumeAttachmentsIDParamsWithTimeout creates a new DeleteServersServerIDVolumeAttachmentsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServersServerIDVolumeAttachmentsIDParamsWithTimeout(timeout time.Duration) *DeleteServersServerIDVolumeAttachmentsIDParams {
	var ()
	return &DeleteServersServerIDVolumeAttachmentsIDParams{

		timeout: timeout,
	}
}

// NewDeleteServersServerIDVolumeAttachmentsIDParamsWithContext creates a new DeleteServersServerIDVolumeAttachmentsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteServersServerIDVolumeAttachmentsIDParamsWithContext(ctx context.Context) *DeleteServersServerIDVolumeAttachmentsIDParams {
	var ()
	return &DeleteServersServerIDVolumeAttachmentsIDParams{

		Context: ctx,
	}
}

// NewDeleteServersServerIDVolumeAttachmentsIDParamsWithHTTPClient creates a new DeleteServersServerIDVolumeAttachmentsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteServersServerIDVolumeAttachmentsIDParamsWithHTTPClient(client *http.Client) *DeleteServersServerIDVolumeAttachmentsIDParams {
	var ()
	return &DeleteServersServerIDVolumeAttachmentsIDParams{
		HTTPClient: client,
	}
}

/*DeleteServersServerIDVolumeAttachmentsIDParams contains all the parameters to send to the API endpoint
for the delete servers server ID volume attachments ID operation typically these are written to a http.Request
*/
type DeleteServersServerIDVolumeAttachmentsIDParams struct {

	/*ID
	  The volume interface identifier

	*/
	ID string
	/*ServerID
	  The server identifier

	*/
	ServerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) WithTimeout(timeout time.Duration) *DeleteServersServerIDVolumeAttachmentsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) WithContext(ctx context.Context) *DeleteServersServerIDVolumeAttachmentsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) WithHTTPClient(client *http.Client) *DeleteServersServerIDVolumeAttachmentsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) WithID(id string) *DeleteServersServerIDVolumeAttachmentsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) SetID(id string) {
	o.ID = id
}

// WithServerID adds the serverID to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) WithServerID(serverID string) *DeleteServersServerIDVolumeAttachmentsIDParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the delete servers server ID volume attachments ID params
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServersServerIDVolumeAttachmentsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param server_id
	if err := r.SetPathParam("server_id", o.ServerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
