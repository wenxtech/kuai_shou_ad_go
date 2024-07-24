# V1ReportMaterialReportPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiserId** | **int64** | 广告主 ID（注：非账户快手 ID），在获取 accessToken 时返回 | 
**StartDate** | **string** | 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明 | 
**EndDate** | **string** | 过滤筛选条件，格式 yyyy-MM-dd 可选时间范围参见文档上方说明 | 
**Page** | **int64** | 请求的页码，默认为 1 | 
**PageSize** | **int64** | 每页行数，默认为 20，最大支持 1000 | 
**ViewType** | **string** | 报表类型：5 - 视频报表 7 - 封面报表 8 - 便利贴报表;16-图片报表； | 
**TemporalGranularity** | Pointer to **string** | 天粒度（DAILY）／小时粒度（HOURLY），默认支持天粒度数据 | [optional] 
**CampaignType** | Pointer to **int64** | 计划类型，过滤筛选条件 1 - 作品推广；2 - 提升应用安装；3 - 获取电商下单；4 - 推广品牌活动；5 - 收集销售线索；6 - 保量广告；7 - 提高应用活跃。 | [optional] 
**CampaignIds** | Pointer to **[]int64** | 广告计划 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**UnitIds** | Pointer to **[]int64** | 广告组 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**CreativeIds** | Pointer to **[]int64** | 广告创意 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**PhotoIds** | Pointer to **[]int64** | 视频 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**CoverIds** | Pointer to **[]int64** | 封面 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**CreativeMaterialType** | Pointer to **int64** | 素材类型：1 - 竖版视频; 2 - 横版视频; 3 - 便利贴;5 -竖版图片; 6- 横版图片; 9-小图；10-组图 | [optional] 

## Methods

### NewV1ReportMaterialReportPostRequest

`func NewV1ReportMaterialReportPostRequest(advertiserId int64, startDate string, endDate string, page int64, pageSize int64, viewType string, ) *V1ReportMaterialReportPostRequest`

NewV1ReportMaterialReportPostRequest instantiates a new V1ReportMaterialReportPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReportMaterialReportPostRequestWithDefaults

`func NewV1ReportMaterialReportPostRequestWithDefaults() *V1ReportMaterialReportPostRequest`

NewV1ReportMaterialReportPostRequestWithDefaults instantiates a new V1ReportMaterialReportPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiserId

`func (o *V1ReportMaterialReportPostRequest) GetAdvertiserId() int64`

GetAdvertiserId returns the AdvertiserId field if non-nil, zero value otherwise.

### GetAdvertiserIdOk

`func (o *V1ReportMaterialReportPostRequest) GetAdvertiserIdOk() (*int64, bool)`

GetAdvertiserIdOk returns a tuple with the AdvertiserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiserId

`func (o *V1ReportMaterialReportPostRequest) SetAdvertiserId(v int64)`

SetAdvertiserId sets AdvertiserId field to given value.


### GetStartDate

`func (o *V1ReportMaterialReportPostRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *V1ReportMaterialReportPostRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *V1ReportMaterialReportPostRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *V1ReportMaterialReportPostRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *V1ReportMaterialReportPostRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *V1ReportMaterialReportPostRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetPage

`func (o *V1ReportMaterialReportPostRequest) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *V1ReportMaterialReportPostRequest) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *V1ReportMaterialReportPostRequest) SetPage(v int64)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *V1ReportMaterialReportPostRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V1ReportMaterialReportPostRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V1ReportMaterialReportPostRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetViewType

`func (o *V1ReportMaterialReportPostRequest) GetViewType() string`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *V1ReportMaterialReportPostRequest) GetViewTypeOk() (*string, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *V1ReportMaterialReportPostRequest) SetViewType(v string)`

SetViewType sets ViewType field to given value.


### GetTemporalGranularity

`func (o *V1ReportMaterialReportPostRequest) GetTemporalGranularity() string`

GetTemporalGranularity returns the TemporalGranularity field if non-nil, zero value otherwise.

### GetTemporalGranularityOk

`func (o *V1ReportMaterialReportPostRequest) GetTemporalGranularityOk() (*string, bool)`

GetTemporalGranularityOk returns a tuple with the TemporalGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalGranularity

`func (o *V1ReportMaterialReportPostRequest) SetTemporalGranularity(v string)`

SetTemporalGranularity sets TemporalGranularity field to given value.

### HasTemporalGranularity

`func (o *V1ReportMaterialReportPostRequest) HasTemporalGranularity() bool`

HasTemporalGranularity returns a boolean if a field has been set.

### GetCampaignType

`func (o *V1ReportMaterialReportPostRequest) GetCampaignType() int64`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *V1ReportMaterialReportPostRequest) GetCampaignTypeOk() (*int64, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *V1ReportMaterialReportPostRequest) SetCampaignType(v int64)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *V1ReportMaterialReportPostRequest) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetCampaignIds

`func (o *V1ReportMaterialReportPostRequest) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *V1ReportMaterialReportPostRequest) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *V1ReportMaterialReportPostRequest) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *V1ReportMaterialReportPostRequest) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### GetUnitIds

