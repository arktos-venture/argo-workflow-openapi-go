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

// IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate ClusterWorkflowTemplate is the definition of a workflow template resource in cluster scope
type IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata"`
	Spec IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec `json:"spec"`
}

// NewIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate instantiates a new IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate(metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta, spec IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate {
	this := IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate{}
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplateWithDefaults() *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate {
	this := IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta {
	if o == nil {
		var ret IoK8sApimachineryPkgApisMetaV1ObjectMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) GetSpec() IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec {
	if o == nil {
		var ret IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) GetSpecOk() (*IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) SetSpec(v IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) {
	o.Spec = v
}

func (o IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate struct {
	value *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) Get() *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) Set(val *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate(val *IoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) *NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate {
	return &NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1ClusterWorkflowTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


