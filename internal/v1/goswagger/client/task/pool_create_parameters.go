// Code generated by go-swagger; DO NOT EDIT.

package task

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

// NewPoolCreateParams creates a new PoolCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPoolCreateParams() *PoolCreateParams {
	return &PoolCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPoolCreateParamsWithTimeout creates a new PoolCreateParams object
// with the ability to set a timeout on a request.
func NewPoolCreateParamsWithTimeout(timeout time.Duration) *PoolCreateParams {
	return &PoolCreateParams{
		timeout: timeout,
	}
}

// NewPoolCreateParamsWithContext creates a new PoolCreateParams object
// with the ability to set a context for a request.
func NewPoolCreateParamsWithContext(ctx context.Context) *PoolCreateParams {
	return &PoolCreateParams{
		Context: ctx,
	}
}

// NewPoolCreateParamsWithHTTPClient creates a new PoolCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPoolCreateParamsWithHTTPClient(client *http.Client) *PoolCreateParams {
	return &PoolCreateParams{
		HTTPClient: client,
	}
}

/*
PoolCreateParams contains all the parameters to send to the API endpoint

	for the pool create operation.

	Typically these are written to a http.Request.
*/
type PoolCreateParams struct {

	/* ID.

	   Pool event ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pool create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoolCreateParams) WithDefaults() *PoolCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pool create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PoolCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pool create params
func (o *PoolCreateParams) WithTimeout(timeout time.Duration) *PoolCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pool create params
func (o *PoolCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pool create params
func (o *PoolCreateParams) WithContext(ctx context.Context) *PoolCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pool create params
func (o *PoolCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pool create params
func (o *PoolCreateParams) WithHTTPClient(client *http.Client) *PoolCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pool create params
func (o *PoolCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the pool create params
func (o *PoolCreateParams) WithID(id string) *PoolCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pool create params
func (o *PoolCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PoolCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
