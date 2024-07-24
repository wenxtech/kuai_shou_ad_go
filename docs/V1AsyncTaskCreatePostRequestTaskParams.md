# V1AsyncTaskCreatePostRequestTaskParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** | 格式如：yyyy-MM-dd，时间跨度不能超过 6 个月 | 
**EndDate** | **string** | 格式如：yyyy-MM-dd，不大于查询开始日期 | 
**ViewType** | **int64** | 1：账户维度查询，2: 广告计划维度查询，3：广告组维度查询，4：广告创意维度查询(自定义创意) 5：视频报表 报表 7：封面报表 8：便利贴报表 10：程序化创意 2.0 报表 21：关键词报表 25：搜索词报表 | 
**Query** | **[]string** | 推广关键词ID集，过滤筛选条件，单次查询数量不超过 5000,推广关键词ID集可通过获取关键词列表接口获取 | 
**TemporalGranularity** | Pointer to **string** | “DAILY”：天粒度；“HOURLY”：小时粒度；默认按天粒度进行聚合 | [optional] 
**CampaignIds** | Pointer to **[]int64** | 广告计划 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**UnitIds** | Pointer to **[]int64** | 广告组 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**CreativeIds** | Pointer to **[]int64** | 广告创意 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**PhotoIds** | Pointer to **[]int64** | 视频 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 
**CoverIds** | Pointer to **[]int64** | 封面 ID 集，过滤筛选条件，单次查询数量不超过 5000 | [optional] 

## Methods

### NewV1AsyncTaskCreatePostRequestTaskParams

`func NewV1AsyncTaskCreatePostRequestTaskParams(startDate string, endDate string, viewType int64, query []string, ) *V1AsyncTaskCreatePostRequestTaskParams`

NewV1AsyncTaskCreatePostRequestTaskParams instantiates a new V1AsyncTaskCreatePostRequestTaskParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AsyncTaskCreatePostRequestTaskParamsWithDefaults

`func NewV1AsyncTaskCreatePostRequestTaskParamsWithDefaults() *V1AsyncTaskCreatePostRequestTaskParams`

NewV1AsyncTaskCreatePostRequestTaskParamsWithDefaults instantiates a new V1AsyncTaskCreatePostRequestTaskParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetViewType

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetViewType() int64`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetViewTypeOk() (*int64, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetViewType(v int64)`

SetViewType sets ViewType field to given value.


### GetQuery

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetQuery() []string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetQueryOk() (*[]string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetQuery(v []string)`

SetQuery sets Query field to given value.


### GetTemporalGranularity

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetTemporalGranularity() string`

GetTemporalGranularity returns the TemporalGranularity field if non-nil, zero value otherwise.

### GetTemporalGranularityOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetTemporalGranularityOk() (*string, bool)`

GetTemporalGranularityOk returns a tuple with the TemporalGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporalGranularity

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetTemporalGranularity(v string)`

SetTemporalGranularity sets TemporalGranularity field to given value.

### HasTemporalGranularity

`func (o *V1AsyncTaskCreatePostRequestTaskParams) HasTemporalGranularity() bool`

HasTemporalGranularity returns a boolean if a field has been set.

### GetCampaignIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.

### HasCampaignIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### GetUnitIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetUnitIds() []int64`

GetUnitIds returns the UnitIds field if non-nil, zero value otherwise.

### GetUnitIdsOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetUnitIdsOk() (*[]int64, bool)`

GetUnitIdsOk returns a tuple with the UnitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetUnitIds(v []int64)`

SetUnitIds sets UnitIds field to given value.

### HasUnitIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) HasUnitIds() bool`

HasUnitIds returns a boolean if a field has been set.

### GetCreativeIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetCreativeIds() []int64`

GetCreativeIds returns the CreativeIds field if non-nil, zero value otherwise.

### GetCreativeIdsOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetCreativeIdsOk() (*[]int64, bool)`

GetCreativeIdsOk returns a tuple with the CreativeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetCreativeIds(v []int64)`

SetCreativeIds sets CreativeIds field to given value.

### HasCreativeIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) HasCreativeIds() bool`

HasCreativeIds returns a boolean if a field has been set.

### GetPhotoIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetPhotoIds() []int64`

GetPhotoIds returns the PhotoIds field if non-nil, zero value otherwise.

### GetPhotoIdsOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetPhotoIdsOk() (*[]int64, bool)`

GetPhotoIdsOk returns a tuple with the PhotoIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetPhotoIds(v []int64)`

SetPhotoIds sets PhotoIds field to given value.

### HasPhotoIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) HasPhotoIds() bool`

HasPhotoIds returns a boolean if a field has been set.

### GetCoverIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetCoverIds() []int64`

GetCoverIds returns the CoverIds field if non-nil, zero value otherwise.

### GetCoverIdsOk

`func (o *V1AsyncTaskCreatePostRequestTaskParams) GetCoverIdsOk() (*[]int64, bool)`

GetCoverIdsOk returns a tuple with the CoverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) SetCoverIds(v []int64)`

SetCoverIds sets CoverIds field to given value.

### HasCoverIds

`func (o *V1AsyncTaskCreatePostRequestTaskParams) HasCoverIds() bool`

HasCoverIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


