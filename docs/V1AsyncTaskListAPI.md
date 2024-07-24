# \V1AsyncTaskListAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AsyncTaskListPost**](V1AsyncTaskListAPI.md#V1AsyncTaskListPost) | **Post** /v1/async_task/list | 获取任务状态



## V1AsyncTaskListPost

> V1AsyncTaskListResponse V1AsyncTaskListPost(ctx).V1AsyncTaskListPostRequest(v1AsyncTaskListPostRequest).Execute()

获取任务状态



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
	v1AsyncTaskListPostRequest := *openapiclient.NewV1AsyncTaskListPostRequest(int64(123), []int64{int64(123)}, int64(123), int64(123)) // V1AsyncTaskListPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.V1AsyncTaskListAPI.V1AsyncTaskListPost(context.Background()).V1AsyncTaskListPostRequest(v1AsyncTaskListPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `V1AsyncTaskListAPI.V1AsyncTaskListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AsyncTaskListPost`: V1AsyncTaskListResponse
	fmt.Fprintf(os.Stdout, "Response from `V1AsyncTaskListAPI.V1AsyncTaskListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AsyncTaskListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AsyncTaskListPostRequest** | [**V1AsyncTaskListPostRequest**](V1AsyncTaskListPostRequest.md) |  | 

### Return type

[**V1AsyncTaskListResponse**](V1AsyncTaskListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

