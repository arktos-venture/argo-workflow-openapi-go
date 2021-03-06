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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus
type GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus struct {
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	LastScaledAt *time.Time `json:"lastScaledAt,omitempty"`
	Message *string `json:"message,omitempty"`
	Phase *string `json:"phase,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Replicas *int32 `json:"replicas,omitempty"`
	Selector *string `json:"selector,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus() *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatusWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatusWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus{}
	return &this
}

// GetLastScaledAt returns the LastScaledAt field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetLastScaledAt() time.Time {
	if o == nil || o.LastScaledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastScaledAt
}

// GetLastScaledAtOk returns a tuple with the LastScaledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetLastScaledAtOk() (*time.Time, bool) {
	if o == nil || o.LastScaledAt == nil {
		return nil, false
	}
	return o.LastScaledAt, true
}

// HasLastScaledAt returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasLastScaledAt() bool {
	if o != nil && o.LastScaledAt != nil {
		return true
	}

	return false
}

// SetLastScaledAt gets a reference to the given time.Time and assigns it to the LastScaledAt field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetLastScaledAt(v time.Time) {
	o.LastScaledAt = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetMessage(v string) {
	o.Message = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetPhase(v string) {
	o.Phase = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetReason(v string) {
	o.Reason = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetSelector(v string) {
	o.Selector = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastScaledAt != nil {
		toSerialize["lastScaledAt"] = o.LastScaledAt
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Selector != nil {
		toSerialize["selector"] = o.Selector
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