`func (o *V1ReportMaterialReportPostRequest) GetUnitIds() []int64`

GetUnitIds returns the UnitIds field if non-nil, zero value otherwise.

### GetUnitIdsOk

`func (o *V1ReportMaterialReportPostRequest) GetUnitIdsOk() (*[]int64, bool)`

GetUnitIdsOk returns a tuple with the UnitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitIds

`func (o *V1ReportMaterialReportPostRequest) SetUnitIds(v []int64)`

SetUnitIds sets UnitIds field to given value.

### HasUnitIds

`func (o *V1ReportMaterialReportPostRequest) HasUnitIds() bool`

HasUnitIds returns a boolean if a field has been set.

### GetCreativeIds

`func (o *V1ReportMaterialReportPostRequest) GetCreativeIds() []int64`

GetCreativeIds returns the CreativeIds field if non-nil, zero value otherwise.

### GetCreativeIdsOk

`func (o *V1ReportMaterialReportPostRequest) GetCreativeIdsOk() (*[]int64, bool)`

GetCreativeIdsOk returns a tuple with the CreativeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeIds

`func (o *V1ReportMaterialReportPostRequest) SetCreativeIds(v []int64)`

SetCreativeIds sets CreativeIds field to given value.

### HasCreativeIds

`func (o *V1ReportMaterialReportPostRequest) HasCreativeIds() bool`

HasCreativeIds returns a boolean if a field has been set.

### GetPhotoIds

`func (o *V1ReportMaterialReportPostRequest) GetPhotoIds() []int64`

GetPhotoIds returns the PhotoIds field if non-nil, zero value otherwise.

### GetPhotoIdsOk

`func (o *V1ReportMaterialReportPostRequest) GetPhotoIdsOk() (*[]int64, bool)`

GetPhotoIdsOk returns a tuple with the PhotoIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoIds

`func (o *V1ReportMaterialReportPostRequest) SetPhotoIds(v []int64)`

SetPhotoIds sets PhotoIds field to given value.

### HasPhotoIds

`func (o *V1ReportMaterialReportPostRequest) HasPhotoIds() bool`

HasPhotoIds returns a boolean if a field has been set.

### GetCoverIds

`func (o *V1ReportMaterialReportPostRequest) GetCoverIds() []int64`

GetCoverIds returns the CoverIds field if non-nil, zero value otherwise.

### GetCoverIdsOk

`func (o *V1ReportMaterialReportPostRequest) GetCoverIdsOk() (*[]int64, bool)`

GetCoverIdsOk returns a tuple with the CoverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverIds

`func (o *V1ReportMaterialReportPostRequest) SetCoverIds(v []int64)`

SetCoverIds sets CoverIds field to given value.

### HasCoverIds

`func (o *V1ReportMaterialReportPostRequest) HasCoverIds() bool`

HasCoverIds returns a boolean if a field has been set.

### GetCreativeMaterialType

`func (o *V1ReportMaterialReportPostRequest) GetCreativeMaterialType() int64`

GetCreativeMaterialType returns the CreativeMaterialType field if non-nil, zero value otherwise.

### GetCreativeMaterialTypeOk

`func (o *V1ReportMaterialReportPostRequest) GetCreativeMaterialTypeOk() (*int64, bool)`

GetCreativeMaterialTypeOk returns a tuple with the CreativeMaterialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeMaterialType

`func (o *V1ReportMaterialReportPostRequest) SetCreativeMaterialType(v int64)`

SetCreativeMaterialType sets CreativeMaterialType field to given value.

### HasCreativeMaterialType

`func (o *V1ReportMaterialReportPostRequest) HasCreativeMaterialType() bool`

HasCreativeMaterialType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


