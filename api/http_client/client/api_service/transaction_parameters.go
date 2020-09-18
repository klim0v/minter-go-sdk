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
)

// NewTransactionParams creates a new TransactionParams object
// with the default values initialized.
func NewTransactionParams() *TransactionParams {
	var ()
	return &TransactionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTransactionParamsWithTimeout creates a new TransactionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTransactionParamsWithTimeout(timeout time.Duration) *TransactionParams {
	var ()
	return &TransactionParams{

		timeout: timeout,
	}
}

// NewTransactionParamsWithContext creates a new TransactionParams object
// with the default values initialized, and the ability to set a context for a request
func NewTransactionParamsWithContext(ctx context.Context) *TransactionParams {
	var ()
	return &TransactionParams{

		Context: ctx,
	}
}

// NewTransactionParamsWithHTTPClient creates a new TransactionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTransactionParamsWithHTTPClient(client *http.Client) *TransactionParams {
	var ()
	return &TransactionParams{
		HTTPClient: client,
	}
}

/*TransactionParams contains all the parameters to send to the API endpoint
for the transaction operation typically these are written to a http.Request
*/
type TransactionParams struct {

	/*Hash*/
	Hash string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the transaction params
func (o *TransactionParams) WithTimeout(timeout time.Duration) *TransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transaction params
func (o *TransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transaction params
func (o *TransactionParams) WithContext(ctx context.Context) *TransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transaction params
func (o *TransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transaction params
func (o *TransactionParams) WithHTTPClient(client *http.Client) *TransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transaction params
func (o *TransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHash adds the hash to the transaction params
func (o *TransactionParams) WithHash(hash string) *TransactionParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the transaction params
func (o *TransactionParams) SetHash(hash string) {
	o.Hash = hash
}

// WriteToRequest writes these params to a swagger request
func (o *TransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param hash
	if err := r.SetPathParam("hash", o.Hash); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
