# \V1ReportAccountReportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ReportAccountReportPost**](V1ReportAccountReportAPI.md#V1ReportAccountReportPost) | **Post** /v1/report/account_report | 广告主数据



## V1ReportAccountReportPost

> V1ReportAccountReportResponse V1ReportAccountReportPost(ctx).V1ReportAccountReportPostRequest(v1ReportAccountReportPostRequest).Execute()

广告主数据



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
	v1ReportAccountReportPostRequest := *openapiclient.NewV1ReportAccountReportPostRequest(int64(123), "StartDate_example", "EndDate_example", int64(123), int64(123)) // V1ReportAccountReportPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1ReportAccountReportAPI.V1ReportAccountReportPost(context.Background()).V1ReportAccountReportPostRequest(v1ReportAccountReportPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1ReportAccountReportAPI.V1ReportAccountReportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReportAccountReportPost`: V1ReportAccountReportResponse
	fmt.Fprintf(os.Stdout, "Response from `V1ReportAccountReportAPI.V1ReportAccountReportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReportAccountReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ReportAccountReportPostRequest** | [**V1ReportAccountReportPostRequest**](V1ReportAccountReportPostRequest.md) |  | 

### Return type

[**V1ReportAccountReportResponse**](V1ReportAccountReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

