# V1AsyncTaskListPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiserId** | **int64** | 广告主 ID | 
**TaskIds** | **[]int64** | 任务 ID 集，不超过 10 个 | 
**Page** | **int64** | 页码，默认 1 | 
**PageSize** | **int64** | 每页行数，默认 20 | 

## Methods

### NewV1AsyncTaskListPostRequest

`func NewV1AsyncTaskListPostRequest(advertiserId int64, taskIds []int64, page int64, pageSize int64, ) *V1AsyncTaskListPostRequest`

NewV1AsyncTaskListPostRequest instantiates a new V1AsyncTaskListPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AsyncTaskListPostRequestWithDefaults

`func NewV1AsyncTaskListPostRequestWithDefaults() *V1AsyncTaskListPostRequest`

NewV1AsyncTaskListPostRequestWithDefaults instantiates a new V1AsyncTaskListPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiserId

`func (o *V1AsyncTaskListPostRequest) GetAdvertiserId() int64`

GetAdvertiserId returns the AdvertiserId field if non-nil, zero value otherwise.

### GetAdvertiserIdOk

`func (o *V1AsyncTaskListPostRequest) GetAdvertiserIdOk() (*int64, bool)`

GetAdvertiserIdOk returns a tuple with the AdvertiserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiserId

`func (o *V1AsyncTaskListPostRequest) SetAdvertiserId(v int64)`

SetAdvertiserId sets AdvertiserId field to given value.


### GetTaskIds

`func (o *V1AsyncTaskListPostRequest) GetTaskIds() []int64`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *V1AsyncTaskListPostRequest) GetTaskIdsOk() (*[]int64, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *V1AsyncTaskListPostRequest) SetTaskIds(v []int64)`

SetTaskIds sets TaskIds field to given value.


### GetPage

`func (o *V1AsyncTaskListPostRequest) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *V1AsyncTaskListPostRequest) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *V1AsyncTaskListPostRequest) SetPage(v int64)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *V1AsyncTaskListPostRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V1AsyncTaskListPostRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V1AsyncTaskListPostRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


