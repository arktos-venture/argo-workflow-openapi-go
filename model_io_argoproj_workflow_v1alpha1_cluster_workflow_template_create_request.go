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

// IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest struct for IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest
type IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest struct {
	CreateOptions *IoK8sApimachineryPkgApisMetaV1CreateOptions `json:"createOptions,omitempty"`
	Template *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate `json:"template,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest instantiates a new IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest() *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest {
	this := IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequestWithDefaults() *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest {
	this := IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest{}
	return &this
}

// GetCreateOptions returns the CreateOptions field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) GetCreateOptions() IoK8sApimachineryPkgApisMetaV1CreateOptions {
	if o == nil || o.CreateOptions == nil {
		var ret IoK8sApimachineryPkgApisMetaV1CreateOptions
		return ret
	}
	return *o.CreateOptions
}

// GetCreateOptionsOk returns a tuple with the CreateOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) GetCreateOptionsOk() (*IoK8sApimachineryPkgApisMetaV1CreateOptions, bool) {
	if o == nil || o.CreateOptions == nil {
		return nil, false
	}
	return o.CreateOptions, true
}

// HasCreateOptions returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) HasCreateOptions() bool {
	if o != nil && o.CreateOptions != nil {
		return true
	}

	return false
}

// SetCreateOptions gets a reference to the given IoK8sApimachineryPkgApisMetaV1CreateOptions and assigns it to the CreateOptions field.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) SetCreateOptions(v IoK8sApimachineryPkgApisMetaV1CreateOptions) {
	o.CreateOptions = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) GetTemplate() IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate {
	if o == nil || o.Template == nil {
		var ret IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) GetTemplateOk() (*IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate and assigns it to the Template field.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) SetTemplate(v IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) {
	o.Template = &v
}

func (o IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateOptions != nil {
		toSerialize["createOptions"] = o.CreateOptions
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest struct {
	value *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) Get() *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) Set(val *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest(val *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) *NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest {
	return &NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


