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
	"github.com/go-openapi/validate"
)

// DeleteClusterDefinitionsResponse Response object for delete cluster definition request.
//
// swagger:model DeleteClusterDefinitionsResponse
type DeleteClusterDefinitionsResponse struct {

	// The clusterDefinitions.
	// Required: true
	ClusterDefinitions []*ClusterDefinition `json:"clusterDefinitions"`
}

// Validate validates this delete cluster definitions response
func (m *DeleteClusterDefinitionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterDefinitions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteClusterDefinitionsResponse) validateClusterDefinitions(formats strfmt.Registry) error {

	if err := validate.Required("clusterDefinitions", "body", m.ClusterDefinitions); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterDefinitions); i++ {
		if swag.IsZero(m.ClusterDefinitions[i]) { // not required
			continue
		}

		if m.ClusterDefinitions[i] != nil {
			if err := m.ClusterDefinitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterDefinitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterDefinitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete cluster definitions response based on the context it is used
func (m *DeleteClusterDefinitionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterDefinitions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteClusterDefinitionsResponse) contextValidateClusterDefinitions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterDefinitions); i++ {

		if m.ClusterDefinitions[i] != nil {
			if err := m.ClusterDefinitions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterDefinitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterDefinitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteClusterDefinitionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteClusterDefinitionsResponse) UnmarshalBinary(b []byte) error {
	var res DeleteClusterDefinitionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
