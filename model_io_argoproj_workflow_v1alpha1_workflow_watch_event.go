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

// IoArgoprojWorkflowV1alpha1WorkflowWatchEvent struct for IoArgoprojWorkflowV1alpha1WorkflowWatchEvent
type IoArgoprojWorkflowV1alpha1WorkflowWatchEvent struct {
	Object *IoArgoprojWorkflowV1alpha1Workflow `json:"object,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowWatchEvent instantiates a new IoArgoprojWorkflowV1alpha1WorkflowWatchEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowWatchEvent() *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent {
	this := IoArgoprojWorkflowV1alpha1WorkflowWatchEvent{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowWatchEventWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowWatchEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowWatchEventWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent {
	this := IoArgoprojWorkflowV1alpha1WorkflowWatchEvent{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) GetObject() IoArgoprojWorkflowV1alpha1Workflow {
	if o == nil || o.Object == nil {
		var ret IoArgoprojWorkflowV1alpha1Workflow
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) GetObjectOk() (*IoArgoprojWorkflowV1alpha1Workflow, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given IoArgoprojWorkflowV1alpha1Workflow and assigns it to the Object field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) SetObject(v IoArgoprojWorkflowV1alpha1Workflow) {
	o.Object = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) SetType(v string) {
	o.Type = &v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent) Get() *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent) Set(val *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent(val *IoArgoprojWorkflowV1alpha1WorkflowWatchEvent) *NullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowWatchEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


