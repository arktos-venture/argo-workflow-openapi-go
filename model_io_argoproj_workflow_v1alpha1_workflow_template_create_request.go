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

// IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest struct for IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest
type IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest struct {
	CreateOptions *IoK8sApimachineryPkgApisMetaV1CreateOptions `json:"createOptions,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Template *IoArgoprojWorkflowV1alpha1WorkflowTemplate `json:"template,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest() *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequestWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest{}
	return &this
}

// GetCreateOptions returns the CreateOptions field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetCreateOptions() IoK8sApimachineryPkgApisMetaV1CreateOptions {
	if o == nil || o.CreateOptions == nil {
		var ret IoK8sApimachineryPkgApisMetaV1CreateOptions
		return ret
	}
	return *o.CreateOptions
}

// GetCreateOptionsOk returns a tuple with the CreateOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetCreateOptionsOk() (*IoK8sApimachineryPkgApisMetaV1CreateOptions, bool) {
	if o == nil || o.CreateOptions == nil {
		return nil, false
	}
	return o.CreateOptions, true
}

// HasCreateOptions returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) HasCreateOptions() bool {
	if o != nil && o.CreateOptions != nil {
		return true
	}

	return false
}

// SetCreateOptions gets a reference to the given IoK8sApimachineryPkgApisMetaV1CreateOptions and assigns it to the CreateOptions field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) SetCreateOptions(v IoK8sApimachineryPkgApisMetaV1CreateOptions) {
	o.CreateOptions = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetTemplate() IoArgoprojWorkflowV1alpha1WorkflowTemplate {
	if o == nil || o.Template == nil {
		var ret IoArgoprojWorkflowV1alpha1WorkflowTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetTemplateOk() (*IoArgoprojWorkflowV1alpha1WorkflowTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given IoArgoprojWorkflowV1alpha1WorkflowTemplate and assigns it to the Template field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) SetTemplate(v IoArgoprojWorkflowV1alpha1WorkflowTemplate) {
	o.Template = &v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateOptions != nil {
		toSerialize["createOptions"] = o.CreateOptions
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) Get() *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) Set(val *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest(val *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


