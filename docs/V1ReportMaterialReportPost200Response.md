# V1ReportMaterialReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** |  | [optional] 
**Details** | Pointer to [**[]V1ReportMaterialReportResponseDetailsInner**](V1ReportMaterialReportResponseDetailsInner.md) |  | [optional] 

## Methods

### NewV1ReportMaterialReportResponse

`func NewV1ReportMaterialReportResponse() *V1ReportMaterialReportResponse`

NewV1ReportMaterialReportResponse instantiates a new V1ReportMaterialReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReportMaterialReportResponseWithDefaults

`func NewV1ReportMaterialReportResponseWithDefaults() *V1ReportMaterialReportResponse`

NewV1ReportMaterialReportResponseWithDefaults instantiates a new V1ReportMaterialReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *V1ReportMaterialReportResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V1ReportMaterialReportResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V1ReportMaterialReportResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *V1ReportMaterialReportResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetDetails

`func (o *V1ReportMaterialReportResponse) GetDetails() []V1ReportMaterialReportResponseDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *V1ReportMaterialReportResponse) GetDetailsOk() (*[]V1ReportMaterialReportResponseDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *V1ReportMaterialReportResponse) SetDetails(v []V1ReportMaterialReportResponseDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *V1ReportMaterialReportResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


