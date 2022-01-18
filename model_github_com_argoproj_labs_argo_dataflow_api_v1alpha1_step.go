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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1Step struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1Step
type GithubComArgoprojLabsArgoDataflowApiV1alpha1Step struct {
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`
	Spec *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec `json:"spec,omitempty"`
	Status *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus `json:"status,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Step instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Step object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Step() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Step{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Step object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Step{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret IoK8sApimachineryPkgApisMetaV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given IoK8sApimachineryPkgApisMetaV1ObjectMeta and assigns it to the Metadata field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) GetSpec() GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec {
	if o == nil || o.Spec == nil {
		var ret GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) GetSpecOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec and assigns it to the Spec field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) SetSpec(v GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) GetStatus() GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus {
	if o == nil || o.Status == nil {
		var ret GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) GetStatusOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus and assigns it to the Status field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) SetStatus(v GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) {
	o.Status = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Step) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Step) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


