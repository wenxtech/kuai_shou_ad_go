# V1AsyncTaskListResponseDataDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiserId** | **int64** | 广告主 ID | 
**TaskId** | **int64** | 任务 ID | 
**TaskName** | **string** | 任务名称 | 
**CreateTime** | **string** | 任务创建时间 格式如：yyyy-MM-dd HH:mm:ss | 
**TaskStatus** | **int64** | 任务状态 0：新建，1：处理中，2：处理成功，3：处理失败 | 
**FileSize** | **int64** | 文件大小 字节数 | 

## Methods

### NewV1AsyncTaskListResponseDataDetailsInner

`func NewV1AsyncTaskListResponseDataDetailsInner(advertiserId int64, taskId int64, taskName string, createTime string, taskStatus int64, fileSize int64, ) *V1AsyncTaskListResponseDataDetailsInner`

NewV1AsyncTaskListResponseDataDetailsInner instantiates a new V1AsyncTaskListResponseDataDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AsyncTaskListResponseDataDetailsInnerWithDefaults

`func NewV1AsyncTaskListResponseDataDetailsInnerWithDefaults() *V1AsyncTaskListResponseDataDetailsInner`

NewV1AsyncTaskListResponseDataDetailsInnerWithDefaults instantiates a new V1AsyncTaskListResponseDataDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiserId

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetAdvertiserId() int64`

GetAdvertiserId returns the AdvertiserId field if non-nil, zero value otherwise.

### GetAdvertiserIdOk

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetAdvertiserIdOk() (*int64, bool)`

GetAdvertiserIdOk returns a tuple with the AdvertiserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiserId

`func (o *V1AsyncTaskListResponseDataDetailsInner) SetAdvertiserId(v int64)`

SetAdvertiserId sets AdvertiserId field to given value.


### GetTaskId

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetTaskId() int64`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetTaskIdOk() (*int64, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *V1AsyncTaskListResponseDataDetailsInner) SetTaskId(v int64)`

SetTaskId sets TaskId field to given value.


### GetTaskName

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *V1AsyncTaskListResponseDataDetailsInner) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetCreateTime

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *V1AsyncTaskListResponseDataDetailsInner) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.


### GetTaskStatus

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetTaskStatus() int64`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetTaskStatusOk() (*int64, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *V1AsyncTaskListResponseDataDetailsInner) SetTaskStatus(v int64)`

SetTaskStatus sets TaskStatus field to given value.


### GetFileSize

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *V1AsyncTaskListResponseDataDetailsInner) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *V1AsyncTaskListResponseDataDetailsInner) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


