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

// IoArgoprojWorkflowV1alpha1LabelValues Labels is list of workflow labels
type IoArgoprojWorkflowV1alpha1LabelValues struct {
	Items *[]string `json:"items,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1LabelValues instantiates a new IoArgoprojWorkflowV1alpha1LabelValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1LabelValues() *IoArgoprojWorkflowV1alpha1LabelValues {
	this := IoArgoprojWorkflowV1alpha1LabelValues{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1LabelValuesWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1LabelValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1LabelValuesWithDefaults() *IoArgoprojWorkflowV1alpha1LabelValues {
	this := IoArgoprojWorkflowV1alpha1LabelValues{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1LabelValues) GetItems() []string {
	if o == nil || o.Items == nil {
		var ret []string
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1LabelValues) GetItemsOk() (*[]string, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1LabelValues) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []string and assigns it to the Items field.
func (o *IoArgoprojWorkflowV1alpha1LabelValues) SetItems(v []string) {
	o.Items = &v
}

func (o IoArgoprojWorkflowV1alpha1LabelValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1LabelValues struct {
	value *IoArgoprojWorkflowV1alpha1LabelValues
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1LabelValues) Get() *IoArgoprojWorkflowV1alpha1LabelValues {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1LabelValues) Set(val *IoArgoprojWorkflowV1alpha1LabelValues) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1LabelValues) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1LabelValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1LabelValues(val *IoArgoprojWorkflowV1alpha1LabelValues) *NullableIoArgoprojWorkflowV1alpha1LabelValues {
	return &NullableIoArgoprojWorkflowV1alpha1LabelValues{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1LabelValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1LabelValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


