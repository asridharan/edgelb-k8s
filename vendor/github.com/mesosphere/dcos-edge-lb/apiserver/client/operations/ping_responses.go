// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PingReader is a Reader for the Ping structure.
type PingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPingOK creates a PingOK with default headers values
func NewPingOK() *PingOK {
	return &PingOK{}
}

/*PingOK handles this case with default header values.

Pong.
*/
type PingOK struct {
	Payload string
}

func (o *PingOK) Error() string {
	return fmt.Sprintf("[GET /ping][%d] pingOK  %+v", 200, o.Payload)
}

func (o *PingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingDefault creates a PingDefault with default headers values
func NewPingDefault(code int) *PingDefault {
	return &PingDefault{
		_statusCode: code,
	}
}

/*PingDefault handles this case with default header values.

Unexpected error.
*/
type PingDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the ping default response
func (o *PingDefault) Code() int {
	return o._statusCode
}

func (o *PingDefault) Error() string {
	return fmt.Sprintf("[GET /ping][%d] Ping default  %+v", o._statusCode, o.Payload)
}

func (o *PingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}