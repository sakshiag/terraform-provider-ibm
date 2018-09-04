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

// DeleteFloatingIpsIDTagsTagNameReader is a Reader for the DeleteFloatingIpsIDTagsTagName structure.
type DeleteFloatingIpsIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFloatingIpsIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFloatingIpsIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteFloatingIpsIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteFloatingIpsIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewDeleteFloatingIpsIDTagsTagNameNoContent creates a DeleteFloatingIpsIDTagsTagNameNoContent with default headers values
func NewDeleteFloatingIpsIDTagsTagNameNoContent() *DeleteFloatingIpsIDTagsTagNameNoContent {
	return &DeleteFloatingIpsIDTagsTagNameNoContent{}
}

/*DeleteFloatingIpsIDTagsTagNameNoContent handles this case with default header values.

error
*/
type DeleteFloatingIpsIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteFloatingIpsIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /floating_ips/{id}/tags/{tag_name}][%d] deleteFloatingIpsIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *DeleteFloatingIpsIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFloatingIpsIDTagsTagNameNotFound creates a DeleteFloatingIpsIDTagsTagNameNotFound with default headers values
func NewDeleteFloatingIpsIDTagsTagNameNotFound() *DeleteFloatingIpsIDTagsTagNameNotFound {
	return &DeleteFloatingIpsIDTagsTagNameNotFound{}
}

/*DeleteFloatingIpsIDTagsTagNameNotFound handles this case with default header values.

error
*/
type DeleteFloatingIpsIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteFloatingIpsIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /floating_ips/{id}/tags/{tag_name}][%d] deleteFloatingIpsIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFloatingIpsIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFloatingIpsIDTagsTagNameDefault creates a DeleteFloatingIpsIDTagsTagNameDefault with default headers values
func NewDeleteFloatingIpsIDTagsTagNameDefault(code int) *DeleteFloatingIpsIDTagsTagNameDefault {
	return &DeleteFloatingIpsIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*DeleteFloatingIpsIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type DeleteFloatingIpsIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the delete floating ips ID tags tag name default response
func (o *DeleteFloatingIpsIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFloatingIpsIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /floating_ips/{id}/tags/{tag_name}][%d] DeleteFloatingIpsIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteFloatingIpsIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
