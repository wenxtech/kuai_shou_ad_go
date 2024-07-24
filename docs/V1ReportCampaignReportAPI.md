# \V1ReportCampaignReportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ReportCampaignReportPost**](V1ReportCampaignReportAPI.md#V1ReportCampaignReportPost) | **Post** /v1/report/campaign_report | 广告计划数据



## V1ReportCampaignReportPost

> V1ReportCampaignReportResponse V1ReportCampaignReportPost(ctx).V1ReportCampaignReportPostRequest(v1ReportCampaignReportPostRequest).Execute()

广告计划数据



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
	v1ReportCampaignReportPostRequest := *openapiclient.NewV1ReportCampaignReportPostRequest(int64(123), "StartDate_example", "EndDate_example", int64(123), int64(123)) // V1ReportCampaignReportPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1ReportCampaignReportAPI.V1ReportCampaignReportPost(context.Background()).V1ReportCampaignReportPostRequest(v1ReportCampaignReportPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1ReportCampaignReportAPI.V1ReportCampaignReportPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ReportCampaignReportPost`: V1ReportCampaignReportResponse
	fmt.Fprintf(os.Stdout, "Response from `V1ReportCampaignReportAPI.V1ReportCampaignReportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ReportCampaignReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ReportCampaignReportPostRequest** | [**V1ReportCampaignReportPostRequest**](V1ReportCampaignReportPostRequest.md) |  | 

### Return type

[**V1ReportCampaignReportResponse**](V1ReportCampaignReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

