# V1AsyncTaskListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int64** |  | 
**Message** | **string** |  | 
**Data** | [**V1AsyncTaskListResponseData**](V1AsyncTaskListResponseData.md) |  | 

## Methods

### NewV1AsyncTaskListResponse

`func NewV1AsyncTaskListResponse(code int64, message string, data V1AsyncTaskListResponseData, ) *V1AsyncTaskListResponse`

NewV1AsyncTaskListResponse instantiates a new V1AsyncTaskListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AsyncTaskListResponseWithDefaults

`func NewV1AsyncTaskListResponseWithDefaults() *V1AsyncTaskListResponse`

NewV1AsyncTaskListResponseWithDefaults instantiates a new V1AsyncTaskListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *V1AsyncTaskListResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *V1AsyncTaskListResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *V1AsyncTaskListResponse) SetCode(v int64)`

SetCode sets Code field to given value.


### GetMessage

`func (o *V1AsyncTaskListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1AsyncTaskListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1AsyncTaskListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *V1AsyncTaskListResponse) GetData() V1AsyncTaskListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1AsyncTaskListResponse) GetDataOk() (*V1AsyncTaskListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1AsyncTaskListResponse) SetData(v V1AsyncTaskListResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


