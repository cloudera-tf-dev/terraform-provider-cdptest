// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListDbcsResponse Response object for the listDbcs method.
//
// swagger:model ListDbcsResponse
type ListDbcsResponse struct {

	// The list of Database Catalogs.
	Dbcs []*DbcSummary `json:"dbcs"`
}

// Validate validates this list dbcs response
func (m *ListDbcsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDbcs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDbcsResponse) validateDbcs(formats strfmt.Registry) error {
	if swag.IsZero(m.Dbcs) { // not required
		return nil
	}

	for i := 0; i < len(m.Dbcs); i++ {
		if swag.IsZero(m.Dbcs[i]) { // not required
			continue
		}

		if m.Dbcs[i] != nil {
			if err := m.Dbcs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dbcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dbcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list dbcs response based on the context it is used
func (m *ListDbcsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDbcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDbcsResponse) contextValidateDbcs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dbcs); i++ {

		if m.Dbcs[i] != nil {
			if err := m.Dbcs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dbcs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dbcs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListDbcsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDbcsResponse) UnmarshalBinary(b []byte) error {
	var res ListDbcsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
