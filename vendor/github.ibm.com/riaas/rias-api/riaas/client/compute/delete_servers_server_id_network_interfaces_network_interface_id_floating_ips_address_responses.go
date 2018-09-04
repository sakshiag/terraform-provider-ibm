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

// DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressReader is a Reader for the DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress structure.
type DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		
		return nil, result
	}
}

// NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent creates a DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent with default headers values
func NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent() *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent {
	return &DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent{}
}

/*DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent handles this case with default header values.

The floating IP was disassociated successfully.
*/
type DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent struct {
}

func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent) Error() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}/network_interfaces/{network_interface_id}/floating_ips/{address}][%d] deleteServersServerIdNetworkInterfacesNetworkInterfaceIdFloatingIpsAddressNoContent ", 204)
}

func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest creates a DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest with default headers values
func NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest() *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest {
	return &DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest{}
}

/*DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest handles this case with default header values.

The specified floating IP could not be disassociated.
*/
type DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest struct {
	Payload *models.Riaaserror
}

func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}/network_interfaces/{network_interface_id}/floating_ips/{address}][%d] deleteServersServerIdNetworkInterfacesNetworkInterfaceIdFloatingIpsAddressBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound creates a DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound with default headers values
func NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound() *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound {
	return &DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound{}
}

/*DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound handles this case with default header values.

The specified floating IP address is not associated with the network interface with the specified identifier
*/
type DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound) Error() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}/network_interfaces/{network_interface_id}/floating_ips/{address}][%d] deleteServersServerIdNetworkInterfacesNetworkInterfaceIdFloatingIpsAddressNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault creates a DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault with default headers values
func NewDeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault(code int) *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault {
	return &DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault{
		_statusCode: code,
	}
}

/*DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault handles this case with default header values.

unexpectederror
*/
type DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault struct {
	_statusCode int

	Payload *models.Riaaserror
}

// Code gets the status code for the delete servers server ID network interfaces network interface ID floating ips address default response
func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault) Error() string {
	return fmt.Sprintf("[DELETE /servers/{server_id}/network_interfaces/{network_interface_id}/floating_ips/{address}][%d] DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServersServerIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
