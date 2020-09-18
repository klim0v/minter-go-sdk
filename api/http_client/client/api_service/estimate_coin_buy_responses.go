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

// EstimateCoinBuyReader is a Reader for the EstimateCoinBuy structure.
type EstimateCoinBuyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EstimateCoinBuyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEstimateCoinBuyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEstimateCoinBuyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEstimateCoinBuyOK creates a EstimateCoinBuyOK with default headers values
func NewEstimateCoinBuyOK() *EstimateCoinBuyOK {
	return &EstimateCoinBuyOK{}
}

/*EstimateCoinBuyOK handles this case with default header values.

A successful response.
*/
type EstimateCoinBuyOK struct {
	Payload *models.EstimateCoinBuyResponse
}

func (o *EstimateCoinBuyOK) Error() string {
	return fmt.Sprintf("[GET /estimate_coin_buy][%d] estimateCoinBuyOK  %+v", 200, o.Payload)
}

func (o *EstimateCoinBuyOK) GetPayload() *models.EstimateCoinBuyResponse {
	return o.Payload
}

func (o *EstimateCoinBuyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EstimateCoinBuyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEstimateCoinBuyDefault creates a EstimateCoinBuyDefault with default headers values
func NewEstimateCoinBuyDefault(code int) *EstimateCoinBuyDefault {
	return &EstimateCoinBuyDefault{
		_statusCode: code,
	}
}

/*EstimateCoinBuyDefault handles this case with default header values.

An unexpected error response
*/
type EstimateCoinBuyDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the estimate coin buy default response
func (o *EstimateCoinBuyDefault) Code() int {
	return o._statusCode
}

func (o *EstimateCoinBuyDefault) Error() string {
	return fmt.Sprintf("[GET /estimate_coin_buy][%d] EstimateCoinBuy default  %+v", o._statusCode, o.Payload)
}

func (o *EstimateCoinBuyDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *EstimateCoinBuyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}