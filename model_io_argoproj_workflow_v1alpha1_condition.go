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

// IoArgoprojWorkflowV1alpha1Condition struct for IoArgoprojWorkflowV1alpha1Condition
type IoArgoprojWorkflowV1alpha1Condition struct {
	// Message is the condition message
	Message *string `json:"message,omitempty"`
	// Status is the status of the condition
	Status *string `json:"status,omitempty"`
	// Type is the type of condition
	Type *string `json:"type,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1Condition instantiates a new IoArgoprojWorkflowV1alpha1Condition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1Condition() *IoArgoprojWorkflowV1alpha1Condition {
	this := IoArgoprojWorkflowV1alpha1Condition{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1ConditionWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Condition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1ConditionWithDefaults() *IoArgoprojWorkflowV1alpha1Condition {
	this := IoArgoprojWorkflowV1alpha1Condition{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Condition) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Condition) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Condition) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IoArgoprojWorkflowV1alpha1Condition) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Condition) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Condition) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Condition) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IoArgoprojWorkflowV1alpha1Condition) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Condition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Condition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Condition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IoArgoprojWorkflowV1alpha1Condition) SetType(v string) {
	o.Type = &v
}

func (o IoArgoprojWorkflowV1alpha1Condition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1Condition struct {
	value *IoArgoprojWorkflowV1alpha1Condition
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1Condition) Get() *IoArgoprojWorkflowV1alpha1Condition {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1Condition) Set(val *IoArgoprojWorkflowV1alpha1Condition) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1Condition) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1Condition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1Condition(val *IoArgoprojWorkflowV1alpha1Condition) *NullableIoArgoprojWorkflowV1alpha1Condition {
	return &NullableIoArgoprojWorkflowV1alpha1Condition{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1Condition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


