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

// IoArgoprojWorkflowV1alpha1WorkflowTemplateRef WorkflowTemplateRef is a reference to a WorkflowTemplate resource.
type IoArgoprojWorkflowV1alpha1WorkflowTemplateRef struct {
	// ClusterScope indicates the referred template is cluster scoped (i.e. a ClusterWorkflowTemplate).
	ClusterScope *bool `json:"clusterScope,omitempty"`
	// Name is the resource name of the workflow template.
	Name *string `json:"name,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowTemplateRef instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateRef() *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef {
	this := IoArgoprojWorkflowV1alpha1WorkflowTemplateRef{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowTemplateRefWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateRefWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef {
	this := IoArgoprojWorkflowV1alpha1WorkflowTemplateRef{}
	return &this
}

// GetClusterScope returns the ClusterScope field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) GetClusterScope() bool {
	if o == nil || o.ClusterScope == nil {
		var ret bool
		return ret
	}
	return *o.ClusterScope
}

// GetClusterScopeOk returns a tuple with the ClusterScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) GetClusterScopeOk() (*bool, bool) {
	if o == nil || o.ClusterScope == nil {
		return nil, false
	}
	return o.ClusterScope, true
}

// HasClusterScope returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) HasClusterScope() bool {
	if o != nil && o.ClusterScope != nil {
		return true
	}

	return false
}

// SetClusterScope gets a reference to the given bool and assigns it to the ClusterScope field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) SetClusterScope(v bool) {
	o.ClusterScope = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) SetName(v string) {
	o.Name = &v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClusterScope != nil {
		toSerialize["clusterScope"] = o.ClusterScope
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef) Get() *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef) Set(val *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef(val *IoArgoprojWorkflowV1alpha1WorkflowTemplateRef) *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


