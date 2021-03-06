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

// IoArgoprojWorkflowV1alpha1CronWorkflow CronWorkflow is the definition of a scheduled workflow resource
type IoArgoprojWorkflowV1alpha1CronWorkflow struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata"`
	Spec IoArgoprojWorkflowV1alpha1CronWorkflowSpec `json:"spec"`
	Status *IoArgoprojWorkflowV1alpha1CronWorkflowStatus `json:"status,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1CronWorkflow instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1CronWorkflow(metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta, spec IoArgoprojWorkflowV1alpha1CronWorkflowSpec) *IoArgoprojWorkflowV1alpha1CronWorkflow {
	this := IoArgoprojWorkflowV1alpha1CronWorkflow{}
	this.Metadata = metadata
	this.Spec = spec
	return &this
}

// NewIoArgoprojWorkflowV1alpha1CronWorkflowWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1CronWorkflowWithDefaults() *IoArgoprojWorkflowV1alpha1CronWorkflow {
	this := IoArgoprojWorkflowV1alpha1CronWorkflow{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta {
	if o == nil {
		var ret IoK8sApimachineryPkgApisMetaV1ObjectMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta) {
	o.Metadata = v
}

// GetSpec returns the Spec field value
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetSpec() IoArgoprojWorkflowV1alpha1CronWorkflowSpec {
	if o == nil {
		var ret IoArgoprojWorkflowV1alpha1CronWorkflowSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetSpecOk() (*IoArgoprojWorkflowV1alpha1CronWorkflowSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) SetSpec(v IoArgoprojWorkflowV1alpha1CronWorkflowSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetStatus() IoArgoprojWorkflowV1alpha1CronWorkflowStatus {
	if o == nil || o.Status == nil {
		var ret IoArgoprojWorkflowV1alpha1CronWorkflowStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) GetStatusOk() (*IoArgoprojWorkflowV1alpha1CronWorkflowStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IoArgoprojWorkflowV1alpha1CronWorkflowStatus and assigns it to the Status field.
func (o *IoArgoprojWorkflowV1alpha1CronWorkflow) SetStatus(v IoArgoprojWorkflowV1alpha1CronWorkflowStatus) {
	o.Status = &v
}

func (o IoArgoprojWorkflowV1alpha1CronWorkflow) MarshalJSON() ([]byte, error) {
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
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1CronWorkflow struct {
	value *IoArgoprojWorkflowV1alpha1CronWorkflow
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1CronWorkflow) Get() *IoArgoprojWorkflowV1alpha1CronWorkflow {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1CronWorkflow) Set(val *IoArgoprojWorkflowV1alpha1CronWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1CronWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1CronWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1CronWorkflow(val *IoArgoprojWorkflowV1alpha1CronWorkflow) *NullableIoArgoprojWorkflowV1alpha1CronWorkflow {
	return &NullableIoArgoprojWorkflowV1alpha1CronWorkflow{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1CronWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1CronWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


