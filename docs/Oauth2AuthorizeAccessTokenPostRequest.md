# Oauth2AuthorizeAccessTokenPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **int64** | 申请应用后快手返回的 app_id | 
**Secret** | **string** | 申请应用后快手返回的 secret | 
**AuthCode** | **string** | 授权时返回的 auth_code | 

## Methods

### NewOauth2AuthorizeAccessTokenPostRequest

`func NewOauth2AuthorizeAccessTokenPostRequest(appId int64, secret string, authCode string, ) *Oauth2AuthorizeAccessTokenPostRequest`

NewOauth2AuthorizeAccessTokenPostRequest instantiates a new Oauth2AuthorizeAccessTokenPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2AuthorizeAccessTokenPostRequestWithDefaults

`func NewOauth2AuthorizeAccessTokenPostRequestWithDefaults() *Oauth2AuthorizeAccessTokenPostRequest`

NewOauth2AuthorizeAccessTokenPostRequestWithDefaults instantiates a new Oauth2AuthorizeAccessTokenPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *Oauth2AuthorizeAccessTokenPostRequest) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Oauth2AuthorizeAccessTokenPostRequest) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Oauth2AuthorizeAccessTokenPostRequest) SetAppId(v int64)`

SetAppId sets AppId field to given value.


### GetSecret

`func (o *Oauth2AuthorizeAccessTokenPostRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Oauth2AuthorizeAccessTokenPostRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Oauth2AuthorizeAccessTokenPostRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetAuthCode

`func (o *Oauth2AuthorizeAccessTokenPostRequest) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *Oauth2AuthorizeAccessTokenPostRequest) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *Oauth2AuthorizeAccessTokenPostRequest) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


