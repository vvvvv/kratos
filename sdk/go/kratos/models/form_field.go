// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FormField Field represents a HTML Form Field
// swagger:model formField
type FormField struct {

	// Errors contains all validation errors this particular field has caused.
	Errors []*Error `json:"errors"`

	// Name is the equivalent of <input name="{{.Name}}">
	Name string `json:"name,omitempty"`

	// Name is the equivalent of <input required="{{.Required}}">
	Required bool `json:"required,omitempty"`

	// Name is the equivalent of <input type="{{.Type}}">
	Type string `json:"type,omitempty"`

	// Name is the equivalent of <input value="{{.Value}}">
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this form field
func (m *FormField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FormField) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FormField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FormField) UnmarshalBinary(b []byte) error {
	var res FormField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
