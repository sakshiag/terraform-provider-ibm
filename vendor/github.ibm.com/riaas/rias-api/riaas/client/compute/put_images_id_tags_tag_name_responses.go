// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/riaas/rias-api/riaas/models"
)

// PutImagesIDTagsTagNameReader is a Reader for the PutImagesIDTagsTagName structure.
type PutImagesIDTagsTagNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutImagesIDTagsTagNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutImagesIDTagsTagNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutImagesIDTagsTagNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutImagesIDTagsTagNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutImagesIDTagsTagNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPutImagesIDTagsTagNameNoContent creates a PutImagesIDTagsTagNameNoContent with default headers values
func NewPutImagesIDTagsTagNameNoContent() *PutImagesIDTagsTagNameNoContent {
	return &PutImagesIDTagsTagNameNoContent{}
}

/*PutImagesIDTagsTagNameNoContent handles this case with default header values.

error
*/
type PutImagesIDTagsTagNameNoContent struct {
	Payload *models.Riaaserror
}

func (o *PutImagesIDTagsTagNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /images/{id}/tags/{tag_name}][%d] putImagesIdTagsTagNameNoContent  %+v", 204, o.Payload)
}

func (o *PutImagesIDTagsTagNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutImagesIDTagsTagNameBadRequest creates a PutImagesIDTagsTagNameBadRequest with default headers values
func NewPutImagesIDTagsTagNameBadRequest() *PutImagesIDTagsTagNameBadRequest {
	return &PutImagesIDTagsTagNameBadRequest{}
}

/*PutImagesIDTagsTagNameBadRequest handles this case with default header values.

error
*/
type PutImagesIDTagsTagNameBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PutImagesIDTagsTagNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /images/{id}/tags/{tag_name}][%d] putImagesIdTagsTagNameBadRequest  %+v", 400, o.Payload)
}

func (o *PutImagesIDTagsTagNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutImagesIDTagsTagNameNotFound creates a PutImagesIDTagsTagNameNotFound with default headers values
func NewPutImagesIDTagsTagNameNotFound() *PutImagesIDTagsTagNameNotFound {
	return &PutImagesIDTagsTagNameNotFound{}
}

/*PutImagesIDTagsTagNameNotFound handles this case with default header values.

error
*/
type PutImagesIDTagsTagNameNotFound struct {
	Payload *models.Riaaserror
}

func (o *PutImagesIDTagsTagNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /images/{id}/tags/{tag_name}][%d] putImagesIdTagsTagNameNotFound  %+v", 404, o.Payload)
}

func (o *PutImagesIDTagsTagNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutImagesIDTagsTagNameDefault creates a PutImagesIDTagsTagNameDefault with default headers values
func NewPutImagesIDTagsTagNameDefault(code int) *PutImagesIDTagsTagNameDefault {
	return &PutImagesIDTagsTagNameDefault{
		_statusCode: code,
	}
}

/*PutImagesIDTagsTagNameDefault handles this case with default header values.

unexpectederror
*/
type PutImagesIDTagsTagNameDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the put images ID tags tag name default response
func (o *PutImagesIDTagsTagNameDefault) Code() int {
	return o._statusCode
}

func (o *PutImagesIDTagsTagNameDefault) Error() string {
	return fmt.Sprintf("[PUT /images/{id}/tags/{tag_name}][%d] PutImagesIDTagsTagName default  %+v", o._statusCode, o.Payload)
}

func (o *PutImagesIDTagsTagNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
