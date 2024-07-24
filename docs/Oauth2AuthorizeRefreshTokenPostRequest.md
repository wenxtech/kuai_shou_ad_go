# Oauth2AuthorizeRefreshTokenPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int64** | 申请应用后快手返回的 app_id | 
**Secret** | **string** | 申请应用后快手返回的 secret | 
**RefreshToken** | **string** | 最近一次快手返回的 refresh_token | 

## Methods

### NewOauth2AuthorizeRefreshTokenPostRequest

`func NewOauth2AuthorizeRefreshTokenPostRequest(appId int64, secret string, refreshToken string, ) *Oauth2AuthorizeRefreshTokenPostRequest`

NewOauth2AuthorizeRefreshTokenPostRequest instantiates a new Oauth2AuthorizeRefreshTokenPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2AuthorizeRefreshTokenPostRequestWithDefaults

`func NewOauth2AuthorizeRefreshTokenPostRequestWithDefaults() *Oauth2AuthorizeRefreshTokenPostRequest`

NewOauth2AuthorizeRefreshTokenPostRequestWithDefaults instantiates a new Oauth2AuthorizeRefreshTokenPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *Oauth2AuthorizeRefreshTokenPostRequest) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Oauth2AuthorizeRefreshTokenPostRequest) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Oauth2AuthorizeRefreshTokenPostRequest) SetAppId(v int64)`

SetAppId sets AppId field to given value.


### GetSecret

`func (o *Oauth2AuthorizeRefreshTokenPostRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Oauth2AuthorizeRefreshTokenPostRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Oauth2AuthorizeRefreshTokenPostRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetRefreshToken

`func (o *Oauth2AuthorizeRefreshTokenPostRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Oauth2AuthorizeRefreshTokenPostRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Oauth2AuthorizeRefreshTokenPostRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


