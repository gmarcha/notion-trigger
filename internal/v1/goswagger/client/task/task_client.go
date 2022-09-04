// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new task API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for task API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CampusCreate(params *CampusCreateParams, opts ...ClientOption) (*CampusCreateCreated, error)

	PoolCreate(params *PoolCreateParams, opts ...ClientOption) (*PoolCreateCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CampusCreate creates campus tasks

Create campus tasks like onboarding steps, etc..
*/
func (a *Client) CampusCreate(params *CampusCreateParams, opts ...ClientOption) (*CampusCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCampusCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "campusCreate",
		Method:             "POST",
		PathPattern:        "/tasks/campuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CampusCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CampusCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for campusCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PoolCreate creates pool tasks

Create a task list linked to pool setup.
*/
func (a *Client) PoolCreate(params *PoolCreateParams, opts ...ClientOption) (*PoolCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPoolCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "poolCreate",
		Method:             "POST",
		PathPattern:        "/tasks/pools/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PoolCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PoolCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for poolCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}