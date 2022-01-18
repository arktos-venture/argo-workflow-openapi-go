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

// IoK8sApimachineryPkgApisMetaV1LabelSelector A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
type IoK8sApimachineryPkgApisMetaV1LabelSelector struct {
	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	MatchExpressions *[]IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement `json:"matchExpressions,omitempty"`
	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.
	MatchLabels *map[string]string `json:"matchLabels,omitempty"`
}

// NewIoK8sApimachineryPkgApisMetaV1LabelSelector instantiates a new IoK8sApimachineryPkgApisMetaV1LabelSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApimachineryPkgApisMetaV1LabelSelector() *IoK8sApimachineryPkgApisMetaV1LabelSelector {
	this := IoK8sApimachineryPkgApisMetaV1LabelSelector{}
	return &this
}

// NewIoK8sApimachineryPkgApisMetaV1LabelSelectorWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1LabelSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApimachineryPkgApisMetaV1LabelSelectorWithDefaults() *IoK8sApimachineryPkgApisMetaV1LabelSelector {
	this := IoK8sApimachineryPkgApisMetaV1LabelSelector{}
	return &this
}

// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelector) GetMatchExpressions() []IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement {
	if o == nil || o.MatchExpressions == nil {
		var ret []IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement
		return ret
	}
	return *o.MatchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelector) GetMatchExpressionsOk() (*[]IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement, bool) {
	if o == nil || o.MatchExpressions == nil {
		return nil, false
	}
	return o.MatchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelector) HasMatchExpressions() bool {
	if o != nil && o.MatchExpressions != nil {
		return true
	}

	return false
}

// SetMatchExpressions gets a reference to the given []IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement and assigns it to the MatchExpressions field.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelector) SetMatchExpressions(v []IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) {
	o.MatchExpressions = &v
}

// GetMatchLabels returns the MatchLabels field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelector) GetMatchLabels() map[string]string {
	if o == nil || o.MatchLabels == nil {
		var ret map[string]string
		return ret
	}
	return *o.MatchLabels
}

// GetMatchLabelsOk returns a tuple with the MatchLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelector) GetMatchLabelsOk() (*map[string]string, bool) {
	if o == nil || o.MatchLabels == nil {
		return nil, false
	}
	return o.MatchLabels, true
}

// HasMatchLabels returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelector) HasMatchLabels() bool {
	if o != nil && o.MatchLabels != nil {
		return true
	}

	return false
}

// SetMatchLabels gets a reference to the given map[string]string and assigns it to the MatchLabels field.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelector) SetMatchLabels(v map[string]string) {
	o.MatchLabels = &v
}

func (o IoK8sApimachineryPkgApisMetaV1LabelSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MatchExpressions != nil {
		toSerialize["matchExpressions"] = o.MatchExpressions
	}
	if o.MatchLabels != nil {
		toSerialize["matchLabels"] = o.MatchLabels
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApimachineryPkgApisMetaV1LabelSelector struct {
	value *IoK8sApimachineryPkgApisMetaV1LabelSelector
	isSet bool
}

func (v NullableIoK8sApimachineryPkgApisMetaV1LabelSelector) Get() *IoK8sApimachineryPkgApisMetaV1LabelSelector {
	return v.value
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1LabelSelector) Set(val *IoK8sApimachineryPkgApisMetaV1LabelSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApimachineryPkgApisMetaV1LabelSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1LabelSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApimachineryPkgApisMetaV1LabelSelector(val *IoK8sApimachineryPkgApisMetaV1LabelSelector) *NullableIoK8sApimachineryPkgApisMetaV1LabelSelector {
	return &NullableIoK8sApimachineryPkgApisMetaV1LabelSelector{value: val, isSet: true}
}

func (v NullableIoK8sApimachineryPkgApisMetaV1LabelSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1LabelSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


