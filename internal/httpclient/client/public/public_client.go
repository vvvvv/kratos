// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new public API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CompleteSelfServiceBrowserProfileManagementFlow completes the browser based profile management flows

This endpoint completes a browser-based profile management flow. This is usually achieved by POSTing data to this
endpoint.

If the provided profile data is valid against the Identity's Traits JSON Schema, the data will be updated and
the browser redirected to `url.profile_ui` for further steps.

> This endpoint is NOT INTENDED for API clients and only works with browsers (Chrome, Firefox, ...) and HTML Forms.

More information can be found at [ORY Kratos Profile Management Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-profile-management).
*/
func (a *Client) CompleteSelfServiceBrowserProfileManagementFlow(params *CompleteSelfServiceBrowserProfileManagementFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteSelfServiceBrowserProfileManagementFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "completeSelfServiceBrowserProfileManagementFlow",
		Method:             "POST",
		PathPattern:        "/self-service/browser/flows/profile/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CompleteSelfServiceBrowserProfileManagementFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetSelfServiceBrowserLoginRequest gets the request context of browser based login user flows

This endpoint returns a login request's context with, for example, error details and
other information.

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) GetSelfServiceBrowserLoginRequest(params *GetSelfServiceBrowserLoginRequestParams) (*GetSelfServiceBrowserLoginRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceBrowserLoginRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceBrowserLoginRequest",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/requests/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceBrowserLoginRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceBrowserLoginRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceBrowserLoginRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSelfServiceBrowserProfileManagementRequest gets the request context of browser based profile management flows

More information can be found at [ORY Kratos Profile Management Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-profile-management).
*/
func (a *Client) GetSelfServiceBrowserProfileManagementRequest(params *GetSelfServiceBrowserProfileManagementRequestParams) (*GetSelfServiceBrowserProfileManagementRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceBrowserProfileManagementRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceBrowserProfileManagementRequest",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/requests/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceBrowserProfileManagementRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceBrowserProfileManagementRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceBrowserProfileManagementRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSelfServiceBrowserRegistrationRequest gets the request context of browser based registration user flows

This endpoint returns a registration request's context with, for example, error details and
other information.

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) GetSelfServiceBrowserRegistrationRequest(params *GetSelfServiceBrowserRegistrationRequestParams) (*GetSelfServiceBrowserRegistrationRequestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSelfServiceBrowserRegistrationRequestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSelfServiceBrowserRegistrationRequest",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/requests/registration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSelfServiceBrowserRegistrationRequestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSelfServiceBrowserRegistrationRequestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSelfServiceBrowserRegistrationRequest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
InitializeSelfServiceBrowserLoginFlow initializes browser based login user flow

This endpoint initializes a browser-based user login flow. Once initialized, the browser will be redirected to
`urls.login_ui` with the request ID set as a query parameter. If a valid user session exists already, the browser will be
redirected to `urls.default_redirect_url`.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) InitializeSelfServiceBrowserLoginFlow(params *InitializeSelfServiceBrowserLoginFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceBrowserLoginFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceBrowserLoginFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceBrowserLoginFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
InitializeSelfServiceBrowserLogoutFlow initializes browser based logout user flow

This endpoint initializes a logout flow.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

On successful logout, the browser will be redirected (HTTP 302 Found) to `urls.default_return_to`.

More information can be found at [ORY Kratos User Logout Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-logout).
*/
func (a *Client) InitializeSelfServiceBrowserLogoutFlow(params *InitializeSelfServiceBrowserLogoutFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceBrowserLogoutFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceBrowserLogoutFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceBrowserLogoutFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
InitializeSelfServiceBrowserRegistrationFlow initializes browser based registration user flow

This endpoint initializes a browser-based user registration flow. Once initialized, the browser will be redirected to
`urls.registration_ui` with the request ID set as a query parameter. If a valid user session exists already, the browser will be
redirected to `urls.default_redirect_url`.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos User Login and User Registration Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-login-user-registration).
*/
func (a *Client) InitializeSelfServiceBrowserRegistrationFlow(params *InitializeSelfServiceBrowserRegistrationFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceBrowserRegistrationFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceBrowserRegistrationFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/registration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceBrowserRegistrationFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
InitializeSelfServiceProfileManagementFlow initializes browser based profile management flow

This endpoint initializes a browser-based profile management flow. Once initialized, the browser will be redirected to
`urls.profile_ui` with the request ID set as a query parameter. If no valid user session exists, a login
flow will be initialized.

> This endpoint is NOT INTENDED for API clients and only works
with browsers (Chrome, Firefox, ...).

More information can be found at [ORY Kratos Profile Management Documentation](https://www.ory.sh/docs/next/kratos/self-service/flows/user-profile-management).
*/
func (a *Client) InitializeSelfServiceProfileManagementFlow(params *InitializeSelfServiceProfileManagementFlowParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitializeSelfServiceProfileManagementFlowParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initializeSelfServiceProfileManagementFlow",
		Method:             "GET",
		PathPattern:        "/self-service/browser/flows/profile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InitializeSelfServiceProfileManagementFlowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
Whoami checks who the current HTTP session belongs to

Uses the HTTP Headers in the GET request to determine (e.g. by using checking the cookies) who is authenticated.
Returns a session object or 401 if the credentials are invalid or no credentials were sent.

This endpoint is useful for reverse proxies and API Gateways.
*/
func (a *Client) Whoami(params *WhoamiParams) (*WhoamiOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWhoamiParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "whoami",
		Method:             "GET",
		PathPattern:        "/sessions/whoami",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WhoamiReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WhoamiOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for whoami: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
