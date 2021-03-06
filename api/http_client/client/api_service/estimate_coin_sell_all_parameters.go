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

// NewEstimateCoinSellAllParams creates a new EstimateCoinSellAllParams object
// with the default values initialized.
func NewEstimateCoinSellAllParams() *EstimateCoinSellAllParams {
	var (
		gasPriceDefault = uint64(1)
	)
	return &EstimateCoinSellAllParams{
		GasPrice: &gasPriceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewEstimateCoinSellAllParamsWithTimeout creates a new EstimateCoinSellAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEstimateCoinSellAllParamsWithTimeout(timeout time.Duration) *EstimateCoinSellAllParams {
	var (
		gasPriceDefault = uint64(1)
	)
	return &EstimateCoinSellAllParams{
		GasPrice: &gasPriceDefault,

		timeout: timeout,
	}
}

// NewEstimateCoinSellAllParamsWithContext creates a new EstimateCoinSellAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewEstimateCoinSellAllParamsWithContext(ctx context.Context) *EstimateCoinSellAllParams {
	var (
		gasPriceDefault = uint64(1)
	)
	return &EstimateCoinSellAllParams{
		GasPrice: &gasPriceDefault,

		Context: ctx,
	}
}

// NewEstimateCoinSellAllParamsWithHTTPClient creates a new EstimateCoinSellAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEstimateCoinSellAllParamsWithHTTPClient(client *http.Client) *EstimateCoinSellAllParams {
	var (
		gasPriceDefault = uint64(1)
	)
	return &EstimateCoinSellAllParams{
		GasPrice:   &gasPriceDefault,
		HTTPClient: client,
	}
}

/*EstimateCoinSellAllParams contains all the parameters to send to the API endpoint
for the estimate coin sell all operation typically these are written to a http.Request
*/
type EstimateCoinSellAllParams struct {

	/*CoinIDToBuy*/
	CoinIDToBuy *uint64
	/*CoinIDToSell*/
	CoinIDToSell *uint64
	/*CoinToBuy*/
	CoinToBuy *string
	/*CoinToSell*/
	CoinToSell *string
	/*GasPrice*/
	GasPrice *uint64
	/*Height*/
	Height *uint64
	/*ValueToSell*/
	ValueToSell string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithTimeout(timeout time.Duration) *EstimateCoinSellAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithContext(ctx context.Context) *EstimateCoinSellAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithHTTPClient(client *http.Client) *EstimateCoinSellAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCoinIDToBuy adds the coinIDToBuy to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithCoinIDToBuy(coinIDToBuy *uint64) *EstimateCoinSellAllParams {
	o.SetCoinIDToBuy(coinIDToBuy)
	return o
}

// SetCoinIDToBuy adds the coinIdToBuy to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetCoinIDToBuy(coinIDToBuy *uint64) {
	o.CoinIDToBuy = coinIDToBuy
}

// WithCoinIDToSell adds the coinIDToSell to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithCoinIDToSell(coinIDToSell *uint64) *EstimateCoinSellAllParams {
	o.SetCoinIDToSell(coinIDToSell)
	return o
}

// SetCoinIDToSell adds the coinIdToSell to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetCoinIDToSell(coinIDToSell *uint64) {
	o.CoinIDToSell = coinIDToSell
}

// WithCoinToBuy adds the coinToBuy to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithCoinToBuy(coinToBuy *string) *EstimateCoinSellAllParams {
	o.SetCoinToBuy(coinToBuy)
	return o
}

// SetCoinToBuy adds the coinToBuy to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetCoinToBuy(coinToBuy *string) {
	o.CoinToBuy = coinToBuy
}

// WithCoinToSell adds the coinToSell to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithCoinToSell(coinToSell *string) *EstimateCoinSellAllParams {
	o.SetCoinToSell(coinToSell)
	return o
}

// SetCoinToSell adds the coinToSell to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetCoinToSell(coinToSell *string) {
	o.CoinToSell = coinToSell
}

// WithGasPrice adds the gasPrice to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithGasPrice(gasPrice *uint64) *EstimateCoinSellAllParams {
	o.SetGasPrice(gasPrice)
	return o
}

// SetGasPrice adds the gasPrice to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetGasPrice(gasPrice *uint64) {
	o.GasPrice = gasPrice
}

// WithHeight adds the height to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithHeight(height *uint64) *EstimateCoinSellAllParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetHeight(height *uint64) {
	o.Height = height
}

// WithValueToSell adds the valueToSell to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) WithValueToSell(valueToSell string) *EstimateCoinSellAllParams {
	o.SetValueToSell(valueToSell)
	return o
}

// SetValueToSell adds the valueToSell to the estimate coin sell all params
func (o *EstimateCoinSellAllParams) SetValueToSell(valueToSell string) {
	o.ValueToSell = valueToSell
}

// WriteToRequest writes these params to a swagger request
func (o *EstimateCoinSellAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CoinIDToBuy != nil {

		// query param coin_id_to_buy
		var qrCoinIDToBuy uint64
		if o.CoinIDToBuy != nil {
			qrCoinIDToBuy = *o.CoinIDToBuy
		}
		qCoinIDToBuy := swag.FormatUint64(qrCoinIDToBuy)
		if qCoinIDToBuy != "" {
			if err := r.SetQueryParam("coin_id_to_buy", qCoinIDToBuy); err != nil {
				return err
			}
		}

	}

	if o.CoinIDToSell != nil {

		// query param coin_id_to_sell
		var qrCoinIDToSell uint64
		if o.CoinIDToSell != nil {
			qrCoinIDToSell = *o.CoinIDToSell
		}
		qCoinIDToSell := swag.FormatUint64(qrCoinIDToSell)
		if qCoinIDToSell != "" {
			if err := r.SetQueryParam("coin_id_to_sell", qCoinIDToSell); err != nil {
				return err
			}
		}

	}

	if o.CoinToBuy != nil {

		// query param coin_to_buy
		var qrCoinToBuy string
		if o.CoinToBuy != nil {
			qrCoinToBuy = *o.CoinToBuy
		}
		qCoinToBuy := qrCoinToBuy
		if qCoinToBuy != "" {
			if err := r.SetQueryParam("coin_to_buy", qCoinToBuy); err != nil {
				return err
			}
		}

	}

	if o.CoinToSell != nil {

		// query param coin_to_sell
		var qrCoinToSell string
		if o.CoinToSell != nil {
			qrCoinToSell = *o.CoinToSell
		}
		qCoinToSell := qrCoinToSell
		if qCoinToSell != "" {
			if err := r.SetQueryParam("coin_to_sell", qCoinToSell); err != nil {
				return err
			}
		}

	}

	if o.GasPrice != nil {

		// query param gas_price
		var qrGasPrice uint64
		if o.GasPrice != nil {
			qrGasPrice = *o.GasPrice
		}
		qGasPrice := swag.FormatUint64(qrGasPrice)
		if qGasPrice != "" {
			if err := r.SetQueryParam("gas_price", qGasPrice); err != nil {
				return err
			}
		}

	}

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

	// query param value_to_sell
	qrValueToSell := o.ValueToSell
	qValueToSell := qrValueToSell
	if qValueToSell != "" {
		if err := r.SetQueryParam("value_to_sell", qValueToSell); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
