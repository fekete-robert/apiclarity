// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// APIType Api type
//
// swagger:model ApiType
type APIType string

func NewAPIType(value APIType) *APIType {
	v := value
	return &v
}

const (

	// APITypeINTERNAL captures enum value "INTERNAL"
	APITypeINTERNAL APIType = "INTERNAL"

	// APITypeEXTERNAL captures enum value "EXTERNAL"
	APITypeEXTERNAL APIType = "EXTERNAL"
)

// for schema
var apiTypeEnum []interface{}

func init() {
	var res []APIType
	if err := json.Unmarshal([]byte(`["INTERNAL","EXTERNAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiTypeEnum = append(apiTypeEnum, v)
	}
}

func (m APIType) validateAPITypeEnum(path, location string, value APIType) error {
	if err := validate.EnumCase(path, location, value, apiTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Api type
func (m APIType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAPITypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Api type based on context it is used
func (m APIType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}