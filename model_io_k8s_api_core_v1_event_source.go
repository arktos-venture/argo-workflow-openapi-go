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

// IoK8sApiCoreV1EventSource EventSource contains information for an event.
type IoK8sApiCoreV1EventSource struct {
	// Component from which the event is generated.
	Component *string `json:"component,omitempty"`
	// Node name on which the event is generated.
	Host *string `json:"host,omitempty"`
}

// NewIoK8sApiCoreV1EventSource instantiates a new IoK8sApiCoreV1EventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1EventSource() *IoK8sApiCoreV1EventSource {
	this := IoK8sApiCoreV1EventSource{}
	return &this
}

// NewIoK8sApiCoreV1EventSourceWithDefaults instantiates a new IoK8sApiCoreV1EventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1EventSourceWithDefaults() *IoK8sApiCoreV1EventSource {
	this := IoK8sApiCoreV1EventSource{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1EventSource) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1EventSource) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1EventSource) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *IoK8sApiCoreV1EventSource) SetComponent(v string) {
	o.Component = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1EventSource) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1EventSource) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1EventSource) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *IoK8sApiCoreV1EventSource) SetHost(v string) {
	o.Host = &v
}

func (o IoK8sApiCoreV1EventSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1EventSource struct {
	value *IoK8sApiCoreV1EventSource
	isSet bool
}

func (v NullableIoK8sApiCoreV1EventSource) Get() *IoK8sApiCoreV1EventSource {
	return v.value
}

func (v *NullableIoK8sApiCoreV1EventSource) Set(val *IoK8sApiCoreV1EventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1EventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1EventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1EventSource(val *IoK8sApiCoreV1EventSource) *NullableIoK8sApiCoreV1EventSource {
	return &NullableIoK8sApiCoreV1EventSource{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1EventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1EventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


