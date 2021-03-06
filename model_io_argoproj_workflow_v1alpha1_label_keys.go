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

// IoArgoprojWorkflowV1alpha1LabelKeys LabelKeys is list of keys
type IoArgoprojWorkflowV1alpha1LabelKeys struct {
	Items *[]string `json:"items,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1LabelKeys instantiates a new IoArgoprojWorkflowV1alpha1LabelKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1LabelKeys() *IoArgoprojWorkflowV1alpha1LabelKeys {
	this := IoArgoprojWorkflowV1alpha1LabelKeys{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1LabelKeysWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1LabelKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1LabelKeysWithDefaults() *IoArgoprojWorkflowV1alpha1LabelKeys {
	this := IoArgoprojWorkflowV1alpha1LabelKeys{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1LabelKeys) GetItems() []string {
	if o == nil || o.Items == nil {
		var ret []string
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1LabelKeys) GetItemsOk() (*[]string, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1LabelKeys) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []string and assigns it to the Items field.
func (o *IoArgoprojWorkflowV1alpha1LabelKeys) SetItems(v []string) {
	o.Items = &v
}

func (o IoArgoprojWorkflowV1alpha1LabelKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1LabelKeys struct {
	value *IoArgoprojWorkflowV1alpha1LabelKeys
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1LabelKeys) Get() *IoArgoprojWorkflowV1alpha1LabelKeys {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1LabelKeys) Set(val *IoArgoprojWorkflowV1alpha1LabelKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1LabelKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1LabelKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1LabelKeys(val *IoArgoprojWorkflowV1alpha1LabelKeys) *NullableIoArgoprojWorkflowV1alpha1LabelKeys {
	return &NullableIoArgoprojWorkflowV1alpha1LabelKeys{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1LabelKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1LabelKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


