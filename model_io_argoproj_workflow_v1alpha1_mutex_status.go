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

// IoArgoprojWorkflowV1alpha1MutexStatus MutexStatus contains which objects hold  mutex locks, and which objects this workflow is waiting on to release locks.
type IoArgoprojWorkflowV1alpha1MutexStatus struct {
	// Holding is a list of mutexes and their respective objects that are held by mutex lock for this io.argoproj.workflow.v1alpha1.
	Holding *[]IoArgoprojWorkflowV1alpha1MutexHolding `json:"holding,omitempty"`
	// Waiting is a list of mutexes and their respective objects this workflow is waiting for.
	Waiting *[]IoArgoprojWorkflowV1alpha1MutexHolding `json:"waiting,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1MutexStatus instantiates a new IoArgoprojWorkflowV1alpha1MutexStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1MutexStatus() *IoArgoprojWorkflowV1alpha1MutexStatus {
	this := IoArgoprojWorkflowV1alpha1MutexStatus{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1MutexStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1MutexStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1MutexStatusWithDefaults() *IoArgoprojWorkflowV1alpha1MutexStatus {
	this := IoArgoprojWorkflowV1alpha1MutexStatus{}
	return &this
}

// GetHolding returns the Holding field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1MutexStatus) GetHolding() []IoArgoprojWorkflowV1alpha1MutexHolding {
	if o == nil || o.Holding == nil {
		var ret []IoArgoprojWorkflowV1alpha1MutexHolding
		return ret
	}
	return *o.Holding
}

// GetHoldingOk returns a tuple with the Holding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1MutexStatus) GetHoldingOk() (*[]IoArgoprojWorkflowV1alpha1MutexHolding, bool) {
	if o == nil || o.Holding == nil {
		return nil, false
	}
	return o.Holding, true
}

// HasHolding returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1MutexStatus) HasHolding() bool {
	if o != nil && o.Holding != nil {
		return true
	}

	return false
}

// SetHolding gets a reference to the given []IoArgoprojWorkflowV1alpha1MutexHolding and assigns it to the Holding field.
func (o *IoArgoprojWorkflowV1alpha1MutexStatus) SetHolding(v []IoArgoprojWorkflowV1alpha1MutexHolding) {
	o.Holding = &v
}

// GetWaiting returns the Waiting field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1MutexStatus) GetWaiting() []IoArgoprojWorkflowV1alpha1MutexHolding {
	if o == nil || o.Waiting == nil {
		var ret []IoArgoprojWorkflowV1alpha1MutexHolding
		return ret
	}
	return *o.Waiting
}

// GetWaitingOk returns a tuple with the Waiting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1MutexStatus) GetWaitingOk() (*[]IoArgoprojWorkflowV1alpha1MutexHolding, bool) {
	if o == nil || o.Waiting == nil {
		return nil, false
	}
	return o.Waiting, true
}

// HasWaiting returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1MutexStatus) HasWaiting() bool {
	if o != nil && o.Waiting != nil {
		return true
	}

	return false
}

// SetWaiting gets a reference to the given []IoArgoprojWorkflowV1alpha1MutexHolding and assigns it to the Waiting field.
func (o *IoArgoprojWorkflowV1alpha1MutexStatus) SetWaiting(v []IoArgoprojWorkflowV1alpha1MutexHolding) {
	o.Waiting = &v
}

func (o IoArgoprojWorkflowV1alpha1MutexStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Holding != nil {
		toSerialize["holding"] = o.Holding
	}
	if o.Waiting != nil {
		toSerialize["waiting"] = o.Waiting
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1MutexStatus struct {
	value *IoArgoprojWorkflowV1alpha1MutexStatus
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1MutexStatus) Get() *IoArgoprojWorkflowV1alpha1MutexStatus {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1MutexStatus) Set(val *IoArgoprojWorkflowV1alpha1MutexStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1MutexStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1MutexStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1MutexStatus(val *IoArgoprojWorkflowV1alpha1MutexStatus) *NullableIoArgoprojWorkflowV1alpha1MutexStatus {
	return &NullableIoArgoprojWorkflowV1alpha1MutexStatus{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1MutexStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1MutexStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

