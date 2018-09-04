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

// PutNetworkAclsIDTagsTagNameReader is a Reader for the PutNetworkAclsIDTagsTagName structure.
type PutNetworkAclsIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworkAclsIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutNetworkAclsIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutNetworkAclsIDTagsTagNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutNetworkAclsIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutNetworkAclsIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPutNetworkAclsIDTagsTagNameNoContent creates a PutNetworkAclsIDTagsTagNameNoContent with default headers values
func NewPutNetworkAclsIDTagsTagNameNoContent() *PutNetworkAclsIDTagsTagNameNoContent {
	return &PutNetworkAclsIDTagsTagNameNoContent{}
}

/*PutNetworkAclsIDTagsTagNameNoContent handles this case with default header values.

error
*/
type PutNetworkAclsIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *PutNetworkAclsIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /network_acls/{id}/tags/{tag_name}][%d] putNetworkAclsIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *PutNetworkAclsIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkAclsIDTagsTagNameBadRequest creates a PutNetworkAclsIDTagsTagNameBadRequest with default headers values
func NewPutNetworkAclsIDTagsTagNameBadRequest() *PutNetworkAclsIDTagsTagNameBadRequest {
	return &PutNetworkAclsIDTagsTagNameBadRequest{}
}

/*PutNetworkAclsIDTagsTagNameBadRequest handles this case with default header values.

error
*/
type PutNetworkAclsIDTagsTagNameBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PutNetworkAclsIDTagsTagNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /network_acls/{id}/tags/{tag_name}][%d] putNetworkAclsIdTagsTagNameBadRequest  %+v", 400, o.Payload)
}

func (o *PutNetworkAclsIDTagsTagNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkAclsIDTagsTagNameNotFound creates a PutNetworkAclsIDTagsTagNameNotFound with default headers values
func NewPutNetworkAclsIDTagsTagNameNotFound() *PutNetworkAclsIDTagsTagNameNotFound {
	return &PutNetworkAclsIDTagsTagNameNotFound{}
}

/*PutNetworkAclsIDTagsTagNameNotFound handles this case with default header values.

error
*/
type PutNetworkAclsIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *PutNetworkAclsIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /network_acls/{id}/tags/{tag_name}][%d] putNetworkAclsIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *PutNetworkAclsIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutNetworkAclsIDTagsTagNameDefault creates a PutNetworkAclsIDTagsTagNameDefault with default headers values
func NewPutNetworkAclsIDTagsTagNameDefault(code int) *PutNetworkAclsIDTagsTagNameDefault {
	return &PutNetworkAclsIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*PutNetworkAclsIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type PutNetworkAclsIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the put network acls ID tags tag name default response
func (o *PutNetworkAclsIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworkAclsIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[PUT /network_acls/{id}/tags/{tag_name}][%d] PutNetworkAclsIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworkAclsIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
