# V1ReportAccountReportPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiserId** | **int64** | 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回 | 
**StartDate** | **string** | 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明 | 
**EndDate** | **string** | 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明 | 
**Page** | **int64** | 请求的页码，默认为 1 | 
**PageSize** | **int64** | 每页行数，默认为 20，最大支持 2000 | 
**TemporalGranularity** | Pointer to **string** | 天粒度（DAILY）／小时粒度（HOURLY），默认支持天粒度数据 | [optional] 

## Methods

### NewV1ReportAccountReportPostRequest

`func NewV1ReportAccountReportPostRequest(advertiserId int64, startDate string, endDate string, page int64, pageSize int64, ) *V1ReportAccountReportPostRequest`

NewV1ReportAccountReportPostRequest instantiates a new V1ReportAccountReportPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReportAccountReportPostRequestWithDefaults

`func NewV1ReportAccountReportPostRequestWithDefaults() *V1ReportAccountReportPostRequest`

NewV1ReportAccountReportPostRequestWithDefaults instantiates a new V1ReportAccountReportPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiserId

`func (o *V1ReportAccountReportPostRequest) GetAdvertiserId() int64`

GetAdvertiserId returns the AdvertiserId field if non-nil, zero value otherwise.

### GetAdvertiserIdOk

`func (o *V1ReportAccountReportPostRequest) GetAdvertiserIdOk() (*int64, bool)`

GetAdvertiserIdOk returns a tuple with the AdvertiserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiserId

`func (o *V1ReportAccountReportPostRequest) SetAdvertiserId(v int64)`

SetAdvertiserId sets AdvertiserId field to given value.


### GetStartDate

`func (o *V1ReportAccountReportPostRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *V1ReportAccountReportPostRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *V1ReportAccountReportPostRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *V1ReportAccountReportPostRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *V1ReportAccountReportPostRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *V1ReportAccountReportPostRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetPage

`func (o *V1ReportAccountReportPostRequest) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *V1ReportAccountReportPostRequest) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *V1ReportAccountReportPostRequest) SetPage(v int64)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *V1ReportAccountReportPostRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V1ReportAccountReportPostRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V1ReportAccountReportPostRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetTemporalGranularity

`func (o *V1ReportAccountReportPostRequest) GetTemporalGranularity() string`

GetTemporalGranularity returns the TemporalGranularity field if non-nil, zero value otherwise.

### GetTemporalGranularityOk

`func (o *V1ReportAccountReportPostRequest) GetTemporalGranularityOk() (*string, bool)`

GetTemporalGranularityOk returns a tuple with the TemporalGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalGranularity

`func (o *V1ReportAccountReportPostRequest) SetTemporalGranularity(v string)`

SetTemporalGranularity sets TemporalGranularity field to given value.

### HasTemporalGranularity

`func (o *V1ReportAccountReportPostRequest) HasTemporalGranularity() bool`

HasTemporalGranularity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


