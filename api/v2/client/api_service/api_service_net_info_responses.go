// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MinterTeam/minter-go-sdk/v2/api/v2/models"
)

// APIServiceNetInfoReader is a Reader for the APIServiceNetInfo structure.
type APIServiceNetInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceNetInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceNetInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAPIServiceNetInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAPIServiceNetInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceNetInfoOK creates a APIServiceNetInfoOK with default headers values
func NewAPIServiceNetInfoOK() *APIServiceNetInfoOK {
	return &APIServiceNetInfoOK{}
}

/*APIServiceNetInfoOK handles this case with default header values.

A successful response.
*/
type APIServiceNetInfoOK struct {
	Payload *models.APIPbNetInfoResponse
}

func (o *APIServiceNetInfoOK) Error() string {
	return fmt.Sprintf("[GET /net_info][%d] apiServiceNetInfoOK  %+v", 200, o.Payload)
}

func (o *APIServiceNetInfoOK) GetPayload() *models.APIPbNetInfoResponse {
	return o.Payload
}

func (o *APIServiceNetInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbNetInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceNetInfoBadRequest creates a APIServiceNetInfoBadRequest with default headers values
func NewAPIServiceNetInfoBadRequest() *APIServiceNetInfoBadRequest {
	return &APIServiceNetInfoBadRequest{}
}

/*APIServiceNetInfoBadRequest handles this case with default header values.

An unexpected error response
*/
type APIServiceNetInfoBadRequest struct {
	Payload *models.APIPbErrorBody
}

func (o *APIServiceNetInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /net_info][%d] apiServiceNetInfoBadRequest  %+v", 400, o.Payload)
}

func (o *APIServiceNetInfoBadRequest) GetPayload() *models.APIPbErrorBody {
	return o.Payload
}

func (o *APIServiceNetInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceNetInfoDefault creates a APIServiceNetInfoDefault with default headers values
func NewAPIServiceNetInfoDefault(code int) *APIServiceNetInfoDefault {
	return &APIServiceNetInfoDefault{
		_statusCode: code,
	}
}

/*APIServiceNetInfoDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceNetInfoDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the Api service net info default response
func (o *APIServiceNetInfoDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceNetInfoDefault) Error() string {
	return fmt.Sprintf("[GET /net_info][%d] ApiService_NetInfo default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceNetInfoDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *APIServiceNetInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
