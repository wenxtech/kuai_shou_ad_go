# \V1ReportUnitReportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ReportUnitReportPost**](V1ReportUnitReportAPI.md#V1ReportUnitReportPost) | **Post** /v1/report/unit_report | 广告组数据



## V1ReportUnitReportPost

> V1ReportUnitReportResponse V1ReportUnitReportPost(ctx).V1ReportUnitReportPostRequest(v1ReportUnitReportPostRequest).Execute()

广告组数据



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wenxtech/kuai_shou_ad_go"
)

func main() {
	v1ReportUnitReportPostRequest := *openapiclient.NewV1ReportUnitReportPostRequest(int64(123), "StartDate_example", "EndDate_example", int64(123), int64(123)) // V1ReportUnitReportPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1ReportUnitReportAPI.V1ReportUnitReportPost(context.Background()).V1ReportUnitReportPostRequest(v1ReportUnitReportPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1ReportUnitReportAPI.V1ReportUnitReportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReportUnitReportPost`: V1ReportUnitReportResponse
	fmt.Fprintf(os.Stdout, "Response from `V1ReportUnitReportAPI.V1ReportUnitReportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReportUnitReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ReportUnitReportPostRequest** | [**V1ReportUnitReportPostRequest**](V1ReportUnitReportPostRequest.md) |  | 

### Return type

[**V1ReportUnitReportResponse**](V1ReportUnitReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

