# V1AsyncTaskCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int64** |  | 
**Message** | **string** |  | 
**Data** | [**V1AsyncTaskCreateResponseData**](V1AsyncTaskCreateResponseData.md) |  | 
**RequestId** | **string** |  | 

## Methods

### NewV1AsyncTaskCreateResponse

`func NewV1AsyncTaskCreateResponse(code int64, message string, data V1AsyncTaskCreateResponseData, requestId string, ) *V1AsyncTaskCreateResponse`

NewV1AsyncTaskCreateResponse instantiates a new V1AsyncTaskCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AsyncTaskCreateResponseWithDefaults

`func NewV1AsyncTaskCreateResponseWithDefaults() *V1AsyncTaskCreateResponse`

NewV1AsyncTaskCreateResponseWithDefaults instantiates a new V1AsyncTaskCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *V1AsyncTaskCreateResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *V1AsyncTaskCreateResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *V1AsyncTaskCreateResponse) SetCode(v int64)`

SetCode sets Code field to given value.


### GetMessage

`func (o *V1AsyncTaskCreateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1AsyncTaskCreateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1AsyncTaskCreateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *V1AsyncTaskCreateResponse) GetData() V1AsyncTaskCreateResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1AsyncTaskCreateResponse) GetDataOk() (*V1AsyncTaskCreateResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1AsyncTaskCreateResponse) SetData(v V1AsyncTaskCreateResponseData)`

SetData sets Data field to given value.


### GetRequestId

`func (o *V1AsyncTaskCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *V1AsyncTaskCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *V1AsyncTaskCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


