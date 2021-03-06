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

// IoArgoprojWorkflowV1alpha1PodGC PodGC describes how to delete completed pods as they complete
type IoArgoprojWorkflowV1alpha1PodGC struct {
	LabelSelector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"labelSelector,omitempty"`
	// Strategy is the strategy to use. One of \"OnPodCompletion\", \"OnPodSuccess\", \"OnWorkflowCompletion\", \"OnWorkflowSuccess\"
	Strategy *string `json:"strategy,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1PodGC instantiates a new IoArgoprojWorkflowV1alpha1PodGC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1PodGC() *IoArgoprojWorkflowV1alpha1PodGC {
	this := IoArgoprojWorkflowV1alpha1PodGC{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1PodGCWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1PodGC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1PodGCWithDefaults() *IoArgoprojWorkflowV1alpha1PodGC {
	this := IoArgoprojWorkflowV1alpha1PodGC{}
	return &this
}

// GetLabelSelector returns the LabelSelector field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1PodGC) GetLabelSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector {
	if o == nil || o.LabelSelector == nil {
		var ret IoK8sApimachineryPkgApisMetaV1LabelSelector
		return ret
	}
	return *o.LabelSelector
}

// GetLabelSelectorOk returns a tuple with the LabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1PodGC) GetLabelSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool) {
	if o == nil || o.LabelSelector == nil {
		return nil, false
	}
	return o.LabelSelector, true
}

// HasLabelSelector returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1PodGC) HasLabelSelector() bool {
	if o != nil && o.LabelSelector != nil {
		return true
	}

	return false
}

// SetLabelSelector gets a reference to the given IoK8sApimachineryPkgApisMetaV1LabelSelector and assigns it to the LabelSelector field.
func (o *IoArgoprojWorkflowV1alpha1PodGC) SetLabelSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector) {
	o.LabelSelector = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1PodGC) GetStrategy() string {
	if o == nil || o.Strategy == nil {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1PodGC) GetStrategyOk() (*string, bool) {
	if o == nil || o.Strategy == nil {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1PodGC) HasStrategy() bool {
	if o != nil && o.Strategy != nil {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *IoArgoprojWorkflowV1alpha1PodGC) SetStrategy(v string) {
	o.Strategy = &v
}

func (o IoArgoprojWorkflowV1alpha1PodGC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LabelSelector != nil {
		toSerialize["labelSelector"] = o.LabelSelector
	}
	if o.Strategy != nil {
		toSerialize["strategy"] = o.Strategy
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1PodGC struct {
	value *IoArgoprojWorkflowV1alpha1PodGC
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1PodGC) Get() *IoArgoprojWorkflowV1alpha1PodGC {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1PodGC) Set(val *IoArgoprojWorkflowV1alpha1PodGC) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1PodGC) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1PodGC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1PodGC(val *IoArgoprojWorkflowV1alpha1PodGC) *NullableIoArgoprojWorkflowV1alpha1PodGC {
	return &NullableIoArgoprojWorkflowV1alpha1PodGC{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1PodGC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1PodGC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


