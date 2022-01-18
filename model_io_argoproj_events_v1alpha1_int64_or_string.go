/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IoArgoprojEventsV1alpha1Int64OrString struct for IoArgoprojEventsV1alpha1Int64OrString
type IoArgoprojEventsV1alpha1Int64OrString struct {
	Int64Val *string `json:"int64Val,omitempty"`
	StrVal *string `json:"strVal,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewIoArgoprojEventsV1alpha1Int64OrString instantiates a new IoArgoprojEventsV1alpha1Int64OrString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1Int64OrString() *IoArgoprojEventsV1alpha1Int64OrString {
	this := IoArgoprojEventsV1alpha1Int64OrString{}
	return &this
}

// NewIoArgoprojEventsV1alpha1Int64OrStringWithDefaults instantiates a new IoArgoprojEventsV1alpha1Int64OrString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1Int64OrStringWithDefaults() *IoArgoprojEventsV1alpha1Int64OrString {
	this := IoArgoprojEventsV1alpha1Int64OrString{}
	return &this
}

// GetInt64Val returns the Int64Val field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Int64OrString) GetInt64Val() string {
	if o == nil || o.Int64Val == nil {
		var ret string
		return ret
	}
	return *o.Int64Val
}

// GetInt64ValOk returns a tuple with the Int64Val field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Int64OrString) GetInt64ValOk() (*string, bool) {
	if o == nil || o.Int64Val == nil {
		return nil, false
	}
	return o.Int64Val, true
}

// HasInt64Val returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Int64OrString) HasInt64Val() bool {
	if o != nil && o.Int64Val != nil {
		return true
	}

	return false
}

// SetInt64Val gets a reference to the given string and assigns it to the Int64Val field.
func (o *IoArgoprojEventsV1alpha1Int64OrString) SetInt64Val(v string) {
	o.Int64Val = &v
}

// GetStrVal returns the StrVal field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Int64OrString) GetStrVal() string {
	if o == nil || o.StrVal == nil {
		var ret string
		return ret
	}
	return *o.StrVal
}

// GetStrValOk returns a tuple with the StrVal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Int64OrString) GetStrValOk() (*string, bool) {
	if o == nil || o.StrVal == nil {
		return nil, false
	}
	return o.StrVal, true
}

// HasStrVal returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Int64OrString) HasStrVal() bool {
	if o != nil && o.StrVal != nil {
		return true
	}

	return false
}

// SetStrVal gets a reference to the given string and assigns it to the StrVal field.
func (o *IoArgoprojEventsV1alpha1Int64OrString) SetStrVal(v string) {
	o.StrVal = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Int64OrString) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Int64OrString) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Int64OrString) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IoArgoprojEventsV1alpha1Int64OrString) SetType(v string) {
	o.Type = &v
}

func (o IoArgoprojEventsV1alpha1Int64OrString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Int64Val != nil {
		toSerialize["int64Val"] = o.Int64Val
	}
	if o.StrVal != nil {
		toSerialize["strVal"] = o.StrVal
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1Int64OrString struct {
	value *IoArgoprojEventsV1alpha1Int64OrString
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1Int64OrString) Get() *IoArgoprojEventsV1alpha1Int64OrString {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1Int64OrString) Set(val *IoArgoprojEventsV1alpha1Int64OrString) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1Int64OrString) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1Int64OrString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1Int64OrString(val *IoArgoprojEventsV1alpha1Int64OrString) *NullableIoArgoprojEventsV1alpha1Int64OrString {
	return &NullableIoArgoprojEventsV1alpha1Int64OrString{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1Int64OrString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1Int64OrString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


