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

// NewDeleteServersServerIDNetworkInterfacesIDTagsTagNameParams creates a new DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams object
// with the default values initialized.
func NewDeleteServersServerIDNetworkInterfacesIDTagsTagNameParams() *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	var ()
	return &DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServersServerIDNetworkInterfacesIDTagsTagNameParamsWithTimeout creates a new DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServersServerIDNetworkInterfacesIDTagsTagNameParamsWithTimeout(timeout time.Duration) *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	var ()
	return &DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewDeleteServersServerIDNetworkInterfacesIDTagsTagNameParamsWithContext creates a new DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteServersServerIDNetworkInterfacesIDTagsTagNameParamsWithContext(ctx context.Context) *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	var ()
	return &DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewDeleteServersServerIDNetworkInterfacesIDTagsTagNameParamsWithHTTPClient creates a new DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteServersServerIDNetworkInterfacesIDTagsTagNameParamsWithHTTPClient(client *http.Client) *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	var ()
	return &DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the delete servers server ID network interfaces ID tags tag name operation typically these are written to a http.Request
*/
type DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams struct {

	/*ID
	  The resource identifier

	*/
	ID string
	/*ServerID
	  The server identifier

	*/
	ServerID string
	/*TagName
	  The name of the tag

	*/
	TagName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) WithTimeout(timeout time.Duration) *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) WithContext(ctx context.Context) *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) WithHTTPClient(client *http.Client) *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) WithID(id string) *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithServerID adds the serverID to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) WithServerID(serverID string) *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) SetServerID(serverID string) {
	o.ServerID = serverID
}

// WithTagName adds the tagName to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) WithTagName(tagName string) *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the delete servers server ID network interfaces ID tags tag name params
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServersServerIDNetworkInterfacesIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tag_name
	if err := r.SetPathParam("tag_name", o.TagName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
