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

// NewCampusCreateParams creates a new CampusCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCampusCreateParams() *CampusCreateParams {
	return &CampusCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCampusCreateParamsWithTimeout creates a new CampusCreateParams object
// with the ability to set a timeout on a request.
func NewCampusCreateParamsWithTimeout(timeout time.Duration) *CampusCreateParams {
	return &CampusCreateParams{
		timeout: timeout,
	}
}

// NewCampusCreateParamsWithContext creates a new CampusCreateParams object
// with the ability to set a context for a request.
func NewCampusCreateParamsWithContext(ctx context.Context) *CampusCreateParams {
	return &CampusCreateParams{
		Context: ctx,
	}
}

// NewCampusCreateParamsWithHTTPClient creates a new CampusCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCampusCreateParamsWithHTTPClient(client *http.Client) *CampusCreateParams {
	return &CampusCreateParams{
		HTTPClient: client,
	}
}

/*
CampusCreateParams contains all the parameters to send to the API endpoint

	for the campus create operation.

	Typically these are written to a http.Request.
*/
type CampusCreateParams struct {

	/* ID.

	   Campus ID.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the campus create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CampusCreateParams) WithDefaults() *CampusCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the campus create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CampusCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the campus create params
func (o *CampusCreateParams) WithTimeout(timeout time.Duration) *CampusCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the campus create params
func (o *CampusCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the campus create params
func (o *CampusCreateParams) WithContext(ctx context.Context) *CampusCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the campus create params
func (o *CampusCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the campus create params
func (o *CampusCreateParams) WithHTTPClient(client *http.Client) *CampusCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the campus create params
func (o *CampusCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the campus create params
func (o *CampusCreateParams) WithID(id string) *CampusCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the campus create params
func (o *CampusCreateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CampusCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
