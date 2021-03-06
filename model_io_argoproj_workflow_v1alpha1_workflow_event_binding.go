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

// IoArgoprojWorkflowV1alpha1WorkflowEventBinding WorkflowEventBinding is the definition of an event resource
type IoArgoprojWorkflowV1alpha1WorkflowEventBinding struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata"`
	Spec IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec `json:"spec"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowEventBinding instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowEventBinding(metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta, spec IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) *IoArgoprojWorkflowV1alpha1WorkflowEventBinding {
	this := IoArgoprojWorkflowV1alpha1WorkflowEventBinding{}
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowEventBinding {
	this := IoArgoprojWorkflowV1alpha1WorkflowEventBinding{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta {
	if o == nil {
		var ret IoK8sApimachineryPkgApisMetaV1ObjectMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetSpec() IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec {
	if o == nil {
		var ret IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetSpecOk() (*IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) SetSpec(v IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec) {
	o.Spec = v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowEventBinding) MarshalJSON() ([]byte, error) {
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

type NullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowEventBinding
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding) Get() *IoArgoprojWorkflowV1alpha1WorkflowEventBinding {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding) Set(val *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding(val *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


