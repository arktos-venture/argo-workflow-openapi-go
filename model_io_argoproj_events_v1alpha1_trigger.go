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

// IoArgoprojEventsV1alpha1Trigger struct for IoArgoprojEventsV1alpha1Trigger
type IoArgoprojEventsV1alpha1Trigger struct {
	Parameters *[]IoArgoprojEventsV1alpha1TriggerParameter `json:"parameters,omitempty"`
	Policy *IoArgoprojEventsV1alpha1TriggerPolicy `json:"policy,omitempty"`
	RetryStrategy *IoArgoprojEventsV1alpha1Backoff `json:"retryStrategy,omitempty"`
	Template *IoArgoprojEventsV1alpha1TriggerTemplate `json:"template,omitempty"`
}

// NewIoArgoprojEventsV1alpha1Trigger instantiates a new IoArgoprojEventsV1alpha1Trigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1Trigger() *IoArgoprojEventsV1alpha1Trigger {
	this := IoArgoprojEventsV1alpha1Trigger{}
	return &this
}

// NewIoArgoprojEventsV1alpha1TriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1Trigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1TriggerWithDefaults() *IoArgoprojEventsV1alpha1Trigger {
	this := IoArgoprojEventsV1alpha1Trigger{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Trigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter {
	if o == nil || o.Parameters == nil {
		var ret []IoArgoprojEventsV1alpha1TriggerParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Trigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Trigger) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []IoArgoprojEventsV1alpha1TriggerParameter and assigns it to the Parameters field.
func (o *IoArgoprojEventsV1alpha1Trigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter) {
	o.Parameters = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Trigger) GetPolicy() IoArgoprojEventsV1alpha1TriggerPolicy {
	if o == nil || o.Policy == nil {
		var ret IoArgoprojEventsV1alpha1TriggerPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Trigger) GetPolicyOk() (*IoArgoprojEventsV1alpha1TriggerPolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Trigger) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given IoArgoprojEventsV1alpha1TriggerPolicy and assigns it to the Policy field.
func (o *IoArgoprojEventsV1alpha1Trigger) SetPolicy(v IoArgoprojEventsV1alpha1TriggerPolicy) {
	o.Policy = &v
}

// GetRetryStrategy returns the RetryStrategy field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Trigger) GetRetryStrategy() IoArgoprojEventsV1alpha1Backoff {
	if o == nil || o.RetryStrategy == nil {
		var ret IoArgoprojEventsV1alpha1Backoff
		return ret
	}
	return *o.RetryStrategy
}

// GetRetryStrategyOk returns a tuple with the RetryStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Trigger) GetRetryStrategyOk() (*IoArgoprojEventsV1alpha1Backoff, bool) {
	if o == nil || o.RetryStrategy == nil {
		return nil, false
	}
	return o.RetryStrategy, true
}

// HasRetryStrategy returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Trigger) HasRetryStrategy() bool {
	if o != nil && o.RetryStrategy != nil {
		return true
	}

	return false
}

// SetRetryStrategy gets a reference to the given IoArgoprojEventsV1alpha1Backoff and assigns it to the RetryStrategy field.
func (o *IoArgoprojEventsV1alpha1Trigger) SetRetryStrategy(v IoArgoprojEventsV1alpha1Backoff) {
	o.RetryStrategy = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Trigger) GetTemplate() IoArgoprojEventsV1alpha1TriggerTemplate {
	if o == nil || o.Template == nil {
		var ret IoArgoprojEventsV1alpha1TriggerTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Trigger) GetTemplateOk() (*IoArgoprojEventsV1alpha1TriggerTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Trigger) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given IoArgoprojEventsV1alpha1TriggerTemplate and assigns it to the Template field.
func (o *IoArgoprojEventsV1alpha1Trigger) SetTemplate(v IoArgoprojEventsV1alpha1TriggerTemplate) {
	o.Template = &v
}

func (o IoArgoprojEventsV1alpha1Trigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	if o.RetryStrategy != nil {
		toSerialize["retryStrategy"] = o.RetryStrategy
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1Trigger struct {
	value *IoArgoprojEventsV1alpha1Trigger
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1Trigger) Get() *IoArgoprojEventsV1alpha1Trigger {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1Trigger) Set(val *IoArgoprojEventsV1alpha1Trigger) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1Trigger) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1Trigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1Trigger(val *IoArgoprojEventsV1alpha1Trigger) *NullableIoArgoprojEventsV1alpha1Trigger {
	return &NullableIoArgoprojEventsV1alpha1Trigger{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1Trigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1Trigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


