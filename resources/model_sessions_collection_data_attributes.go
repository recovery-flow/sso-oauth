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

// checks if the SessionsCollectionDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionsCollectionDataAttributes{}

// SessionsCollectionDataAttributes struct for SessionsCollectionDataAttributes
type SessionsCollectionDataAttributes struct {
	Sessions []SessionData `json:"sessions"`
}

type _SessionsCollectionDataAttributes SessionsCollectionDataAttributes

// NewSessionsCollectionDataAttributes instantiates a new SessionsCollectionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionsCollectionDataAttributes(sessions []SessionData) *SessionsCollectionDataAttributes {
	this := SessionsCollectionDataAttributes{}
	this.Sessions = sessions
	return &this
}

// NewSessionsCollectionDataAttributesWithDefaults instantiates a new SessionsCollectionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionsCollectionDataAttributesWithDefaults() *SessionsCollectionDataAttributes {
	this := SessionsCollectionDataAttributes{}
	return &this
}

// GetSessions returns the Sessions field value
func (o *SessionsCollectionDataAttributes) GetSessions() []SessionData {
	if o == nil {
		var ret []SessionData
		return ret
	}

	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *SessionsCollectionDataAttributes) GetSessionsOk() ([]SessionData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sessions, true
}

// SetSessions sets field value
func (o *SessionsCollectionDataAttributes) SetSessions(v []SessionData) {
	o.Sessions = v
}

func (o SessionsCollectionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionsCollectionDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessions"] = o.Sessions
	return toSerialize, nil
}

func (o *SessionsCollectionDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sessions",
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

	varSessionsCollectionDataAttributes := _SessionsCollectionDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSessionsCollectionDataAttributes)

	if err != nil {
		return err
	}

	*o = SessionsCollectionDataAttributes(varSessionsCollectionDataAttributes)

	return err
}

type NullableSessionsCollectionDataAttributes struct {
	value *SessionsCollectionDataAttributes
	isSet bool
}

func (v NullableSessionsCollectionDataAttributes) Get() *SessionsCollectionDataAttributes {
	return v.value
}

func (v *NullableSessionsCollectionDataAttributes) Set(val *SessionsCollectionDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionsCollectionDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionsCollectionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionsCollectionDataAttributes(val *SessionsCollectionDataAttributes) *NullableSessionsCollectionDataAttributes {
	return &NullableSessionsCollectionDataAttributes{value: val, isSet: true}
}

func (v NullableSessionsCollectionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionsCollectionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


