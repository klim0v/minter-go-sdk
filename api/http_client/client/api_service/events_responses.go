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

// EventsReader is a Reader for the Events structure.
type EventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEventsOK creates a EventsOK with default headers values
func NewEventsOK() *EventsOK {
	return &EventsOK{}
}

/*EventsOK handles this case with default header values.

A successful response.
*/
type EventsOK struct {
	Payload *models.EventsResponse
}

func (o *EventsOK) Error() string {
	return fmt.Sprintf("[GET /events/{height}][%d] eventsOK  %+v", 200, o.Payload)
}

func (o *EventsOK) GetPayload() *models.EventsResponse {
	return o.Payload
}

func (o *EventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventsDefault creates a EventsDefault with default headers values
func NewEventsDefault(code int) *EventsDefault {
	return &EventsDefault{
		_statusCode: code,
	}
}

/*EventsDefault handles this case with default header values.

An unexpected error response
*/
type EventsDefault struct {
	_statusCode int

	Payload *models.ErrorBody
}

// Code gets the status code for the events default response
func (o *EventsDefault) Code() int {
	return o._statusCode
}

func (o *EventsDefault) Error() string {
	return fmt.Sprintf("[GET /events/{height}][%d] Events default  %+v", o._statusCode, o.Payload)
}

func (o *EventsDefault) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *EventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
