# V1AsyncTaskDownloadGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiserId** | **string** | 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回 | 
**TaskId** | **string** | 任务 ID | 

## Methods

### NewV1AsyncTaskDownloadGetRequest

`func NewV1AsyncTaskDownloadGetRequest(advertiserId string, taskId string, ) *V1AsyncTaskDownloadGetRequest`

NewV1AsyncTaskDownloadGetRequest instantiates a new V1AsyncTaskDownloadGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AsyncTaskDownloadGetRequestWithDefaults

`func NewV1AsyncTaskDownloadGetRequestWithDefaults() *V1AsyncTaskDownloadGetRequest`

NewV1AsyncTaskDownloadGetRequestWithDefaults instantiates a new V1AsyncTaskDownloadGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiserId

`func (o *V1AsyncTaskDownloadGetRequest) GetAdvertiserId() string`

GetAdvertiserId returns the AdvertiserId field if non-nil, zero value otherwise.

### GetAdvertiserIdOk

`func (o *V1AsyncTaskDownloadGetRequest) GetAdvertiserIdOk() (*string, bool)`

GetAdvertiserIdOk returns a tuple with the AdvertiserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiserId

`func (o *V1AsyncTaskDownloadGetRequest) SetAdvertiserId(v string)`

SetAdvertiserId sets AdvertiserId field to given value.


### GetTaskId

`func (o *V1AsyncTaskDownloadGetRequest) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *V1AsyncTaskDownloadGetRequest) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *V1AsyncTaskDownloadGetRequest) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


