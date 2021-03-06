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

// IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec struct for IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec
type IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec struct {
	Event IoArgoprojWorkflowV1alpha1Event `json:"event"`
	Submit *IoArgoprojWorkflowV1alpha1Submit `json:"submit,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec(event IoArgoprojWorkflowV1alpha1Event) *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec {
	this := IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec{}
	this.Event = event
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpecWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpecWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec {
	this := IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec{}
	return &this
}

// GetEvent returns the Event field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) GetEvent() IoArgoprojWorkflowV1alpha1Event {
	if o == nil {
		var ret IoArgoprojWorkflowV1alpha1Event
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) GetEventOk() (*IoArgoprojWorkflowV1alpha1Event, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) SetEvent(v IoArgoprojWorkflowV1alpha1Event) {
	o.Event = v
}

// GetSubmit returns the Submit field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) GetSubmit() IoArgoprojWorkflowV1alpha1Submit {
	if o == nil || o.Submit == nil {
		var ret IoArgoprojWorkflowV1alpha1Submit
		return ret
	}
	return *o.Submit
}

// GetSubmitOk returns a tuple with the Submit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) GetSubmitOk() (*IoArgoprojWorkflowV1alpha1Submit, bool) {
	if o == nil || o.Submit == nil {
		return nil, false
	}
	return o.Submit, true
}

// HasSubmit returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) HasSubmit() bool {
	if o != nil && o.Submit != nil {
		return true
	}

	return false
}

// SetSubmit gets a reference to the given IoArgoprojWorkflowV1alpha1Submit and assigns it to the Submit field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) SetSubmit(v IoArgoprojWorkflowV1alpha1Submit) {
	o.Submit = &v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if o.Submit != nil {
		toSerialize["submit"] = o.Submit
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) Get() *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) Set(val *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec(val *IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


