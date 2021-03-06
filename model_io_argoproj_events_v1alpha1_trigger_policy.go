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

// IoArgoprojEventsV1alpha1TriggerPolicy struct for IoArgoprojEventsV1alpha1TriggerPolicy
type IoArgoprojEventsV1alpha1TriggerPolicy struct {
	K8s *IoArgoprojEventsV1alpha1K8SResourcePolicy `json:"k8s,omitempty"`
	Status *IoArgoprojEventsV1alpha1StatusPolicy `json:"status,omitempty"`
}

// NewIoArgoprojEventsV1alpha1TriggerPolicy instantiates a new IoArgoprojEventsV1alpha1TriggerPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1TriggerPolicy() *IoArgoprojEventsV1alpha1TriggerPolicy {
	this := IoArgoprojEventsV1alpha1TriggerPolicy{}
	return &this
}

// NewIoArgoprojEventsV1alpha1TriggerPolicyWithDefaults instantiates a new IoArgoprojEventsV1alpha1TriggerPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1TriggerPolicyWithDefaults() *IoArgoprojEventsV1alpha1TriggerPolicy {
	this := IoArgoprojEventsV1alpha1TriggerPolicy{}
	return &this
}

// GetK8s returns the K8s field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1TriggerPolicy) GetK8s() IoArgoprojEventsV1alpha1K8SResourcePolicy {
	if o == nil || o.K8s == nil {
		var ret IoArgoprojEventsV1alpha1K8SResourcePolicy
		return ret
	}
	return *o.K8s
}

// GetK8sOk returns a tuple with the K8s field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1TriggerPolicy) GetK8sOk() (*IoArgoprojEventsV1alpha1K8SResourcePolicy, bool) {
	if o == nil || o.K8s == nil {
		return nil, false
	}
	return o.K8s, true
}

// HasK8s returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1TriggerPolicy) HasK8s() bool {
	if o != nil && o.K8s != nil {
		return true
	}

	return false
}

// SetK8s gets a reference to the given IoArgoprojEventsV1alpha1K8SResourcePolicy and assigns it to the K8s field.
func (o *IoArgoprojEventsV1alpha1TriggerPolicy) SetK8s(v IoArgoprojEventsV1alpha1K8SResourcePolicy) {
	o.K8s = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1TriggerPolicy) GetStatus() IoArgoprojEventsV1alpha1StatusPolicy {
	if o == nil || o.Status == nil {
		var ret IoArgoprojEventsV1alpha1StatusPolicy
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1TriggerPolicy) GetStatusOk() (*IoArgoprojEventsV1alpha1StatusPolicy, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1TriggerPolicy) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IoArgoprojEventsV1alpha1StatusPolicy and assigns it to the Status field.
func (o *IoArgoprojEventsV1alpha1TriggerPolicy) SetStatus(v IoArgoprojEventsV1alpha1StatusPolicy) {
	o.Status = &v
}

func (o IoArgoprojEventsV1alpha1TriggerPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.K8s != nil {
		toSerialize["k8s"] = o.K8s
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1TriggerPolicy struct {
	value *IoArgoprojEventsV1alpha1TriggerPolicy
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1TriggerPolicy) Get() *IoArgoprojEventsV1alpha1TriggerPolicy {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1TriggerPolicy) Set(val *IoArgoprojEventsV1alpha1TriggerPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1TriggerPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1TriggerPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1TriggerPolicy(val *IoArgoprojEventsV1alpha1TriggerPolicy) *NullableIoArgoprojEventsV1alpha1TriggerPolicy {
	return &NullableIoArgoprojEventsV1alpha1TriggerPolicy{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1TriggerPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1TriggerPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


