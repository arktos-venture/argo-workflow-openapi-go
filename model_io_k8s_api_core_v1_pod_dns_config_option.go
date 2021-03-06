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

// IoK8sApiCoreV1PodDNSConfigOption PodDNSConfigOption defines DNS resolver options of a pod.
type IoK8sApiCoreV1PodDNSConfigOption struct {
	// Required.
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewIoK8sApiCoreV1PodDNSConfigOption instantiates a new IoK8sApiCoreV1PodDNSConfigOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1PodDNSConfigOption() *IoK8sApiCoreV1PodDNSConfigOption {
	this := IoK8sApiCoreV1PodDNSConfigOption{}
	return &this
}

// NewIoK8sApiCoreV1PodDNSConfigOptionWithDefaults instantiates a new IoK8sApiCoreV1PodDNSConfigOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1PodDNSConfigOptionWithDefaults() *IoK8sApiCoreV1PodDNSConfigOption {
	this := IoK8sApiCoreV1PodDNSConfigOption{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PodDNSConfigOption) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PodDNSConfigOption) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PodDNSConfigOption) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IoK8sApiCoreV1PodDNSConfigOption) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PodDNSConfigOption) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PodDNSConfigOption) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PodDNSConfigOption) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IoK8sApiCoreV1PodDNSConfigOption) SetValue(v string) {
	o.Value = &v
}

func (o IoK8sApiCoreV1PodDNSConfigOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1PodDNSConfigOption struct {
	value *IoK8sApiCoreV1PodDNSConfigOption
	isSet bool
}

func (v NullableIoK8sApiCoreV1PodDNSConfigOption) Get() *IoK8sApiCoreV1PodDNSConfigOption {
	return v.value
}

func (v *NullableIoK8sApiCoreV1PodDNSConfigOption) Set(val *IoK8sApiCoreV1PodDNSConfigOption) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1PodDNSConfigOption) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1PodDNSConfigOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1PodDNSConfigOption(val *IoK8sApiCoreV1PodDNSConfigOption) *NullableIoK8sApiCoreV1PodDNSConfigOption {
	return &NullableIoK8sApiCoreV1PodDNSConfigOption{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1PodDNSConfigOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1PodDNSConfigOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


