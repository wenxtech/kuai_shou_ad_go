# \V1ReportCreativeReportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ReportCreativeReportPost**](V1ReportCreativeReportAPI.md#V1ReportCreativeReportPost) | **Post** /v1/report/creative_report | 广告创意数据-自定义



## V1ReportCreativeReportPost

> V1ReportCreativeReportResponse V1ReportCreativeReportPost(ctx).V1ReportCreativeReportPostRequest(v1ReportCreativeReportPostRequest).Execute()

广告创意数据-自定义



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
	v1ReportCreativeReportPostRequest := *openapiclient.NewV1ReportCreativeReportPostRequest(int64(123), "StartDate_example", "EndDate_example", int64(123), int64(123)) // V1ReportCreativeReportPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1ReportCreativeReportAPI.V1ReportCreativeReportPost(context.Background()).V1ReportCreativeReportPostRequest(v1ReportCreativeReportPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1ReportCreativeReportAPI.V1ReportCreativeReportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReportCreativeReportPost`: V1ReportCreativeReportResponse
	fmt.Fprintf(os.Stdout, "Response from `V1ReportCreativeReportAPI.V1ReportCreativeReportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReportCreativeReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ReportCreativeReportPostRequest** | [**V1ReportCreativeReportPostRequest**](V1ReportCreativeReportPostRequest.md) |  | 

### Return type

[**V1ReportCreativeReportResponse**](V1ReportCreativeReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

