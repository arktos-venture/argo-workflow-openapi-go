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

// IoArgoprojWorkflowV1alpha1Backoff Backoff is a backoff strategy to use within retryStrategy
type IoArgoprojWorkflowV1alpha1Backoff struct {
	// Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. \"2m\", \"1h\")
	Duration *string `json:"duration,omitempty"`
	Factor *string `json:"factor,omitempty"`
	// MaxDuration is the maximum amount of time allowed for the backoff strategy
	MaxDuration *string `json:"maxDuration,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1Backoff instantiates a new IoArgoprojWorkflowV1alpha1Backoff object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1Backoff() *IoArgoprojWorkflowV1alpha1Backoff {
	this := IoArgoprojWorkflowV1alpha1Backoff{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1BackoffWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Backoff object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1BackoffWithDefaults() *IoArgoprojWorkflowV1alpha1Backoff {
	this := IoArgoprojWorkflowV1alpha1Backoff{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Backoff) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Backoff) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Backoff) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *IoArgoprojWorkflowV1alpha1Backoff) SetDuration(v string) {
	o.Duration = &v
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Backoff) GetFactor() string {
	if o == nil || o.Factor == nil {
		var ret string
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Backoff) GetFactorOk() (*string, bool) {
	if o == nil || o.Factor == nil {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Backoff) HasFactor() bool {
	if o != nil && o.Factor != nil {
		return true
	}

	return false
}

// SetFactor gets a reference to the given string and assigns it to the Factor field.
func (o *IoArgoprojWorkflowV1alpha1Backoff) SetFactor(v string) {
	o.Factor = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Backoff) GetMaxDuration() string {
	if o == nil || o.MaxDuration == nil {
		var ret string
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Backoff) GetMaxDurationOk() (*string, bool) {
	if o == nil || o.MaxDuration == nil {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Backoff) HasMaxDuration() bool {
	if o != nil && o.MaxDuration != nil {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given string and assigns it to the MaxDuration field.
func (o *IoArgoprojWorkflowV1alpha1Backoff) SetMaxDuration(v string) {
	o.MaxDuration = &v
}

func (o IoArgoprojWorkflowV1alpha1Backoff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Factor != nil {
		toSerialize["factor"] = o.Factor
	}
	if o.MaxDuration != nil {
		toSerialize["maxDuration"] = o.MaxDuration
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1Backoff struct {
	value *IoArgoprojWorkflowV1alpha1Backoff
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1Backoff) Get() *IoArgoprojWorkflowV1alpha1Backoff {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1Backoff) Set(val *IoArgoprojWorkflowV1alpha1Backoff) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1Backoff) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1Backoff) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1Backoff(val *IoArgoprojWorkflowV1alpha1Backoff) *NullableIoArgoprojWorkflowV1alpha1Backoff {
	return &NullableIoArgoprojWorkflowV1alpha1Backoff{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1Backoff) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1Backoff) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


