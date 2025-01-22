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

// checks if the RegSimpleDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegSimpleDataAttributes{}

// RegSimpleDataAttributes struct for RegSimpleDataAttributes
type RegSimpleDataAttributes struct {
	// User email
	Email string `json:"email"`
}

type _RegSimpleDataAttributes RegSimpleDataAttributes

// NewRegSimpleDataAttributes instantiates a new RegSimpleDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegSimpleDataAttributes(email string) *RegSimpleDataAttributes {
	this := RegSimpleDataAttributes{}
	this.Email = email
	return &this
}

// NewRegSimpleDataAttributesWithDefaults instantiates a new RegSimpleDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegSimpleDataAttributesWithDefaults() *RegSimpleDataAttributes {
	this := RegSimpleDataAttributes{}
	return &this
}

// GetEmail returns the Email field value
func (o *RegSimpleDataAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RegSimpleDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RegSimpleDataAttributes) SetEmail(v string) {
	o.Email = v
}

func (o RegSimpleDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegSimpleDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *RegSimpleDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varRegSimpleDataAttributes := _RegSimpleDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegSimpleDataAttributes)

	if err != nil {
		return err
	}

	*o = RegSimpleDataAttributes(varRegSimpleDataAttributes)

	return err
}

type NullableRegSimpleDataAttributes struct {
	value *RegSimpleDataAttributes
	isSet bool
}

func (v NullableRegSimpleDataAttributes) Get() *RegSimpleDataAttributes {
	return v.value
}

func (v *NullableRegSimpleDataAttributes) Set(val *RegSimpleDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableRegSimpleDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableRegSimpleDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegSimpleDataAttributes(val *RegSimpleDataAttributes) *NullableRegSimpleDataAttributes {
	return &NullableRegSimpleDataAttributes{value: val, isSet: true}
}

func (v NullableRegSimpleDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegSimpleDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

