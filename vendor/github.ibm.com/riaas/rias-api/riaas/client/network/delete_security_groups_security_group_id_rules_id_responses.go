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

// DeleteSecurityGroupsSecurityGroupIDRulesIDReader is a Reader for the DeleteSecurityGroupsSecurityGroupIDRulesID structure.
type DeleteSecurityGroupsSecurityGroupIDRulesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSecurityGroupsSecurityGroupIDRulesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteSecurityGroupsSecurityGroupIDRulesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSecurityGroupsSecurityGroupIDRulesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewDeleteSecurityGroupsSecurityGroupIDRulesIDNoContent creates a DeleteSecurityGroupsSecurityGroupIDRulesIDNoContent with default headers values
func NewDeleteSecurityGroupsSecurityGroupIDRulesIDNoContent() *DeleteSecurityGroupsSecurityGroupIDRulesIDNoContent {
	return &DeleteSecurityGroupsSecurityGroupIDRulesIDNoContent{}
}

/*DeleteSecurityGroupsSecurityGroupIDRulesIDNoContent handles this case with default header values.

error
*/
type DeleteSecurityGroupsSecurityGroupIDRulesIDNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /security_groups/{security_group_id}/rules/{id}][%d] deleteSecurityGroupsSecurityGroupIdRulesIdNoContent  %+v", 204, o.Payload)
}

func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecurityGroupsSecurityGroupIDRulesIDNotFound creates a DeleteSecurityGroupsSecurityGroupIDRulesIDNotFound with default headers values
func NewDeleteSecurityGroupsSecurityGroupIDRulesIDNotFound() *DeleteSecurityGroupsSecurityGroupIDRulesIDNotFound {
	return &DeleteSecurityGroupsSecurityGroupIDRulesIDNotFound{}
}

/*DeleteSecurityGroupsSecurityGroupIDRulesIDNotFound handles this case with default header values.

error
*/
type DeleteSecurityGroupsSecurityGroupIDRulesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /security_groups/{security_group_id}/rules/{id}][%d] deleteSecurityGroupsSecurityGroupIdRulesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecurityGroupsSecurityGroupIDRulesIDDefault creates a DeleteSecurityGroupsSecurityGroupIDRulesIDDefault with default headers values
func NewDeleteSecurityGroupsSecurityGroupIDRulesIDDefault(code int) *DeleteSecurityGroupsSecurityGroupIDRulesIDDefault {
	return &DeleteSecurityGroupsSecurityGroupIDRulesIDDefault{
		_statusCode: code,
	}
}

/*DeleteSecurityGroupsSecurityGroupIDRulesIDDefault handles this case with default header values.

unexpectederror
*/
type DeleteSecurityGroupsSecurityGroupIDRulesIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the delete security groups security group ID rules ID default response
func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /security_groups/{security_group_id}/rules/{id}][%d] DeleteSecurityGroupsSecurityGroupIDRulesID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSecurityGroupsSecurityGroupIDRulesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
