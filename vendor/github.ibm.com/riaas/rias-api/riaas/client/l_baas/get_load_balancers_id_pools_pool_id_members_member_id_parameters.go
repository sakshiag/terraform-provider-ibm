// Code generated by go-swagger; DO NOT EDIT.

package l_baas

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

// NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParams creates a new GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams object
// with the default values initialized.
func NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParams() *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	var ()
	return &GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithTimeout creates a new GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithTimeout(timeout time.Duration) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	var ()
	return &GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams{

		timeout: timeout,
	}
}

// NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithContext creates a new GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithContext(ctx context.Context) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	var ()
	return &GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams{

		Context: ctx,
	}
}

// NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithHTTPClient creates a new GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithHTTPClient(client *http.Client) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	var ()
	return &GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams{
		HTTPClient: client,
	}
}

/*GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams contains all the parameters to send to the API endpoint
for the get load balancers ID pools pool ID members member ID operation typically these are written to a http.Request
*/
type GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams struct {

	/*ID
	  The load balancer identifier

	*/
	ID string
	/*MemberID
	  The member identifier

	*/
	MemberID string
	/*PoolID
	  The pool identifier

	*/
	PoolID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithTimeout(timeout time.Duration) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithContext(ctx context.Context) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithHTTPClient(client *http.Client) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithID(id string) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetID(id string) {
	o.ID = id
}

// WithMemberID adds the memberID to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithMemberID(memberID string) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetMemberID(memberID string) {
	o.MemberID = memberID
}

// WithPoolID adds the poolID to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithPoolID(poolID string) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithVersion adds the version to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WithVersion(version string) *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get load balancers ID pools pool ID members member ID params
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoadBalancersIDPoolsPoolIDMembersMemberIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param member_id
	if err := r.SetPathParam("member_id", o.MemberID); err != nil {
		return err
	}

	// path param pool_id
	if err := r.SetPathParam("pool_id", o.PoolID); err != nil {
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
