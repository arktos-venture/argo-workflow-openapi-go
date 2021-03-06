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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint
type GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint struct {
	Url *string `json:"url,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint() *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpointWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpointWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1AWSEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


