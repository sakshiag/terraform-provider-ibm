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

// DeleteInstancesIDReader is a Reader for the DeleteInstancesID structure.
type DeleteInstancesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstancesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteInstancesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteInstancesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteInstancesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteInstancesIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewDeleteInstancesIDNoContent creates a DeleteInstancesIDNoContent with default headers values
func NewDeleteInstancesIDNoContent() *DeleteInstancesIDNoContent {
	return &DeleteInstancesIDNoContent{}
}

/*DeleteInstancesIDNoContent handles this case with default header values.

error
*/
type DeleteInstancesIDNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteInstancesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstancesIdNoContent  %+v", 204, o.Payload)
}

func (o *DeleteInstancesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstancesIDBadRequest creates a DeleteInstancesIDBadRequest with default headers values
func NewDeleteInstancesIDBadRequest() *DeleteInstancesIDBadRequest {
	return &DeleteInstancesIDBadRequest{}
}

/*DeleteInstancesIDBadRequest handles this case with default header values.

error
*/
type DeleteInstancesIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *DeleteInstancesIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstancesIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteInstancesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstancesIDNotFound creates a DeleteInstancesIDNotFound with default headers values
func NewDeleteInstancesIDNotFound() *DeleteInstancesIDNotFound {
	return &DeleteInstancesIDNotFound{}
}

/*DeleteInstancesIDNotFound handles this case with default header values.

error
*/
type DeleteInstancesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteInstancesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] deleteInstancesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteInstancesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInstancesIDDefault creates a DeleteInstancesIDDefault with default headers values
func NewDeleteInstancesIDDefault(code int) *DeleteInstancesIDDefault {
	return &DeleteInstancesIDDefault{
		_statusCode: code,
	}
}

/*DeleteInstancesIDDefault handles this case with default header values.

unexpectederror
*/
type DeleteInstancesIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the delete instances ID default response
func (o *DeleteInstancesIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteInstancesIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /instances/{id}][%d] DeleteInstancesID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteInstancesIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
