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

// IoArgoprojEventsV1alpha1EventSourceList struct for IoArgoprojEventsV1alpha1EventSourceList
type IoArgoprojEventsV1alpha1EventSourceList struct {
	Items *[]IoArgoprojEventsV1alpha1EventSource `json:"items,omitempty"`
	Metadata *IoK8sApimachineryPkgApisMetaV1ListMeta `json:"metadata,omitempty"`
}

// NewIoArgoprojEventsV1alpha1EventSourceList instantiates a new IoArgoprojEventsV1alpha1EventSourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1EventSourceList() *IoArgoprojEventsV1alpha1EventSourceList {
	this := IoArgoprojEventsV1alpha1EventSourceList{}
	return &this
}

// NewIoArgoprojEventsV1alpha1EventSourceListWithDefaults instantiates a new IoArgoprojEventsV1alpha1EventSourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1EventSourceListWithDefaults() *IoArgoprojEventsV1alpha1EventSourceList {
	this := IoArgoprojEventsV1alpha1EventSourceList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceList) GetItems() []IoArgoprojEventsV1alpha1EventSource {
	if o == nil || o.Items == nil {
		var ret []IoArgoprojEventsV1alpha1EventSource
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceList) GetItemsOk() (*[]IoArgoprojEventsV1alpha1EventSource, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []IoArgoprojEventsV1alpha1EventSource and assigns it to the Items field.
func (o *IoArgoprojEventsV1alpha1EventSourceList) SetItems(v []IoArgoprojEventsV1alpha1EventSource) {
	o.Items = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta {
	if o == nil || o.Metadata == nil {
		var ret IoK8sApimachineryPkgApisMetaV1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceList) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given IoK8sApimachineryPkgApisMetaV1ListMeta and assigns it to the Metadata field.
func (o *IoArgoprojEventsV1alpha1EventSourceList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta) {
	o.Metadata = &v
}

func (o IoArgoprojEventsV1alpha1EventSourceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1EventSourceList struct {
	value *IoArgoprojEventsV1alpha1EventSourceList
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1EventSourceList) Get() *IoArgoprojEventsV1alpha1EventSourceList {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1EventSourceList) Set(val *IoArgoprojEventsV1alpha1EventSourceList) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1EventSourceList) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1EventSourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1EventSourceList(val *IoArgoprojEventsV1alpha1EventSourceList) *NullableIoArgoprojEventsV1alpha1EventSourceList {
	return &NullableIoArgoprojEventsV1alpha1EventSourceList{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1EventSourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1EventSourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


