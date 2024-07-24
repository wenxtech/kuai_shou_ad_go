# \Oauth2AuthorizeAccessTokenAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2AuthorizeAccessTokenPost**](Oauth2AuthorizeAccessTokenAPI.md#Oauth2AuthorizeAccessTokenPost) | **Post** /oauth2/authorize/access_token | 获取 token



## Oauth2AuthorizeAccessTokenPost

> Oauth2AuthorizeAccessTokenResponse Oauth2AuthorizeAccessTokenPost(ctx).Oauth2AuthorizeAccessTokenPostRequest(oauth2AuthorizeAccessTokenPostRequest).Execute()

获取 token



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
	oauth2AuthorizeAccessTokenPostRequest := *openapiclient.NewOauth2AuthorizeAccessTokenPostRequest(int64(123), "Secret_example", "AuthCode_example") // Oauth2AuthorizeAccessTokenPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Oauth2AuthorizeAccessTokenAPI.Oauth2AuthorizeAccessTokenPost(context.Background()).Oauth2AuthorizeAccessTokenPostRequest(oauth2AuthorizeAccessTokenPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2AuthorizeAccessTokenAPI.Oauth2AuthorizeAccessTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Oauth2AuthorizeAccessTokenPost`: Oauth2AuthorizeAccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `Oauth2AuthorizeAccessTokenAPI.Oauth2AuthorizeAccessTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AuthorizeAccessTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauth2AuthorizeAccessTokenPostRequest** | [**Oauth2AuthorizeAccessTokenPostRequest**](Oauth2AuthorizeAccessTokenPostRequest.md) |  | 

### Return type

[**Oauth2AuthorizeAccessTokenResponse**](Oauth2AuthorizeAccessTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

