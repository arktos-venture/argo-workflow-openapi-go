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

// IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest struct for IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest
type IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest struct {
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest() *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowTerminateRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowTerminateRequestWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) SetNamespace(v string) {
	o.Namespace = &v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) Get() *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) Set(val *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest(val *IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) *NullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


