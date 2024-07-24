# \V1AsyncTaskCreateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AsyncTaskCreatePost**](V1AsyncTaskCreateAPI.md#V1AsyncTaskCreatePost) | **Post** /v1/async_task/create | 创建历史数据查询任务



## V1AsyncTaskCreatePost

> V1AsyncTaskCreateResponse V1AsyncTaskCreatePost(ctx).V1AsyncTaskCreatePostRequest(v1AsyncTaskCreatePostRequest).Execute()

创建历史数据查询任务



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
	v1AsyncTaskCreatePostRequest := *openapiclient.NewV1AsyncTaskCreatePostRequest(int64(123), "TaskName_example", *openapiclient.NewV1AsyncTaskCreatePostRequestTaskParams("StartDate_example", "EndDate_example", int64(123), []string{"Query_example"})) // V1AsyncTaskCreatePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1AsyncTaskCreateAPI.V1AsyncTaskCreatePost(context.Background()).V1AsyncTaskCreatePostRequest(v1AsyncTaskCreatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1AsyncTaskCreateAPI.V1AsyncTaskCreatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AsyncTaskCreatePost`: V1AsyncTaskCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `V1AsyncTaskCreateAPI.V1AsyncTaskCreatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AsyncTaskCreatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AsyncTaskCreatePostRequest** | [**V1AsyncTaskCreatePostRequest**](V1AsyncTaskCreatePostRequest.md) |  | 

### Return type

[**V1AsyncTaskCreateResponse**](V1AsyncTaskCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

