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

// GetLogsResponse GetLogsResponse contains all the logs for a given request id.
//
// swagger:model GetLogsResponse
type GetLogsResponse struct {

	// Contains all the logs for a given request id.
	// Required: true
	AuditLogs *AuditLogs `json:"auditLogs"`
}

// Validate validates this get logs response
func (m *GetLogsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetLogsResponse) validateAuditLogs(formats strfmt.Registry) error {

	if err := validate.Required("auditLogs", "body", m.AuditLogs); err != nil {
		return err
	}

	if m.AuditLogs != nil {
		if err := m.AuditLogs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auditLogs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auditLogs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get logs response based on the context it is used
func (m *GetLogsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetLogsResponse) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

	if m.AuditLogs != nil {
		if err := m.AuditLogs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auditLogs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auditLogs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetLogsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLogsResponse) UnmarshalBinary(b []byte) error {
	var res GetLogsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
