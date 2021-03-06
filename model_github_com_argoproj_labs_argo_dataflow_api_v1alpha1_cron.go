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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron
type GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron struct {
	Layout *string `json:"layout,omitempty"`
	Schedule *string `json:"schedule,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1CronWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1CronWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron{}
	return &this
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) GetLayout() string {
	if o == nil || o.Layout == nil {
		var ret string
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) GetLayoutOk() (*string, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) HasLayout() bool {
	if o != nil && o.Layout != nil {
		return true
	}

	return false
}

// SetLayout gets a reference to the given string and assigns it to the Layout field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) SetLayout(v string) {
	o.Layout = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) GetSchedule() string {
	if o == nil || o.Schedule == nil {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) GetScheduleOk() (*string, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) SetSchedule(v string) {
	o.Schedule = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Cron) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


