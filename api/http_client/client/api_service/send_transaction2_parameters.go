// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/MinterTeam/minter-go-sdk/v2/api/http_client/models"
)

// NewSendTransaction2Params creates a new SendTransaction2Params object
// with the default values initialized.
func NewSendTransaction2Params() *SendTransaction2Params {
	var ()
	return &SendTransaction2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendTransaction2ParamsWithTimeout creates a new SendTransaction2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendTransaction2ParamsWithTimeout(timeout time.Duration) *SendTransaction2Params {
	var ()
	return &SendTransaction2Params{

		timeout: timeout,
	}
}

// NewSendTransaction2ParamsWithContext creates a new SendTransaction2Params object
// with the default values initialized, and the ability to set a context for a request
func NewSendTransaction2ParamsWithContext(ctx context.Context) *SendTransaction2Params {
	var ()
	return &SendTransaction2Params{

		Context: ctx,
	}
}

// NewSendTransaction2ParamsWithHTTPClient creates a new SendTransaction2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendTransaction2ParamsWithHTTPClient(client *http.Client) *SendTransaction2Params {
	var ()
	return &SendTransaction2Params{
		HTTPClient: client,
	}
}

/*SendTransaction2Params contains all the parameters to send to the API endpoint
for the send transaction2 operation typically these are written to a http.Request
*/
type SendTransaction2Params struct {

	/*Body*/
	Body *models.SendTransactionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send transaction2 params
func (o *SendTransaction2Params) WithTimeout(timeout time.Duration) *SendTransaction2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send transaction2 params
func (o *SendTransaction2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send transaction2 params
func (o *SendTransaction2Params) WithContext(ctx context.Context) *SendTransaction2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send transaction2 params
func (o *SendTransaction2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send transaction2 params
func (o *SendTransaction2Params) WithHTTPClient(client *http.Client) *SendTransaction2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send transaction2 params
func (o *SendTransaction2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send transaction2 params
func (o *SendTransaction2Params) WithBody(body *models.SendTransactionRequest) *SendTransaction2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send transaction2 params
func (o *SendTransaction2Params) SetBody(body *models.SendTransactionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SendTransaction2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}