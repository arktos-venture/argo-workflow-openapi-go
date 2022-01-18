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

// IoArgoprojEventsV1alpha1AMQPQueueBindConfig struct for IoArgoprojEventsV1alpha1AMQPQueueBindConfig
type IoArgoprojEventsV1alpha1AMQPQueueBindConfig struct {
	NoWait *bool `json:"noWait,omitempty"`
}

// NewIoArgoprojEventsV1alpha1AMQPQueueBindConfig instantiates a new IoArgoprojEventsV1alpha1AMQPQueueBindConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1AMQPQueueBindConfig() *IoArgoprojEventsV1alpha1AMQPQueueBindConfig {
	this := IoArgoprojEventsV1alpha1AMQPQueueBindConfig{}
	return &this
}

// NewIoArgoprojEventsV1alpha1AMQPQueueBindConfigWithDefaults instantiates a new IoArgoprojEventsV1alpha1AMQPQueueBindConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1AMQPQueueBindConfigWithDefaults() *IoArgoprojEventsV1alpha1AMQPQueueBindConfig {
	this := IoArgoprojEventsV1alpha1AMQPQueueBindConfig{}
	return &this
}

// GetNoWait returns the NoWait field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1AMQPQueueBindConfig) GetNoWait() bool {
	if o == nil || o.NoWait == nil {
		var ret bool
		return ret
	}
	return *o.NoWait
}

// GetNoWaitOk returns a tuple with the NoWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1AMQPQueueBindConfig) GetNoWaitOk() (*bool, bool) {
	if o == nil || o.NoWait == nil {
		return nil, false
	}
	return o.NoWait, true
}

// HasNoWait returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1AMQPQueueBindConfig) HasNoWait() bool {
	if o != nil && o.NoWait != nil {
		return true
	}

	return false
}

// SetNoWait gets a reference to the given bool and assigns it to the NoWait field.
func (o *IoArgoprojEventsV1alpha1AMQPQueueBindConfig) SetNoWait(v bool) {
	o.NoWait = &v
}

func (o IoArgoprojEventsV1alpha1AMQPQueueBindConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NoWait != nil {
		toSerialize["noWait"] = o.NoWait
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig struct {
	value *IoArgoprojEventsV1alpha1AMQPQueueBindConfig
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig) Get() *IoArgoprojEventsV1alpha1AMQPQueueBindConfig {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig) Set(val *IoArgoprojEventsV1alpha1AMQPQueueBindConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig(val *IoArgoprojEventsV1alpha1AMQPQueueBindConfig) *NullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig {
	return &NullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1AMQPQueueBindConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

