# \V1ReportMaterialReportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ReportMaterialReportPost**](V1ReportMaterialReportAPI.md#V1ReportMaterialReportPost) | **Post** /v1/report/material_report | 素材报表



## V1ReportMaterialReportPost

> V1ReportMaterialReportResponse V1ReportMaterialReportPost(ctx).V1ReportMaterialReportPostRequest(v1ReportMaterialReportPostRequest).Execute()

素材报表



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
	v1ReportMaterialReportPostRequest := *openapiclient.NewV1ReportMaterialReportPostRequest(int64(123), "StartDate_example", "EndDate_example", int64(123), int64(123), "ViewType_example") // V1ReportMaterialReportPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1ReportMaterialReportAPI.V1ReportMaterialReportPost(context.Background()).V1ReportMaterialReportPostRequest(v1ReportMaterialReportPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1ReportMaterialReportAPI.V1ReportMaterialReportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReportMaterialReportPost`: V1ReportMaterialReportResponse
	fmt.Fprintf(os.Stdout, "Response from `V1ReportMaterialReportAPI.V1ReportMaterialReportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReportMaterialReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ReportMaterialReportPostRequest** | [**V1ReportMaterialReportPostRequest**](V1ReportMaterialReportPostRequest.md) |  | 

### Return type

[**V1ReportMaterialReportResponse**](V1ReportMaterialReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

