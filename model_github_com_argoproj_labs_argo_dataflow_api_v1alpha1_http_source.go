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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource
type GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource struct {
	ServiceName *string `json:"serviceName,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource() *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSourceWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSourceWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource{}
	return &this
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) SetServiceName(v string) {
	o.ServiceName = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceName != nil {
		toSerialize["serviceName"] = o.ServiceName
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


