// Code generated by go-swagger; DO NOT EDIT.

package notify

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

// NewNotifySMSParams creates a new NotifySMSParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotifySMSParams() *NotifySMSParams {
	return &NotifySMSParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotifySMSParamsWithTimeout creates a new NotifySMSParams object
// with the ability to set a timeout on a request.
func NewNotifySMSParamsWithTimeout(timeout time.Duration) *NotifySMSParams {
	return &NotifySMSParams{
		timeout: timeout,
	}
}

// NewNotifySMSParamsWithContext creates a new NotifySMSParams object
// with the ability to set a context for a request.
func NewNotifySMSParamsWithContext(ctx context.Context) *NotifySMSParams {
	return &NotifySMSParams{
		Context: ctx,
	}
}

// NewNotifySMSParamsWithHTTPClient creates a new NotifySMSParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotifySMSParamsWithHTTPClient(client *http.Client) *NotifySMSParams {
	return &NotifySMSParams{
		HTTPClient: client,
	}
}

/*
NotifySMSParams contains all the parameters to send to the API endpoint

	for the notify s m s operation.

	Typically these are written to a http.Request.
*/
type NotifySMSParams struct {

	/* Text.

	   Текст сообщения
	*/
	Text string

	/* To.

	   Телефон получателя
	*/
	To string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notify s m s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifySMSParams) WithDefaults() *NotifySMSParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notify s m s params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotifySMSParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notify s m s params
func (o *NotifySMSParams) WithTimeout(timeout time.Duration) *NotifySMSParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notify s m s params
func (o *NotifySMSParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notify s m s params
func (o *NotifySMSParams) WithContext(ctx context.Context) *NotifySMSParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notify s m s params
func (o *NotifySMSParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notify s m s params
func (o *NotifySMSParams) WithHTTPClient(client *http.Client) *NotifySMSParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notify s m s params
func (o *NotifySMSParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithText adds the text to the notify s m s params
func (o *NotifySMSParams) WithText(text string) *NotifySMSParams {
	o.SetText(text)
	return o
}

// SetText adds the text to the notify s m s params
func (o *NotifySMSParams) SetText(text string) {
	o.Text = text
}

// WithTo adds the to to the notify s m s params
func (o *NotifySMSParams) WithTo(to string) *NotifySMSParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the notify s m s params
func (o *NotifySMSParams) SetTo(to string) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *NotifySMSParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param text
	frText := o.Text
	fText := frText
	if fText != "" {
		if err := r.SetFormParam("text", fText); err != nil {
			return err
		}
	}

	// form param to
	frTo := o.To
	fTo := frTo
	if fTo != "" {
		if err := r.SetFormParam("to", fTo); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
