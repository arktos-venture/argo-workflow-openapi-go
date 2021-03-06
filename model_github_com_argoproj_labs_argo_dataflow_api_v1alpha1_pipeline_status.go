/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus
type GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus struct {
	Conditions *[]IoK8sApimachineryPkgApisMetaV1Condition `json:"conditions,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Message *string `json:"message,omitempty"`
	Phase *string `json:"phase,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus() *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatusWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatusWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetConditions() []IoK8sApimachineryPkgApisMetaV1Condition {
	if o == nil || o.Conditions == nil {
		var ret []IoK8sApimachineryPkgApisMetaV1Condition
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetConditionsOk() (*[]IoK8sApimachineryPkgApisMetaV1Condition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []IoK8sApimachineryPkgApisMetaV1Condition and assigns it to the Conditions field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) SetConditions(v []IoK8sApimachineryPkgApisMetaV1Condition) {
	o.Conditions = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) SetMessage(v string) {
	o.Message = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) SetPhase(v string) {
	o.Phase = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


