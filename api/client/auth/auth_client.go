package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new auth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
UserLogin logs the user

User login
*/
func (a *Client) UserLogin(params *UserLoginParams) (*UserLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserLoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "userLogin",
		Method:             "POST",
		PathPattern:        "/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UserLoginReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserLoginOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}