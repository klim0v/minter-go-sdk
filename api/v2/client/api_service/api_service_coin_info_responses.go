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

// APIServiceCoinInfoReader is a Reader for the APIServiceCoinInfo structure.
type APIServiceCoinInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceCoinInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceCoinInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAPIServiceCoinInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAPIServiceCoinInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceCoinInfoOK creates a APIServiceCoinInfoOK with default headers values
func NewAPIServiceCoinInfoOK() *APIServiceCoinInfoOK {
	return &APIServiceCoinInfoOK{}
}

/*APIServiceCoinInfoOK handles this case with default header values.

A successful response.
*/
type APIServiceCoinInfoOK struct {
	Payload *models.APIPbCoinInfoResponse
}

func (o *APIServiceCoinInfoOK) Error() string {
	return fmt.Sprintf("[GET /coin_info/{symbol}][%d] apiServiceCoinInfoOK  %+v", 200, o.Payload)
}

func (o *APIServiceCoinInfoOK) GetPayload() *models.APIPbCoinInfoResponse {
	return o.Payload
}

func (o *APIServiceCoinInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbCoinInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceCoinInfoBadRequest creates a APIServiceCoinInfoBadRequest with default headers values
func NewAPIServiceCoinInfoBadRequest() *APIServiceCoinInfoBadRequest {
	return &APIServiceCoinInfoBadRequest{}
}

/*APIServiceCoinInfoBadRequest handles this case with default header values.

An unexpected error response
*/
type APIServiceCoinInfoBadRequest struct {
	Payload *models.APIPbErrorBody
}

func (o *APIServiceCoinInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /coin_info/{symbol}][%d] apiServiceCoinInfoBadRequest  %+v", 400, o.Payload)
}

func (o *APIServiceCoinInfoBadRequest) GetPayload() *models.APIPbErrorBody {
	return o.Payload
}

func (o *APIServiceCoinInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceCoinInfoDefault creates a APIServiceCoinInfoDefault with default headers values
func NewAPIServiceCoinInfoDefault(code int) *APIServiceCoinInfoDefault {
	return &APIServiceCoinInfoDefault{
		_statusCode: code,
	}
}

/*APIServiceCoinInfoDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceCoinInfoDefault struct {
	_statusCode int

	Payload *models.GatewayruntimeError
}

// Code gets the status code for the Api service coin info default response
func (o *APIServiceCoinInfoDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceCoinInfoDefault) Error() string {
	return fmt.Sprintf("[GET /coin_info/{symbol}][%d] ApiService_CoinInfo default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceCoinInfoDefault) GetPayload() *models.GatewayruntimeError {
	return o.Payload
}

func (o *APIServiceCoinInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
