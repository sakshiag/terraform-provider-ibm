// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.ibm.com/riaas/rias-api/riaas/models"
)

// GetInstancesInstanceIDVolumeAttachmentsReader is a Reader for the GetInstancesInstanceIDVolumeAttachments structure.
type GetInstancesInstanceIDVolumeAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesInstanceIDVolumeAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesInstanceIDVolumeAttachmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInstancesInstanceIDVolumeAttachmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetInstancesInstanceIDVolumeAttachmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetInstancesInstanceIDVolumeAttachmentsOK creates a GetInstancesInstanceIDVolumeAttachmentsOK with default headers values
func NewGetInstancesInstanceIDVolumeAttachmentsOK() *GetInstancesInstanceIDVolumeAttachmentsOK {
	return &GetInstancesInstanceIDVolumeAttachmentsOK{}
}

/*GetInstancesInstanceIDVolumeAttachmentsOK handles this case with default header values.

dummy
*/
type GetInstancesInstanceIDVolumeAttachmentsOK struct {
	Payload *models.GetInstancesInstanceIDVolumeAttachmentsOKBody
}

func (o *GetInstancesInstanceIDVolumeAttachmentsOK) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/volume_attachments][%d] getInstancesInstanceIdVolumeAttachmentsOK  %+v", 200, o.Payload)
}

func (o *GetInstancesInstanceIDVolumeAttachmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetInstancesInstanceIDVolumeAttachmentsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesInstanceIDVolumeAttachmentsNotFound creates a GetInstancesInstanceIDVolumeAttachmentsNotFound with default headers values
func NewGetInstancesInstanceIDVolumeAttachmentsNotFound() *GetInstancesInstanceIDVolumeAttachmentsNotFound {
	return &GetInstancesInstanceIDVolumeAttachmentsNotFound{}
}

/*GetInstancesInstanceIDVolumeAttachmentsNotFound handles this case with default header values.

error
*/
type GetInstancesInstanceIDVolumeAttachmentsNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetInstancesInstanceIDVolumeAttachmentsNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/volume_attachments][%d] getInstancesInstanceIdVolumeAttachmentsNotFound  %+v", 404, o.Payload)
}

func (o *GetInstancesInstanceIDVolumeAttachmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesInstanceIDVolumeAttachmentsDefault creates a GetInstancesInstanceIDVolumeAttachmentsDefault with default headers values
func NewGetInstancesInstanceIDVolumeAttachmentsDefault(code int) *GetInstancesInstanceIDVolumeAttachmentsDefault {
	return &GetInstancesInstanceIDVolumeAttachmentsDefault{
		_statusCode: code,
	}
}

/*GetInstancesInstanceIDVolumeAttachmentsDefault handles this case with default header values.

unexpectederror
*/
type GetInstancesInstanceIDVolumeAttachmentsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get instances instance ID volume attachments default response
func (o *GetInstancesInstanceIDVolumeAttachmentsDefault) Code() int {
	return o._statusCode
}

func (o *GetInstancesInstanceIDVolumeAttachmentsDefault) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/volume_attachments][%d] GetInstancesInstanceIDVolumeAttachments default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstancesInstanceIDVolumeAttachmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
