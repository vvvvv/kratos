// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/kratos/internal/httpclient/models"
)

// InitializeSelfServiceBrowserLoginFlowReader is a Reader for the InitializeSelfServiceBrowserLoginFlow structure.
type InitializeSelfServiceBrowserLoginFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitializeSelfServiceBrowserLoginFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewInitializeSelfServiceBrowserLoginFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInitializeSelfServiceBrowserLoginFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInitializeSelfServiceBrowserLoginFlowFound creates a InitializeSelfServiceBrowserLoginFlowFound with default headers values
func NewInitializeSelfServiceBrowserLoginFlowFound() *InitializeSelfServiceBrowserLoginFlowFound {
	return &InitializeSelfServiceBrowserLoginFlowFound{}
}

/*InitializeSelfServiceBrowserLoginFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type InitializeSelfServiceBrowserLoginFlowFound struct {
}

func (o *InitializeSelfServiceBrowserLoginFlowFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/login][%d] initializeSelfServiceBrowserLoginFlowFound ", 302)
}

func (o *InitializeSelfServiceBrowserLoginFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitializeSelfServiceBrowserLoginFlowInternalServerError creates a InitializeSelfServiceBrowserLoginFlowInternalServerError with default headers values
func NewInitializeSelfServiceBrowserLoginFlowInternalServerError() *InitializeSelfServiceBrowserLoginFlowInternalServerError {
	return &InitializeSelfServiceBrowserLoginFlowInternalServerError{}
}

/*InitializeSelfServiceBrowserLoginFlowInternalServerError handles this case with default header values.

genericError
*/
type InitializeSelfServiceBrowserLoginFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *InitializeSelfServiceBrowserLoginFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/login][%d] initializeSelfServiceBrowserLoginFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *InitializeSelfServiceBrowserLoginFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *InitializeSelfServiceBrowserLoginFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
