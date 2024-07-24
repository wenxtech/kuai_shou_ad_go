# \Oauth2AuthorizeRefreshTokenAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2AuthorizeRefreshTokenPost**](Oauth2AuthorizeRefreshTokenAPI.md#Oauth2AuthorizeRefreshTokenPost) | **Post** /oauth2/authorize/refresh_token | 刷新 token



## Oauth2AuthorizeRefreshTokenPost

> Oauth2AuthorizeAccessTokenResponse Oauth2AuthorizeRefreshTokenPost(ctx).Oauth2AuthorizeRefreshTokenPostRequest(oauth2AuthorizeRefreshTokenPostRequest).Execute()

刷新 token



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
	oauth2AuthorizeRefreshTokenPostRequest := *openapiclient.NewOauth2AuthorizeRefreshTokenPostRequest(int64(123), "Secret_example", "RefreshToken_example") // Oauth2AuthorizeRefreshTokenPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Oauth2AuthorizeRefreshTokenAPI.Oauth2AuthorizeRefreshTokenPost(context.Background()).Oauth2AuthorizeRefreshTokenPostRequest(oauth2AuthorizeRefreshTokenPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oauth2AuthorizeRefreshTokenAPI.Oauth2AuthorizeRefreshTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Oauth2AuthorizeRefreshTokenPost`: Oauth2AuthorizeAccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `Oauth2AuthorizeRefreshTokenAPI.Oauth2AuthorizeRefreshTokenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2AuthorizeRefreshTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oauth2AuthorizeRefreshTokenPostRequest** | [**Oauth2AuthorizeRefreshTokenPostRequest**](Oauth2AuthorizeRefreshTokenPostRequest.md) |  | 

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

