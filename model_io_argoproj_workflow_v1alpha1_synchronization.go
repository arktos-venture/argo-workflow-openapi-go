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

// IoArgoprojWorkflowV1alpha1Synchronization Synchronization holds synchronization lock configuration
type IoArgoprojWorkflowV1alpha1Synchronization struct {
	Mutex *IoArgoprojWorkflowV1alpha1Mutex `json:"mutex,omitempty"`
	Semaphore *IoArgoprojWorkflowV1alpha1SemaphoreRef `json:"semaphore,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1Synchronization instantiates a new IoArgoprojWorkflowV1alpha1Synchronization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1Synchronization() *IoArgoprojWorkflowV1alpha1Synchronization {
	this := IoArgoprojWorkflowV1alpha1Synchronization{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1SynchronizationWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Synchronization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1SynchronizationWithDefaults() *IoArgoprojWorkflowV1alpha1Synchronization {
	this := IoArgoprojWorkflowV1alpha1Synchronization{}
	return &this
}

// GetMutex returns the Mutex field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Synchronization) GetMutex() IoArgoprojWorkflowV1alpha1Mutex {
	if o == nil || o.Mutex == nil {
		var ret IoArgoprojWorkflowV1alpha1Mutex
		return ret
	}
	return *o.Mutex
}

// GetMutexOk returns a tuple with the Mutex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Synchronization) GetMutexOk() (*IoArgoprojWorkflowV1alpha1Mutex, bool) {
	if o == nil || o.Mutex == nil {
		return nil, false
	}
	return o.Mutex, true
}

// HasMutex returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Synchronization) HasMutex() bool {
	if o != nil && o.Mutex != nil {
		return true
	}

	return false
}

// SetMutex gets a reference to the given IoArgoprojWorkflowV1alpha1Mutex and assigns it to the Mutex field.
func (o *IoArgoprojWorkflowV1alpha1Synchronization) SetMutex(v IoArgoprojWorkflowV1alpha1Mutex) {
	o.Mutex = &v
}

// GetSemaphore returns the Semaphore field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Synchronization) GetSemaphore() IoArgoprojWorkflowV1alpha1SemaphoreRef {
	if o == nil || o.Semaphore == nil {
		var ret IoArgoprojWorkflowV1alpha1SemaphoreRef
		return ret
	}
	return *o.Semaphore
}

// GetSemaphoreOk returns a tuple with the Semaphore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Synchronization) GetSemaphoreOk() (*IoArgoprojWorkflowV1alpha1SemaphoreRef, bool) {
	if o == nil || o.Semaphore == nil {
		return nil, false
	}
	return o.Semaphore, true
}

// HasSemaphore returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Synchronization) HasSemaphore() bool {
	if o != nil && o.Semaphore != nil {
		return true
	}

	return false
}

// SetSemaphore gets a reference to the given IoArgoprojWorkflowV1alpha1SemaphoreRef and assigns it to the Semaphore field.
func (o *IoArgoprojWorkflowV1alpha1Synchronization) SetSemaphore(v IoArgoprojWorkflowV1alpha1SemaphoreRef) {
	o.Semaphore = &v
}

func (o IoArgoprojWorkflowV1alpha1Synchronization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mutex != nil {
		toSerialize["mutex"] = o.Mutex
	}
	if o.Semaphore != nil {
		toSerialize["semaphore"] = o.Semaphore
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1Synchronization struct {
	value *IoArgoprojWorkflowV1alpha1Synchronization
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1Synchronization) Get() *IoArgoprojWorkflowV1alpha1Synchronization {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1Synchronization) Set(val *IoArgoprojWorkflowV1alpha1Synchronization) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1Synchronization) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1Synchronization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1Synchronization(val *IoArgoprojWorkflowV1alpha1Synchronization) *NullableIoArgoprojWorkflowV1alpha1Synchronization {
	return &NullableIoArgoprojWorkflowV1alpha1Synchronization{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1Synchronization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1Synchronization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


