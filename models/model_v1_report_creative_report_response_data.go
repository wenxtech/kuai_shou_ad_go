/*
磁力引擎 copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// V1ReportCreativeReportResponseData struct for V1ReportCreativeReportResponseData
type V1ReportCreativeReportResponseData struct {
	Empty *bool `json:"empty,omitempty"`
	TotalCount *int64 `json:"total_count,omitempty"`
	Details []V1ReportCreativeReportResponseDataDetailsInner `json:"details,omitempty"`
}


