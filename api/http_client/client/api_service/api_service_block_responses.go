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

// APIServiceBlockReader is a Reader for the APIServiceBlock structure.
type APIServiceBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceBlockDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceBlockOK creates a APIServiceBlockOK with default headers values
func NewAPIServiceBlockOK() *APIServiceBlockOK {
	return &APIServiceBlockOK{}
}

/*APIServiceBlockOK handles this case with default header values.

A successful response.
*/
type APIServiceBlockOK struct {
	Payload *models.APIPbBlockResponse
}

func (o *APIServiceBlockOK) Error() string {
	return fmt.Sprintf("[GET /block/{height}][%d] apiServiceBlockOK  %+v", 200, o.Payload)
}

func (o *APIServiceBlockOK) GetPayload() *models.APIPbBlockResponse {
	return o.Payload
}

func (o *APIServiceBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbBlockResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceBlockDefault creates a APIServiceBlockDefault with default headers values
func NewAPIServiceBlockDefault(code int) *APIServiceBlockDefault {
	return &APIServiceBlockDefault{
		_statusCode: code,
	}
}

/*APIServiceBlockDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceBlockDefault struct {
	_statusCode int

	Payload *models.APIPbErrorBody
}

// Code gets the status code for the Api service block default response
func (o *APIServiceBlockDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceBlockDefault) Error() string {
	return fmt.Sprintf("[GET /block/{height}][%d] ApiService_Block default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceBlockDefault) GetPayload() *models.APIPbErrorBody {
	return o.Payload
}

func (o *APIServiceBlockDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}