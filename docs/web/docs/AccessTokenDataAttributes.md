# AccessTokenDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshToken** | Pointer to **string** | The access token to generate a new access token. | [optional] 

## Methods

### NewAccessTokenDataAttributes

`func NewAccessTokenDataAttributes() *AccessTokenDataAttributes`

NewAccessTokenDataAttributes instantiates a new AccessTokenDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenDataAttributesWithDefaults

`func NewAccessTokenDataAttributesWithDefaults() *AccessTokenDataAttributes`

NewAccessTokenDataAttributesWithDefaults instantiates a new AccessTokenDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshToken

`func (o *AccessTokenDataAttributes) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AccessTokenDataAttributes) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AccessTokenDataAttributes) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AccessTokenDataAttributes) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


