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

// NewAPIServiceEstimateTxCommission2Params creates a new APIServiceEstimateTxCommission2Params object
// with the default values initialized.
func NewAPIServiceEstimateTxCommission2Params() *APIServiceEstimateTxCommission2Params {
	var (
		dataTypeDefault = string("unknown")
	)
	return &APIServiceEstimateTxCommission2Params{
		DataType: &dataTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceEstimateTxCommission2ParamsWithTimeout creates a new APIServiceEstimateTxCommission2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceEstimateTxCommission2ParamsWithTimeout(timeout time.Duration) *APIServiceEstimateTxCommission2Params {
	var (
		dataTypeDefault = string("unknown")
	)
	return &APIServiceEstimateTxCommission2Params{
		DataType: &dataTypeDefault,

		timeout: timeout,
	}
}

// NewAPIServiceEstimateTxCommission2ParamsWithContext creates a new APIServiceEstimateTxCommission2Params object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceEstimateTxCommission2ParamsWithContext(ctx context.Context) *APIServiceEstimateTxCommission2Params {
	var (
		dataTypeDefault = string("unknown")
	)
	return &APIServiceEstimateTxCommission2Params{
		DataType: &dataTypeDefault,

		Context: ctx,
	}
}

// NewAPIServiceEstimateTxCommission2ParamsWithHTTPClient creates a new APIServiceEstimateTxCommission2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceEstimateTxCommission2ParamsWithHTTPClient(client *http.Client) *APIServiceEstimateTxCommission2Params {
	var (
		dataTypeDefault = string("unknown")
	)
	return &APIServiceEstimateTxCommission2Params{
		DataType:   &dataTypeDefault,
		HTTPClient: client,
	}
}

/*APIServiceEstimateTxCommission2Params contains all the parameters to send to the API endpoint
for the Api service estimate tx commission2 operation typically these are written to a http.Request
*/
type APIServiceEstimateTxCommission2Params struct {

	/*DataGasCoinID*/
	DataGasCoinID *int64
	/*DataMtxs*/
	DataMtxs *string
	/*DataPayload*/
	DataPayload *strfmt.Base64
	/*DataType*/
	DataType *string
	/*Height*/
	Height *string
	/*Tx*/
	Tx *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) WithTimeout(timeout time.Duration) *APIServiceEstimateTxCommission2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) WithContext(ctx context.Context) *APIServiceEstimateTxCommission2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) WithHTTPClient(client *http.Client) *APIServiceEstimateTxCommission2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataGasCoinID adds the dataGasCoinID to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) WithDataGasCoinID(dataGasCoinID *int64) *APIServiceEstimateTxCommission2Params {
	o.SetDataGasCoinID(dataGasCoinID)
	return o
}

// SetDataGasCoinID adds the dataGasCoinId to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) SetDataGasCoinID(dataGasCoinID *int64) {
	o.DataGasCoinID = dataGasCoinID
}

// WithDataMtxs adds the dataMtxs to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) WithDataMtxs(dataMtxs *string) *APIServiceEstimateTxCommission2Params {
	o.SetDataMtxs(dataMtxs)
	return o
}

// SetDataMtxs adds the dataMtxs to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) SetDataMtxs(dataMtxs *string) {
	o.DataMtxs = dataMtxs
}

// WithDataPayload adds the dataPayload to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) WithDataPayload(dataPayload *strfmt.Base64) *APIServiceEstimateTxCommission2Params {
	o.SetDataPayload(dataPayload)
	return o
}

// SetDataPayload adds the dataPayload to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) SetDataPayload(dataPayload *strfmt.Base64) {
	o.DataPayload = dataPayload
}

// WithDataType adds the dataType to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) WithDataType(dataType *string) *APIServiceEstimateTxCommission2Params {
	o.SetDataType(dataType)
	return o
}

// SetDataType adds the dataType to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) SetDataType(dataType *string) {
	o.DataType = dataType
}

// WithHeight adds the height to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) WithHeight(height *string) *APIServiceEstimateTxCommission2Params {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) SetHeight(height *string) {
	o.Height = height
}

// WithTx adds the tx to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) WithTx(tx *string) *APIServiceEstimateTxCommission2Params {
	o.SetTx(tx)
	return o
}

// SetTx adds the tx to the Api service estimate tx commission2 params
func (o *APIServiceEstimateTxCommission2Params) SetTx(tx *string) {
	o.Tx = tx
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceEstimateTxCommission2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DataGasCoinID != nil {

		// query param data.gas_coin_id
		var qrDataGasCoinID int64
		if o.DataGasCoinID != nil {
			qrDataGasCoinID = *o.DataGasCoinID
		}
		qDataGasCoinID := swag.FormatInt64(qrDataGasCoinID)
		if qDataGasCoinID != "" {
			if err := r.SetQueryParam("data.gas_coin_id", qDataGasCoinID); err != nil {
				return err
			}
		}

	}

	if o.DataMtxs != nil {

		// query param data.mtxs
		var qrDataMtxs string
		if o.DataMtxs != nil {
			qrDataMtxs = *o.DataMtxs
		}
		qDataMtxs := qrDataMtxs
		if qDataMtxs != "" {
			if err := r.SetQueryParam("data.mtxs", qDataMtxs); err != nil {
				return err
			}
		}

	}

	if o.DataPayload != nil {

		// query param data.payload
		var qrDataPayload strfmt.Base64
		if o.DataPayload != nil {
			qrDataPayload = *o.DataPayload
		}
		qDataPayload := qrDataPayload.String()
		if qDataPayload != "" {
			if err := r.SetQueryParam("data.payload", qDataPayload); err != nil {
				return err
			}
		}

	}

	if o.DataType != nil {

		// query param data.type
		var qrDataType string
		if o.DataType != nil {
			qrDataType = *o.DataType
		}
		qDataType := qrDataType
		if qDataType != "" {
			if err := r.SetQueryParam("data.type", qDataType); err != nil {
				return err
			}
		}

	}

	if o.Height != nil {

		// query param height
		var qrHeight string
		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := qrHeight
		if qHeight != "" {
			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}

	}

	if o.Tx != nil {

		// query param tx
		var qrTx string
		if o.Tx != nil {
			qrTx = *o.Tx
		}
		qTx := qrTx
		if qTx != "" {
			if err := r.SetQueryParam("tx", qTx); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}