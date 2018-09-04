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

// GetVpcsIDTagsReader is a Reader for the GetVpcsIDTags structure.
type GetVpcsIDTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpcsIDTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpcsIDTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetVpcsIDTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetVpcsIDTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewGetVpcsIDTagsOK creates a GetVpcsIDTagsOK with default headers values
func NewGetVpcsIDTagsOK() *GetVpcsIDTagsOK {
	return &GetVpcsIDTagsOK{}
}

/*GetVpcsIDTagsOK handles this case with default header values.

dummy
*/
type GetVpcsIDTagsOK struct {
	Payload *models.GetVpcsIDTagsOKBody
}

func (o *GetVpcsIDTagsOK) Error() string {
	return fmt.Sprintf("[GET /vpcs/{id}/tags][%d] getVpcsIdTagsOK  %+v", 200, o.Payload)
}

func (o *GetVpcsIDTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVpcsIDTagsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpcsIDTagsNotFound creates a GetVpcsIDTagsNotFound with default headers values
func NewGetVpcsIDTagsNotFound() *GetVpcsIDTagsNotFound {
	return &GetVpcsIDTagsNotFound{}
}

/*GetVpcsIDTagsNotFound handles this case with default header values.

error
*/
type GetVpcsIDTagsNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetVpcsIDTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /vpcs/{id}/tags][%d] getVpcsIdTagsNotFound  %+v", 404, o.Payload)
}

func (o *GetVpcsIDTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpcsIDTagsDefault creates a GetVpcsIDTagsDefault with default headers values
func NewGetVpcsIDTagsDefault(code int) *GetVpcsIDTagsDefault {
	return &GetVpcsIDTagsDefault{
		_statusCode: code,
	}
}

/*GetVpcsIDTagsDefault handles this case with default header values.

unexpectederror
*/
type GetVpcsIDTagsDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the get vpcs ID tags default response
func (o *GetVpcsIDTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetVpcsIDTagsDefault) Error() string {
	return fmt.Sprintf("[GET /vpcs/{id}/tags][%d] GetVpcsIDTags default  %+v", o._statusCode, o.Payload)
}

func (o *GetVpcsIDTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
