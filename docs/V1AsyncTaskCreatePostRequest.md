# V1AsyncTaskCreatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiserId** | **int64** | 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回 | 
**TaskName** | **string** | 任务名称，最大不超过 50字符，不能为空；每个账户提交的任务名称不能重复 | 
**TaskParams** | [**V1AsyncTaskCreatePostRequestTaskParams**](V1AsyncTaskCreatePostRequestTaskParams.md) |  | 

## Methods

### NewV1AsyncTaskCreatePostRequest

`func NewV1AsyncTaskCreatePostRequest(advertiserId int64, taskName string, taskParams V1AsyncTaskCreatePostRequestTaskParams, ) *V1AsyncTaskCreatePostRequest`

NewV1AsyncTaskCreatePostRequest instantiates a new V1AsyncTaskCreatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AsyncTaskCreatePostRequestWithDefaults

`func NewV1AsyncTaskCreatePostRequestWithDefaults() *V1AsyncTaskCreatePostRequest`

NewV1AsyncTaskCreatePostRequestWithDefaults instantiates a new V1AsyncTaskCreatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiserId

`func (o *V1AsyncTaskCreatePostRequest) GetAdvertiserId() int64`

GetAdvertiserId returns the AdvertiserId field if non-nil, zero value otherwise.

### GetAdvertiserIdOk

`func (o *V1AsyncTaskCreatePostRequest) GetAdvertiserIdOk() (*int64, bool)`

GetAdvertiserIdOk returns a tuple with the AdvertiserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiserId

`func (o *V1AsyncTaskCreatePostRequest) SetAdvertiserId(v int64)`

SetAdvertiserId sets AdvertiserId field to given value.


### GetTaskName

`func (o *V1AsyncTaskCreatePostRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *V1AsyncTaskCreatePostRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *V1AsyncTaskCreatePostRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetTaskParams

`func (o *V1AsyncTaskCreatePostRequest) GetTaskParams() V1AsyncTaskCreatePostRequestTaskParams`

GetTaskParams returns the TaskParams field if non-nil, zero value otherwise.

### GetTaskParamsOk

`func (o *V1AsyncTaskCreatePostRequest) GetTaskParamsOk() (*V1AsyncTaskCreatePostRequestTaskParams, bool)`

GetTaskParamsOk returns a tuple with the TaskParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskParams

`func (o *V1AsyncTaskCreatePostRequest) SetTaskParams(v V1AsyncTaskCreatePostRequestTaskParams)`

SetTaskParams sets TaskParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


