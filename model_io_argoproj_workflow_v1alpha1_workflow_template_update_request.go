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

// IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest struct for IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest
type IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest struct {
	// DEPRECATED: This field is ignored.
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Template *IoArgoprojWorkflowV1alpha1WorkflowTemplate `json:"template,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest() *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequestWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) GetTemplate() IoArgoprojWorkflowV1alpha1WorkflowTemplate {
	if o == nil || o.Template == nil {
		var ret IoArgoprojWorkflowV1alpha1WorkflowTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) GetTemplateOk() (*IoArgoprojWorkflowV1alpha1WorkflowTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given IoArgoprojWorkflowV1alpha1WorkflowTemplate and assigns it to the Template field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) SetTemplate(v IoArgoprojWorkflowV1alpha1WorkflowTemplate) {
	o.Template = &v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) Get() *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) Set(val *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest(val *IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


