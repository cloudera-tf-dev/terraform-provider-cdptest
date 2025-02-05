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

// VwFilter Filter object for Virtual Warehouses.
//
// swagger:model VwFilter
type VwFilter struct {

	// Filter Virtual Warehouses whether it is a Compactor or not.
	Compactor bool `json:"compactor,omitempty"`

	// Filter Virtual Warehouses by Database Catalog Id.
	DbcID string `json:"dbcId,omitempty"`

	// Filter the Virtual Warehouses based on whether it has Data Visualisation.
	Viz bool `json:"viz,omitempty"`

	// Virtual Warehouse type.
	VwType VwType `json:"vwType,omitempty"`
}

// Validate validates this vw filter
func (m *VwFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVwType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VwFilter) validateVwType(formats strfmt.Registry) error {
	if swag.IsZero(m.VwType) { // not required
		return nil
	}

	if err := m.VwType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vwType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vwType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this vw filter based on the context it is used
func (m *VwFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVwType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VwFilter) contextValidateVwType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VwType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vwType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vwType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VwFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VwFilter) UnmarshalBinary(b []byte) error {
	var res VwFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
