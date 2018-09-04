// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.ibm.com/riaas/rias-api/riaas/models"
)

// DeleteNetworkAclsNetworkACLIDRulesIDReader is a Reader for the DeleteNetworkAclsNetworkACLIDRulesID structure.
type DeleteNetworkAclsNetworkACLIDRulesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkAclsNetworkACLIDRulesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNetworkAclsNetworkACLIDRulesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteNetworkAclsNetworkACLIDRulesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteNetworkAclsNetworkACLIDRulesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewDeleteNetworkAclsNetworkACLIDRulesIDNoContent creates a DeleteNetworkAclsNetworkACLIDRulesIDNoContent with default headers values
func NewDeleteNetworkAclsNetworkACLIDRulesIDNoContent() *DeleteNetworkAclsNetworkACLIDRulesIDNoContent {
	return &DeleteNetworkAclsNetworkACLIDRulesIDNoContent{}
}

/*DeleteNetworkAclsNetworkACLIDRulesIDNoContent handles this case with default header values.

error
*/
type DeleteNetworkAclsNetworkACLIDRulesIDNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteNetworkAclsNetworkACLIDRulesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /network_acls/{network_acl_id}/rules/{id}][%d] deleteNetworkAclsNetworkAclIdRulesIdNoContent  %+v", 204, o.Payload)
}

func (o *DeleteNetworkAclsNetworkACLIDRulesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkAclsNetworkACLIDRulesIDNotFound creates a DeleteNetworkAclsNetworkACLIDRulesIDNotFound with default headers values
func NewDeleteNetworkAclsNetworkACLIDRulesIDNotFound() *DeleteNetworkAclsNetworkACLIDRulesIDNotFound {
	return &DeleteNetworkAclsNetworkACLIDRulesIDNotFound{}
}

/*DeleteNetworkAclsNetworkACLIDRulesIDNotFound handles this case with default header values.

error
*/
type DeleteNetworkAclsNetworkACLIDRulesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteNetworkAclsNetworkACLIDRulesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /network_acls/{network_acl_id}/rules/{id}][%d] deleteNetworkAclsNetworkAclIdRulesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNetworkAclsNetworkACLIDRulesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkAclsNetworkACLIDRulesIDDefault creates a DeleteNetworkAclsNetworkACLIDRulesIDDefault with default headers values
func NewDeleteNetworkAclsNetworkACLIDRulesIDDefault(code int) *DeleteNetworkAclsNetworkACLIDRulesIDDefault {
	return &DeleteNetworkAclsNetworkACLIDRulesIDDefault{
		_statusCode: code,
	}
}

/*DeleteNetworkAclsNetworkACLIDRulesIDDefault handles this case with default header values.

unexpectederror
*/
type DeleteNetworkAclsNetworkACLIDRulesIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the delete network acls network ACL ID rules ID default response
func (o *DeleteNetworkAclsNetworkACLIDRulesIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworkAclsNetworkACLIDRulesIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /network_acls/{network_acl_id}/rules/{id}][%d] DeleteNetworkAclsNetworkACLIDRulesID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNetworkAclsNetworkACLIDRulesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
