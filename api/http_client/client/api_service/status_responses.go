// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
)

// StatusReader is a Reader for the Status structure.
type StatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatusOK creates a StatusOK with default headers values
func NewStatusOK() *StatusOK {
	return &StatusOK{}
}

/*StatusOK handles this case with default header values.

A successful response.
*/
type StatusOK struct {
	Payload *models.StatusResponse
}

func (o *StatusOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] statusOK  %+v", 200, o.Payload)
}

func (o *StatusOK) GetPayload() *models.StatusResponse {
	return o.Payload
}

func (o *StatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatusDefault creates a StatusDefault with default headers values
func NewStatusDefault(code int) *StatusDefault {
	return &StatusDefault{
		_statusCode: code,
	}
}

/*StatusDefault handles this case with default header values.

An unexpected error response
*/
type StatusDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the status default response
func (o *StatusDefault) Code() int {
	return o._statusCode
}

func (o *StatusDefault) Error() string {
	return fmt.Sprintf("[GET /status][%d] Status default  %+v", o._statusCode, o.Payload)
}

func (o *StatusDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *StatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
