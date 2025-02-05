// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAzureCredentialRequest Request object for a create Azure credential request.
//
// swagger:model CreateAzureCredentialRequest
type CreateAzureCredentialRequest struct {

	// app based
	// Required: true
	AppBased *CreateAzureCredentialRequestAppBased `json:"appBased"`

	// The name of the credential.
	// Required: true
	CredentialName *string `json:"credentialName"`

	// A description for the credential.
	Description string `json:"description,omitempty"`

	// The Azure subscription ID. Required for secret based credentials and optional for certificate based ones.
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// The Azure AD tenant ID for the Azure subscription. Required for secret based credentials and optional for certificate based ones.
	TenantID string `json:"tenantId,omitempty"`
}

// Validate validates this create azure credential request
func (m *CreateAzureCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppBased(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureCredentialRequest) validateAppBased(formats strfmt.Registry) error {

	if err := validate.Required("appBased", "body", m.AppBased); err != nil {
		return err
	}

	if m.AppBased != nil {
		if err := m.AppBased.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appBased")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appBased")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureCredentialRequest) validateCredentialName(formats strfmt.Registry) error {

	if err := validate.Required("credentialName", "body", m.CredentialName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create azure credential request based on the context it is used
func (m *CreateAzureCredentialRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppBased(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureCredentialRequest) contextValidateAppBased(ctx context.Context, formats strfmt.Registry) error {

	if m.AppBased != nil {
		if err := m.AppBased.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appBased")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appBased")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAzureCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureCredentialRequest) UnmarshalBinary(b []byte) error {
	var res CreateAzureCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateAzureCredentialRequestAppBased Additional configurations needed for app-based authentication.
//
// swagger:model CreateAzureCredentialRequestAppBased
type CreateAzureCredentialRequestAppBased struct {

	// The id of the application registered in Azure.
	ApplicationID string `json:"applicationId,omitempty"`

	// Authentication type of the credential
	AuthenticationType AzureAuthenticationTypeProperties `json:"authenticationType,omitempty"`

	// The client secret key (also referred to as application password) for the registered application.
	SecretKey string `json:"secretKey,omitempty"`
}

// Validate validates this create azure credential request app based
func (m *CreateAzureCredentialRequestAppBased) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureCredentialRequestAppBased) validateAuthenticationType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationType) { // not required
		return nil
	}

	if err := m.AuthenticationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("appBased" + "." + "authenticationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("appBased" + "." + "authenticationType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this create azure credential request app based based on the context it is used
func (m *CreateAzureCredentialRequestAppBased) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthenticationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureCredentialRequestAppBased) contextValidateAuthenticationType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AuthenticationType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("appBased" + "." + "authenticationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("appBased" + "." + "authenticationType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAzureCredentialRequestAppBased) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureCredentialRequestAppBased) UnmarshalBinary(b []byte) error {
	var res CreateAzureCredentialRequestAppBased
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
