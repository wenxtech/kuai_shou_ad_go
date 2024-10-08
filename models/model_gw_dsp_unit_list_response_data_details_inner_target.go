/*
磁力引擎 copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// GwDspUnitListResponseDataDetailsInnerTarget struct for GwDspUnitListResponseDataDetailsInnerTarget
type GwDspUnitListResponseDataDetailsInnerTarget struct {
	Gender *int64 `json:"gender,omitempty"`
	FilterConvertedWechatId []string `json:"filter_converted_wechat_id,omitempty"`
	Media []string `json:"media,omitempty"`
	Network *int64 `json:"network,omitempty"`
	DevicePrice []string `json:"device_price,omitempty"`
	MediaSourceType *int64 `json:"media_source_type,omitempty"`
	TargetSource *int64 `json:"target_source,omitempty"`
	UserType *int64 `json:"user_type,omitempty"`
	Operators []string `json:"operators,omitempty"`
	DistanceShow []string `json:"distance_show,omitempty"`
	ExcludePopulation []int64 `json:"exclude_population,omitempty"`
	DeviceBrandIds []int64 `json:"device_brand_ids,omitempty"`
	AgesRange []int64 `json:"ages_range,omitempty"`
	DistrictIds []string `json:"district_ids,omitempty"`
	AndroidOsv *int64 `json:"android_osv,omitempty"`
	BehaviorType *int64 `json:"behavior_type,omitempty"`
	DisableInstalledAppSwitch *int64 `json:"disable_installed_app_switch,omitempty"`
	AppIds []int64 `json:"app_ids,omitempty"`
	AppNames []string `json:"app_names,omitempty"`
	IpType *int64 `json:"ip_type,omitempty"`
	FilterConvertedLevel *int64 `json:"filter_converted_level,omitempty"`
	Population []int64 `json:"population,omitempty"`
	PlatformOs *int64 `json:"platform_os,omitempty"`
	FilterTimeRange *int64 `json:"filter_time_range,omitempty"`
	BusinessInterest *string `json:"business_interest,omitempty"`
	SeedPopulation []string `json:"seed_population,omitempty"`
	IosOsv *int64 `json:"ios_osv,omitempty"`
	AutoPopulation *int64 `json:"auto_population,omitempty"`
	AppInterestIds []string `json:"app_interest_ids,omitempty"`
	PaidAudience []string `json:"paid_audience,omitempty"`
	IntelliExtendOption *int64 `json:"intelli_extend_option,omitempty"`
	Region []int64 `json:"region,omitempty"`
	ExcludeMedia []string `json:"exclude_media,omitempty"`
}


