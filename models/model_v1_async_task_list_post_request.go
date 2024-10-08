/*
磁力引擎 copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// V1AsyncTaskListPostRequest struct for V1AsyncTaskListPostRequest
type V1AsyncTaskListPostRequest struct {
	// 广告主 ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 任务 ID 集，不超过 10 个
	TaskIds []int64 `json:"task_ids"`
	// 页码，默认 1
	Page int64 `json:"page"`
	// 每页行数，默认 20
	PageSize int64 `json:"page_size"`
}

type _V1AsyncTaskListPostRequest V1AsyncTaskListPostRequest


