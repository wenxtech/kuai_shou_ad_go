/*
磁力引擎 copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
    "fmt"
)
// TemporalGranularity the model 'TemporalGranularity'
type TemporalGranularity string

// List of TemporalGranularity
const (
	TemporalGranularity_DAILY = "DAILY"
	TemporalGranularity_HOURLY = "HOURLY"
)

// All allowed values of TemporalGranularity enum
var AllowedTemporalGranularityEnumValues = []TemporalGranularity{
	"DAILY",
	"HOURLY",
}

// NewTemporalGranularityFromValue returns a pointer to a valid TemporalGranularity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTemporalGranularityFromValue(v string) (*TemporalGranularity, error) {
	ev := TemporalGranularity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TemporalGranularity: valid values are %v", v, AllowedTemporalGranularityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TemporalGranularity) IsValid() bool {
	for _, existing := range AllowedTemporalGranularityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TemporalGranularity value
func (v TemporalGranularity) Ptr() *TemporalGranularity {
	return &v
}


