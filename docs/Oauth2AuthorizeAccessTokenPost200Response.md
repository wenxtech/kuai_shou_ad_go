# Oauth2AuthorizeAccessTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int64** |  | 
**Message** | **string** |  | 
**Data** | [**Oauth2AuthorizeAccessTokenResponseData**](Oauth2AuthorizeAccessTokenResponseData.md) |  | 

## Methods

### NewOauth2AuthorizeAccessTokenResponse

`func NewOauth2AuthorizeAccessTokenResponse(code int64, message string, data Oauth2AuthorizeAccessTokenResponseData, ) *Oauth2AuthorizeAccessTokenResponse`

NewOauth2AuthorizeAccessTokenResponse instantiates a new Oauth2AuthorizeAccessTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2AuthorizeAccessTokenResponseWithDefaults

`func NewOauth2AuthorizeAccessTokenResponseWithDefaults() *Oauth2AuthorizeAccessTokenResponse`

NewOauth2AuthorizeAccessTokenResponseWithDefaults instantiates a new Oauth2AuthorizeAccessTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Oauth2AuthorizeAccessTokenResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Oauth2AuthorizeAccessTokenResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Oauth2AuthorizeAccessTokenResponse) SetCode(v int64)`

SetCode sets Code field to given value.


### GetMessage

`func (o *Oauth2AuthorizeAccessTokenResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Oauth2AuthorizeAccessTokenResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Oauth2AuthorizeAccessTokenResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetData

`func (o *Oauth2AuthorizeAccessTokenResponse) GetData() Oauth2AuthorizeAccessTokenResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Oauth2AuthorizeAccessTokenResponse) GetDataOk() (*Oauth2AuthorizeAccessTokenResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Oauth2AuthorizeAccessTokenResponse) SetData(v Oauth2AuthorizeAccessTokenResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


