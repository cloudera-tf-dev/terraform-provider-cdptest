// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeVwResponse Response object for the describeVw method.
//
// swagger:model DescribeVwResponse
type DescribeVwResponse struct {

	// The Virtual Warehouse.
	Vw *VwSummary `json:"vw,omitempty"`
}

// Validate validates this describe vw response
func (m *DescribeVwResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVw(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeVwResponse) validateVw(formats strfmt.Registry) error {
	if swag.IsZero(m.Vw) { // not required
		return nil
	}

	if m.Vw != nil {
		if err := m.Vw.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vw")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe vw response based on the context it is used
func (m *DescribeVwResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVw(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeVwResponse) contextValidateVw(ctx context.Context, formats strfmt.Registry) error {

	if m.Vw != nil {
		if err := m.Vw.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vw")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vw")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeVwResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeVwResponse) UnmarshalBinary(b []byte) error {
	var res DescribeVwResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
