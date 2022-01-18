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

// IoArgoprojWorkflowV1alpha1MemoizationStatus MemoizationStatus is the status of this memoized node
type IoArgoprojWorkflowV1alpha1MemoizationStatus struct {
	// Cache is the name of the cache that was used
	CacheName string `json:"cacheName"`
	// Hit indicates whether this node was created from a cache entry
	Hit bool `json:"hit"`
	// Key is the name of the key used for this node's cache
	Key string `json:"key"`
}

// NewIoArgoprojWorkflowV1alpha1MemoizationStatus instantiates a new IoArgoprojWorkflowV1alpha1MemoizationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1MemoizationStatus(cacheName string, hit bool, key string) *IoArgoprojWorkflowV1alpha1MemoizationStatus {
	this := IoArgoprojWorkflowV1alpha1MemoizationStatus{}
	this.CacheName = cacheName
	this.Hit = hit
	this.Key = key
	return &this
}

// NewIoArgoprojWorkflowV1alpha1MemoizationStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1MemoizationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1MemoizationStatusWithDefaults() *IoArgoprojWorkflowV1alpha1MemoizationStatus {
	this := IoArgoprojWorkflowV1alpha1MemoizationStatus{}
	return &this
}

// GetCacheName returns the CacheName field value
func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetCacheName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CacheName
}

// GetCacheNameOk returns a tuple with the CacheName field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetCacheNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CacheName, true
}

// SetCacheName sets field value
func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) SetCacheName(v string) {
	o.CacheName = v
}

// GetHit returns the Hit field value
func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetHit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hit
}

// GetHitOk returns a tuple with the Hit field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetHitOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hit, true
}

// SetHit sets field value
func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) SetHit(v bool) {
	o.Hit = v
}

// GetKey returns the Key field value
func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) SetKey(v string) {
	o.Key = v
}

func (o IoArgoprojWorkflowV1alpha1MemoizationStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cacheName"] = o.CacheName
	}
	if true {
		toSerialize["hit"] = o.Hit
	}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1MemoizationStatus struct {
	value *IoArgoprojWorkflowV1alpha1MemoizationStatus
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1MemoizationStatus) Get() *IoArgoprojWorkflowV1alpha1MemoizationStatus {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1MemoizationStatus) Set(val *IoArgoprojWorkflowV1alpha1MemoizationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1MemoizationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1MemoizationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1MemoizationStatus(val *IoArgoprojWorkflowV1alpha1MemoizationStatus) *NullableIoArgoprojWorkflowV1alpha1MemoizationStatus {
	return &NullableIoArgoprojWorkflowV1alpha1MemoizationStatus{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1MemoizationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1MemoizationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


