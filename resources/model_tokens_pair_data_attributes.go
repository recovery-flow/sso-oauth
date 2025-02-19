/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TokensPairDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokensPairDataAttributes{}

// TokensPairDataAttributes struct for TokensPairDataAttributes
type TokensPairDataAttributes struct {
	// Access Token
	AccessToken string `json:"access_token"`
	// TokenRefresh Token
	RefreshToken string `json:"refresh_token"`
}

type _TokensPairDataAttributes TokensPairDataAttributes

// NewTokensPairDataAttributes instantiates a new TokensPairDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokensPairDataAttributes(accessToken string, refreshToken string) *TokensPairDataAttributes {
	this := TokensPairDataAttributes{}
	this.AccessToken = accessToken
	this.RefreshToken = refreshToken
	return &this
}

// NewTokensPairDataAttributesWithDefaults instantiates a new TokensPairDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokensPairDataAttributesWithDefaults() *TokensPairDataAttributes {
	this := TokensPairDataAttributes{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *TokensPairDataAttributes) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TokensPairDataAttributes) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TokensPairDataAttributes) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *TokensPairDataAttributes) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *TokensPairDataAttributes) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *TokensPairDataAttributes) SetRefreshToken(v string) {
	o.RefreshToken = v
}

func (o TokensPairDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokensPairDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["refresh_token"] = o.RefreshToken
	return toSerialize, nil
}

func (o *TokensPairDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
		"refresh_token",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTokensPairDataAttributes := _TokensPairDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokensPairDataAttributes)

	if err != nil {
		return err
	}

	*o = TokensPairDataAttributes(varTokensPairDataAttributes)

	return err
}

type NullableTokensPairDataAttributes struct {
	value *TokensPairDataAttributes
	isSet bool
}

func (v NullableTokensPairDataAttributes) Get() *TokensPairDataAttributes {
	return v.value
}

func (v *NullableTokensPairDataAttributes) Set(val *TokensPairDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTokensPairDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTokensPairDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokensPairDataAttributes(val *TokensPairDataAttributes) *NullableTokensPairDataAttributes {
	return &NullableTokensPairDataAttributes{value: val, isSet: true}
}

func (v NullableTokensPairDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokensPairDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


