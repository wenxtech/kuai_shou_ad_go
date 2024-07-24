# \V1AsyncTaskDownloadAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AsyncTaskDownloadGet**](V1AsyncTaskDownloadAPI.md#V1AsyncTaskDownloadGet) | **Get** /v1/async_task/download | 数据下载



## V1AsyncTaskDownloadGet

> map[string]interface{} V1AsyncTaskDownloadGet(ctx).V1AsyncTaskDownloadGetRequest(v1AsyncTaskDownloadGetRequest).Execute()

数据下载



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
	v1AsyncTaskDownloadGetRequest := *openapiclient.NewV1AsyncTaskDownloadGetRequest("AdvertiserId_example", "TaskId_example") // V1AsyncTaskDownloadGetRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1AsyncTaskDownloadAPI.V1AsyncTaskDownloadGet(context.Background()).V1AsyncTaskDownloadGetRequest(v1AsyncTaskDownloadGetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1AsyncTaskDownloadAPI.V1AsyncTaskDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AsyncTaskDownloadGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `V1AsyncTaskDownloadAPI.V1AsyncTaskDownloadGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AsyncTaskDownloadGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AsyncTaskDownloadGetRequest** | [**V1AsyncTaskDownloadGetRequest**](V1AsyncTaskDownloadGetRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

