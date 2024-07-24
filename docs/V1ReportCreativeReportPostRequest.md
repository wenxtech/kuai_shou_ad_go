# V1ReportCreativeReportPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiserId** | **int64** | 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回 | 
**StartDate** | **string** | 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明 | 
**EndDate** | **string** | 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明 | 
**Page** | **int64** | 请求的页码，默认为 1 | 
**PageSize** | **int64** | 每页行数，默认为 20，最大支持 2000 | 
**TemporalGranularity** | Pointer to **string** | 天粒度（DAILY）／小时粒度（HOURLY），默认支持天粒度数据 | [optional] 
**CampaignType** | Pointer to **int64** | 计划类型，过滤筛选条件 1 - 作品推广；2 - 提升应用安装；3 - 获取电商下单；4 - 推广品牌活动；5 - 收集销售线索；6 - 保量广告；7 - 提高应用活跃。 | [optional] 
**CampaignIds** | Pointer to **[]int64** | 广告计划 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**UnitIds** | Pointer to **[]int64** | 广告组 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**CreativeIds** | Pointer to **[]int64** | 广告创意 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**ReportDims** | Pointer to **[]string** | \&quot;adScene\&quot;：按广告场景；激励视频广告/联盟/上下滑大屏广告/信息流广告/开屏广告/搜索广告/视频播放页广告/全部 \&quot;placementType\&quot;：按广告范围，快手/联盟;不传/传空/传空数组：不限 | [optional] 
**ExtendInfo** | Pointer to **[]string** | 获取额外信息， \&quot;photo\&quot;获取视频信息，包含视频Id与视频md5 | [optional] 

## Methods

### NewV1ReportCreativeReportPostRequest

`func NewV1ReportCreativeReportPostRequest(advertiserId int64, startDate string, endDate string, page int64, pageSize int64, ) *V1ReportCreativeReportPostRequest`

NewV1ReportCreativeReportPostRequest instantiates a new V1ReportCreativeReportPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReportCreativeReportPostRequestWithDefaults

`func NewV1ReportCreativeReportPostRequestWithDefaults() *V1ReportCreativeReportPostRequest`

NewV1ReportCreativeReportPostRequestWithDefaults instantiates a new V1ReportCreativeReportPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiserId

`func (o *V1ReportCreativeReportPostRequest) GetAdvertiserId() int64`

GetAdvertiserId returns the AdvertiserId field if non-nil, zero value otherwise.

### GetAdvertiserIdOk

`func (o *V1ReportCreativeReportPostRequest) GetAdvertiserIdOk() (*int64, bool)`

GetAdvertiserIdOk returns a tuple with the AdvertiserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiserId

`func (o *V1ReportCreativeReportPostRequest) SetAdvertiserId(v int64)`

SetAdvertiserId sets AdvertiserId field to given value.


### GetStartDate

`func (o *V1ReportCreativeReportPostRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *V1ReportCreativeReportPostRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *V1ReportCreativeReportPostRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *V1ReportCreativeReportPostRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *V1ReportCreativeReportPostRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *V1ReportCreativeReportPostRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetPage

`func (o *V1ReportCreativeReportPostRequest) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *V1ReportCreativeReportPostRequest) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *V1ReportCreativeReportPostRequest) SetPage(v int64)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *V1ReportCreativeReportPostRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V1ReportCreativeReportPostRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V1ReportCreativeReportPostRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetTemporalGranularity

`func (o *V1ReportCreativeReportPostRequest) GetTemporalGranularity() string`

GetTemporalGranularity returns the TemporalGranularity field if non-nil, zero value otherwise.

### GetTemporalGranularityOk

`func (o *V1ReportCreativeReportPostRequest) GetTemporalGranularityOk() (*string, bool)`

GetTemporalGranularityOk returns a tuple with the TemporalGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalGranularity

`func (o *V1ReportCreativeReportPostRequest) SetTemporalGranularity(v string)`

SetTemporalGranularity sets TemporalGranularity field to given value.

### HasTemporalGranularity

`func (o *V1ReportCreativeReportPostRequest) HasTemporalGranularity() bool`

HasTemporalGranularity returns a boolean if a field has been set.

### GetCampaignType

`func (o *V1ReportCreativeReportPostRequest) GetCampaignType() int64`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *V1ReportCreativeReportPostRequest) GetCampaignTypeOk() (*int64, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *V1ReportCreativeReportPostRequest) SetCampaignType(v int64)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *V1ReportCreativeReportPostRequest) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetCampaignIds

`func (o *V1ReportCreativeReportPostRequest) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *V1ReportCreativeReportPostRequest) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *V1ReportCreativeReportPostRequest) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *V1ReportCreativeReportPostRequest) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### GetUnitIds

`func (o *V1ReportCreativeReportPostRequest) GetUnitIds() []int64`

GetUnitIds returns the UnitIds field if non-nil, zero value otherwise.

### GetUnitIdsOk

`func (o *V1ReportCreativeReportPostRequest) GetUnitIdsOk() (*[]int64, bool)`

GetUnitIdsOk returns a tuple with the UnitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitIds

`func (o *V1ReportCreativeReportPostRequest) SetUnitIds(v []int64)`

SetUnitIds sets UnitIds field to given value.

### HasUnitIds

`func (o *V1ReportCreativeReportPostRequest) HasUnitIds() bool`

HasUnitIds returns a boolean if a field has been set.

### GetCreativeIds

`func (o *V1ReportCreativeReportPostRequest) GetCreativeIds() []int64`

GetCreativeIds returns the CreativeIds field if non-nil, zero value otherwise.

### GetCreativeIdsOk

`func (o *V1ReportCreativeReportPostRequest) GetCreativeIdsOk() (*[]int64, bool)`

GetCreativeIdsOk returns a tuple with the CreativeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeIds

`func (o *V1ReportCreativeReportPostRequest) SetCreativeIds(v []int64)`

SetCreativeIds sets CreativeIds field to given value.

### HasCreativeIds

`func (o *V1ReportCreativeReportPostRequest) HasCreativeIds() bool`

HasCreativeIds returns a boolean if a field has been set.

### GetReportDims

`func (o *V1ReportCreativeReportPostRequest) GetReportDims() []string`

GetReportDims returns the ReportDims field if non-nil, zero value otherwise.

### GetReportDimsOk

`func (o *V1ReportCreativeReportPostRequest) GetReportDimsOk() (*[]string, bool)`

GetReportDimsOk returns a tuple with the ReportDims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDims

`func (o *V1ReportCreativeReportPostRequest) SetReportDims(v []string)`

SetReportDims sets ReportDims field to given value.

### HasReportDims

`func (o *V1ReportCreativeReportPostRequest) HasReportDims() bool`

HasReportDims returns a boolean if a field has been set.

### GetExtendInfo

`func (o *V1ReportCreativeReportPostRequest) GetExtendInfo() []string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *V1ReportCreativeReportPostRequest) GetExtendInfoOk() (*[]string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *V1ReportCreativeReportPostRequest) SetExtendInfo(v []string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *V1ReportCreativeReportPostRequest) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


