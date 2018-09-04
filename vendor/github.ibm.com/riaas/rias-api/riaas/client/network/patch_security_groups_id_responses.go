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

// PatchSecurityGroupsIDReader is a Reader for the PatchSecurityGroupsID structure.
type PatchSecurityGroupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSecurityGroupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchSecurityGroupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchSecurityGroupsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSecurityGroupsIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewPatchSecurityGroupsIDOK creates a PatchSecurityGroupsIDOK with default headers values
func NewPatchSecurityGroupsIDOK() *PatchSecurityGroupsIDOK {
	return &PatchSecurityGroupsIDOK{}
}

/*PatchSecurityGroupsIDOK handles this case with default header values.

dummy
*/
type PatchSecurityGroupsIDOK struct {
	Payload *models.SecurityGroup
}

func (o *PatchSecurityGroupsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /security_groups/{id}][%d] patchSecurityGroupsIdOK  %+v", 200, o.Payload)
}

func (o *PatchSecurityGroupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSecurityGroupsIDBadRequest creates a PatchSecurityGroupsIDBadRequest with default headers values
func NewPatchSecurityGroupsIDBadRequest() *PatchSecurityGroupsIDBadRequest {
	return &PatchSecurityGroupsIDBadRequest{}
}

/*PatchSecurityGroupsIDBadRequest handles this case with default header values.

error
*/
type PatchSecurityGroupsIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchSecurityGroupsIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /security_groups/{id}][%d] patchSecurityGroupsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchSecurityGroupsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSecurityGroupsIDDefault creates a PatchSecurityGroupsIDDefault with default headers values
func NewPatchSecurityGroupsIDDefault(code int) *PatchSecurityGroupsIDDefault {
	return &PatchSecurityGroupsIDDefault{
		_statusCode: code,
	}
}

/*PatchSecurityGroupsIDDefault handles this case with default header values.

unexpectederror
*/
type PatchSecurityGroupsIDDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the patch security groups ID default response
func (o *PatchSecurityGroupsIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchSecurityGroupsIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /security_groups/{id}][%d] PatchSecurityGroupsID default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSecurityGroupsIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
