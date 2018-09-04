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

// GetSecurityGroupsIDTagsTagNameReader is a Reader for the GetSecurityGroupsIDTagsTagName structure.
type GetSecurityGroupsIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupsIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewGetSecurityGroupsIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSecurityGroupsIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSecurityGroupsIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetSecurityGroupsIDTagsTagNameNoContent creates a GetSecurityGroupsIDTagsTagNameNoContent with default headers values
func NewGetSecurityGroupsIDTagsTagNameNoContent() *GetSecurityGroupsIDTagsTagNameNoContent {
	return &GetSecurityGroupsIDTagsTagNameNoContent{}
}

/*GetSecurityGroupsIDTagsTagNameNoContent handles this case with default header values.

error
*/
type GetSecurityGroupsIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *GetSecurityGroupsIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[GET /security_groups/{id}/tags/{tag_name}][%d] getSecurityGroupsIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *GetSecurityGroupsIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsIDTagsTagNameNotFound creates a GetSecurityGroupsIDTagsTagNameNotFound with default headers values
func NewGetSecurityGroupsIDTagsTagNameNotFound() *GetSecurityGroupsIDTagsTagNameNotFound {
	return &GetSecurityGroupsIDTagsTagNameNotFound{}
}

/*GetSecurityGroupsIDTagsTagNameNotFound handles this case with default header values.

error
*/
type GetSecurityGroupsIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetSecurityGroupsIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[GET /security_groups/{id}/tags/{tag_name}][%d] getSecurityGroupsIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *GetSecurityGroupsIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsIDTagsTagNameDefault creates a GetSecurityGroupsIDTagsTagNameDefault with default headers values
func NewGetSecurityGroupsIDTagsTagNameDefault(code int) *GetSecurityGroupsIDTagsTagNameDefault {
	return &GetSecurityGroupsIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*GetSecurityGroupsIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type GetSecurityGroupsIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get security groups ID tags tag name default response
func (o *GetSecurityGroupsIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *GetSecurityGroupsIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[GET /security_groups/{id}/tags/{tag_name}][%d] GetSecurityGroupsIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *GetSecurityGroupsIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
