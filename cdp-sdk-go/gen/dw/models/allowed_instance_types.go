// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AllowedInstanceTypes Allowed compute instance type values and default compute instance type value.
//
// swagger:model AllowedInstanceTypes
type AllowedInstanceTypes struct {

	// Allowed values for the instance type.
	Allowed []string `json:"allowed"`

	// Default value for the compute instance type usage.
	Default []string `json:"default"`
}

// Validate validates this allowed instance types
func (m *AllowedInstanceTypes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this allowed instance types based on context it is used
func (m *AllowedInstanceTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AllowedInstanceTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllowedInstanceTypes) UnmarshalBinary(b []byte) error {
	var res AllowedInstanceTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
