# V1AsyncTaskListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | 任务总数 | 
**Details** | [**[]V1AsyncTaskListResponseDataDetailsInner**](V1AsyncTaskListResponseDataDetailsInner.md) | 详细 | 

## Methods

### NewV1AsyncTaskListResponseData

`func NewV1AsyncTaskListResponseData(totalCount int64, details []V1AsyncTaskListResponseDataDetailsInner, ) *V1AsyncTaskListResponseData`

NewV1AsyncTaskListResponseData instantiates a new V1AsyncTaskListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AsyncTaskListResponseDataWithDefaults

`func NewV1AsyncTaskListResponseDataWithDefaults() *V1AsyncTaskListResponseData`

NewV1AsyncTaskListResponseDataWithDefaults instantiates a new V1AsyncTaskListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *V1AsyncTaskListResponseData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V1AsyncTaskListResponseData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V1AsyncTaskListResponseData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetDetails

`func (o *V1AsyncTaskListResponseData) GetDetails() []V1AsyncTaskListResponseDataDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *V1AsyncTaskListResponseData) GetDetailsOk() (*[]V1AsyncTaskListResponseDataDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *V1AsyncTaskListResponseData) SetDetails(v []V1AsyncTaskListResponseDataDetailsInner)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


