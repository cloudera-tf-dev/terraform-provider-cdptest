// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CollectDatalakeDiagnosticsRequest Request object for collecting DataLake diagnostics.
//
// swagger:model CollectDatalakeDiagnosticsRequest
type CollectDatalakeDiagnosticsRequest struct {

	// Array of log descriptors that should be additionally collected during diagnostics collection.
	AdditionalLogs []*DatalakeVMLogRequest `json:"additionalLogs"`

	// Optional support case number in case of SUPPORT destination, otherwise only act as additional data.
	CaseNumber string `json:"caseNumber,omitempty"`

	// CRN of the Datalake cluster.
	// Required: true
	Crn *string `json:"crn"`

	// Additional information / title for the diagnostics collection.
	// Required: true
	Description *string `json:"description"`

	// Destination of the diagnostics collection (Support, Own cloud storage, Engineering or collect only on the nodes)
	// Required: true
	// Enum: [SUPPORT CLOUD_STORAGE ENG LOCAL]
	Destination *string `json:"destination"`

	// Date timestamp - collect files only for diagnostics that has lower created timestamp value than this.
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// Array of host names (fqdn or IP address), collection will not run on the excluded hosts.
	// Unique: true
	ExcludeHosts []string `json:"excludeHosts"`

	// Array of host groups, collection will run only on the dedicated hosts that belongs to these host groups.
	// Unique: true
	HostGroups []string `json:"hostGroups"`

	// Array of host names (fqdn), collection will run only on the dedicated hosts.
	// Unique: true
	Hosts []string `json:"hosts"`

	// Include Nginx report generated by GoAccess (if available).
	IncludeNginxReport bool `json:"includeNginxReport,omitempty"`

	// Include salt minion/master/api system logs in the diagnostics collection.
	IncludeSaltLogs bool `json:"includeSaltLogs,omitempty"`

	// Include SAR (System Activity Report) generated outputs in the diagnostics collection (if available).
	IncludeSarOutput bool `json:"includeSarOutput,omitempty"`

	// Array of labels that can filter logs that are collected during diagnostics collection.
	// Unique: true
	Labels []string `json:"labels"`

	// Skip unhealthy hosts from the diagnostics collection.
	SkipUnresponsiveHosts bool `json:"skipUnresponsiveHosts,omitempty"`

	// Date timestamp - collect files only for diagnostics that has higher last modified timestamp value than this.
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Enable/disable node level storage validation (can be disabled for example, if you have too many hosts and do not want to do too much parallel writes to s3/abfs)
	StorageValidation *bool `json:"storageValidation,omitempty"`

	// If enabled, required package (cdp-telemetry) will be upgraded or installed on the nodes. (useful if package is not installed or needs to be upgraded) Network is required for this operation.
	UpdatePackage *bool `json:"updatePackage,omitempty"`
}

// Validate validates this collect datalake diagnostics request
func (m *CollectDatalakeDiagnosticsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludeHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateAdditionalLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalLogs); i++ {
		if swag.IsZero(m.AdditionalLogs[i]) { // not required
			continue
		}

		if m.AdditionalLogs[i] != nil {
			if err := m.AdditionalLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

var collectDatalakeDiagnosticsRequestTypeDestinationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUPPORT","CLOUD_STORAGE","ENG","LOCAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		collectDatalakeDiagnosticsRequestTypeDestinationPropEnum = append(collectDatalakeDiagnosticsRequestTypeDestinationPropEnum, v)
	}
}

const (

	// CollectDatalakeDiagnosticsRequestDestinationSUPPORT captures enum value "SUPPORT"
	CollectDatalakeDiagnosticsRequestDestinationSUPPORT string = "SUPPORT"

	// CollectDatalakeDiagnosticsRequestDestinationCLOUDSTORAGE captures enum value "CLOUD_STORAGE"
	CollectDatalakeDiagnosticsRequestDestinationCLOUDSTORAGE string = "CLOUD_STORAGE"

	// CollectDatalakeDiagnosticsRequestDestinationENG captures enum value "ENG"
	CollectDatalakeDiagnosticsRequestDestinationENG string = "ENG"

	// CollectDatalakeDiagnosticsRequestDestinationLOCAL captures enum value "LOCAL"
	CollectDatalakeDiagnosticsRequestDestinationLOCAL string = "LOCAL"
)

// prop value enum
func (m *CollectDatalakeDiagnosticsRequest) validateDestinationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, collectDatalakeDiagnosticsRequestTypeDestinationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	// value enum
	if err := m.validateDestinationEnum("destination", "body", *m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateExcludeHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.ExcludeHosts) { // not required
		return nil
	}

	if err := validate.UniqueItems("excludeHosts", "body", m.ExcludeHosts); err != nil {
		return err
	}

	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateHostGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.HostGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	if err := validate.UniqueItems("hosts", "body", m.Hosts); err != nil {
		return err
	}

	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := validate.UniqueItems("labels", "body", m.Labels); err != nil {
		return err
	}

	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this collect datalake diagnostics request based on the context it is used
func (m *CollectDatalakeDiagnosticsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CollectDatalakeDiagnosticsRequest) contextValidateAdditionalLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalLogs); i++ {

		if m.AdditionalLogs[i] != nil {
			if err := m.AdditionalLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CollectDatalakeDiagnosticsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CollectDatalakeDiagnosticsRequest) UnmarshalBinary(b []byte) error {
	var res CollectDatalakeDiagnosticsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
