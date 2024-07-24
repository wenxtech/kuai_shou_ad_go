/*
磁力引擎 copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// V1AsyncTaskCreatePostRequest struct for V1AsyncTaskCreatePostRequest
type V1AsyncTaskCreatePostRequest struct {
	// 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回
	AdvertiserId int64 `json:"advertiser_id"`
	// 任务名称，最大不超过 50字符，不能为空；每个账户提交的任务名称不能重复
	TaskName string `json:"task_name"`
	TaskParams V1AsyncTaskCreatePostRequestTaskParams `json:"task_params"`
}

type _V1AsyncTaskCreatePostRequest V1AsyncTaskCreatePostRequest


