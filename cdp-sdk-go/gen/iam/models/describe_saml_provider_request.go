// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeSamlProviderRequest Request object for a describe SAML provider request.
//
// swagger:model DescribeSamlProviderRequest
type DescribeSamlProviderRequest struct {

	// The name or CRN of the SAML provider to describe.
	SamlProviderName string `json:"samlProviderName,omitempty"`
}

// Validate validates this describe saml provider request
func (m *DescribeSamlProviderRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this describe saml provider request based on context it is used
func (m *DescribeSamlProviderRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeSamlProviderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeSamlProviderRequest) UnmarshalBinary(b []byte) error {
	var res DescribeSamlProviderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
