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

// checks if the ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner{}

// ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner struct for ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner
type ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner struct {
	// Title is a short, human-readable summary of the problem
	Title string `json:"title"`
	// Detail is a human-readable explanation specific to this occurrence of the problem
	Detail *string `json:"detail,omitempty"`
	// Status is the HTTP status code applicable to this problem
	Status int32 `json:"status"`
}

type _ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner

// NewReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner instantiates a new ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner(title string, status int32) *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner {
	this := ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner{}
	this.Title = title
	this.Status = status
	return &this
}

// NewReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInnerWithDefaults instantiates a new ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInnerWithDefaults() *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner {
	this := ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner{}
	return &this
}

// GetTitle returns the Title field value
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) SetTitle(v string) {
	o.Title = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) SetDetail(v string) {
	o.Detail = &v
}

// GetStatus returns the Status field value
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) SetStatus(v int32) {
	o.Status = v
}

func (o ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"status",
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

	varReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner := _ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner)

	if err != nil {
		return err
	}

	*o = ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner(varReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner)

	return err
}

type NullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner struct {
	value *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner
	isSet bool
}

func (v NullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) Get() *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner {
	return v.value
}

func (v *NullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) Set(val *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner(val *ReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) *NullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner {
	return &NullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner{value: val, isSet: true}
}

func (v NullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReNewsSsoV1PublicGoogleLoginPost400ResponseErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


