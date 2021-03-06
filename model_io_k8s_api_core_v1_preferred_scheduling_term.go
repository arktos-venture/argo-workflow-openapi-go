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

// IoK8sApiCoreV1PreferredSchedulingTerm An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).
type IoK8sApiCoreV1PreferredSchedulingTerm struct {
	Preference IoK8sApiCoreV1NodeSelectorTerm `json:"preference"`
	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
	Weight int32 `json:"weight"`
}

// NewIoK8sApiCoreV1PreferredSchedulingTerm instantiates a new IoK8sApiCoreV1PreferredSchedulingTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1PreferredSchedulingTerm(preference IoK8sApiCoreV1NodeSelectorTerm, weight int32) *IoK8sApiCoreV1PreferredSchedulingTerm {
	this := IoK8sApiCoreV1PreferredSchedulingTerm{}
	this.Preference = preference
	this.Weight = weight
	return &this
}

// NewIoK8sApiCoreV1PreferredSchedulingTermWithDefaults instantiates a new IoK8sApiCoreV1PreferredSchedulingTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1PreferredSchedulingTermWithDefaults() *IoK8sApiCoreV1PreferredSchedulingTerm {
	this := IoK8sApiCoreV1PreferredSchedulingTerm{}
	return &this
}

// GetPreference returns the Preference field value
func (o *IoK8sApiCoreV1PreferredSchedulingTerm) GetPreference() IoK8sApiCoreV1NodeSelectorTerm {
	if o == nil {
		var ret IoK8sApiCoreV1NodeSelectorTerm
		return ret
	}

	return o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PreferredSchedulingTerm) GetPreferenceOk() (*IoK8sApiCoreV1NodeSelectorTerm, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Preference, true
}

// SetPreference sets field value
func (o *IoK8sApiCoreV1PreferredSchedulingTerm) SetPreference(v IoK8sApiCoreV1NodeSelectorTerm) {
	o.Preference = v
}

// GetWeight returns the Weight field value
func (o *IoK8sApiCoreV1PreferredSchedulingTerm) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PreferredSchedulingTerm) GetWeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *IoK8sApiCoreV1PreferredSchedulingTerm) SetWeight(v int32) {
	o.Weight = v
}

func (o IoK8sApiCoreV1PreferredSchedulingTerm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["preference"] = o.Preference
	}
	if true {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1PreferredSchedulingTerm struct {
	value *IoK8sApiCoreV1PreferredSchedulingTerm
	isSet bool
}

func (v NullableIoK8sApiCoreV1PreferredSchedulingTerm) Get() *IoK8sApiCoreV1PreferredSchedulingTerm {
	return v.value
}

func (v *NullableIoK8sApiCoreV1PreferredSchedulingTerm) Set(val *IoK8sApiCoreV1PreferredSchedulingTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1PreferredSchedulingTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1PreferredSchedulingTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1PreferredSchedulingTerm(val *IoK8sApiCoreV1PreferredSchedulingTerm) *NullableIoK8sApiCoreV1PreferredSchedulingTerm {
	return &NullableIoK8sApiCoreV1PreferredSchedulingTerm{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1PreferredSchedulingTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1PreferredSchedulingTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


