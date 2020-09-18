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

// APIServiceCandidatesReader is a Reader for the APIServiceCandidates structure.
type APIServiceCandidatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceCandidatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceCandidatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceCandidatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceCandidatesOK creates a APIServiceCandidatesOK with default headers values
func NewAPIServiceCandidatesOK() *APIServiceCandidatesOK {
	return &APIServiceCandidatesOK{}
}

/*APIServiceCandidatesOK handles this case with default header values.

A successful response.
*/
type APIServiceCandidatesOK struct {
	Payload *models.APIPbCandidatesResponse
}

func (o *APIServiceCandidatesOK) Error() string {
	return fmt.Sprintf("[GET /candidates][%d] apiServiceCandidatesOK  %+v", 200, o.Payload)
}

func (o *APIServiceCandidatesOK) GetPayload() *models.APIPbCandidatesResponse {
	return o.Payload
}

func (o *APIServiceCandidatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbCandidatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceCandidatesDefault creates a APIServiceCandidatesDefault with default headers values
func NewAPIServiceCandidatesDefault(code int) *APIServiceCandidatesDefault {
	return &APIServiceCandidatesDefault{
		_statusCode: code,
	}
}

/*APIServiceCandidatesDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceCandidatesDefault struct {
	_statusCode int

	Payload *models.APIPbErrorBody
}

// Code gets the status code for the Api service candidates default response
func (o *APIServiceCandidatesDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceCandidatesDefault) Error() string {
	return fmt.Sprintf("[GET /candidates][%d] ApiService_Candidates default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceCandidatesDefault) GetPayload() *models.APIPbErrorBody {
	return o.Payload
}

func (o *APIServiceCandidatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}