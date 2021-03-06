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
	"github.com/go-openapi/swag"
)

// NewMaxGasPriceParams creates a new MaxGasPriceParams object
// with the default values initialized.
func NewMaxGasPriceParams() *MaxGasPriceParams {
	var ()
	return &MaxGasPriceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMaxGasPriceParamsWithTimeout creates a new MaxGasPriceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMaxGasPriceParamsWithTimeout(timeout time.Duration) *MaxGasPriceParams {
	var ()
	return &MaxGasPriceParams{

		timeout: timeout,
	}
}

// NewMaxGasPriceParamsWithContext creates a new MaxGasPriceParams object
// with the default values initialized, and the ability to set a context for a request
func NewMaxGasPriceParamsWithContext(ctx context.Context) *MaxGasPriceParams {
	var ()
	return &MaxGasPriceParams{

		Context: ctx,
	}
}

// NewMaxGasPriceParamsWithHTTPClient creates a new MaxGasPriceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMaxGasPriceParamsWithHTTPClient(client *http.Client) *MaxGasPriceParams {
	var ()
	return &MaxGasPriceParams{
		HTTPClient: client,
	}
}

/*MaxGasPriceParams contains all the parameters to send to the API endpoint
for the max gas price operation typically these are written to a http.Request
*/
type MaxGasPriceParams struct {

	/*Height*/
	Height *uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the max gas price params
func (o *MaxGasPriceParams) WithTimeout(timeout time.Duration) *MaxGasPriceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the max gas price params
func (o *MaxGasPriceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the max gas price params
func (o *MaxGasPriceParams) WithContext(ctx context.Context) *MaxGasPriceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the max gas price params
func (o *MaxGasPriceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the max gas price params
func (o *MaxGasPriceParams) WithHTTPClient(client *http.Client) *MaxGasPriceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the max gas price params
func (o *MaxGasPriceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the max gas price params
func (o *MaxGasPriceParams) WithHeight(height *uint64) *MaxGasPriceParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the max gas price params
func (o *MaxGasPriceParams) SetHeight(height *uint64) {
	o.Height = height
}

// WriteToRequest writes these params to a swagger request
func (o *MaxGasPriceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Height != nil {

		// query param height
		var qrHeight uint64
		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := swag.FormatUint64(qrHeight)
		if qHeight != "" {
			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
