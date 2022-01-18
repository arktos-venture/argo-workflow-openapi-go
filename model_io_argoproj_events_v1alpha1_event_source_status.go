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

// IoArgoprojEventsV1alpha1EventSourceStatus struct for IoArgoprojEventsV1alpha1EventSourceStatus
type IoArgoprojEventsV1alpha1EventSourceStatus struct {
	Status *IoArgoprojEventsV1alpha1Status `json:"status,omitempty"`
}

// NewIoArgoprojEventsV1alpha1EventSourceStatus instantiates a new IoArgoprojEventsV1alpha1EventSourceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1EventSourceStatus() *IoArgoprojEventsV1alpha1EventSourceStatus {
	this := IoArgoprojEventsV1alpha1EventSourceStatus{}
	return &this
}

// NewIoArgoprojEventsV1alpha1EventSourceStatusWithDefaults instantiates a new IoArgoprojEventsV1alpha1EventSourceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1EventSourceStatusWithDefaults() *IoArgoprojEventsV1alpha1EventSourceStatus {
	this := IoArgoprojEventsV1alpha1EventSourceStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceStatus) GetStatus() IoArgoprojEventsV1alpha1Status {
	if o == nil || o.Status == nil {
		var ret IoArgoprojEventsV1alpha1Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceStatus) GetStatusOk() (*IoArgoprojEventsV1alpha1Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IoArgoprojEventsV1alpha1Status and assigns it to the Status field.
func (o *IoArgoprojEventsV1alpha1EventSourceStatus) SetStatus(v IoArgoprojEventsV1alpha1Status) {
	o.Status = &v
}

func (o IoArgoprojEventsV1alpha1EventSourceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1EventSourceStatus struct {
	value *IoArgoprojEventsV1alpha1EventSourceStatus
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1EventSourceStatus) Get() *IoArgoprojEventsV1alpha1EventSourceStatus {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1EventSourceStatus) Set(val *IoArgoprojEventsV1alpha1EventSourceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1EventSourceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1EventSourceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1EventSourceStatus(val *IoArgoprojEventsV1alpha1EventSourceStatus) *NullableIoArgoprojEventsV1alpha1EventSourceStatus {
	return &NullableIoArgoprojEventsV1alpha1EventSourceStatus{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1EventSourceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1EventSourceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


