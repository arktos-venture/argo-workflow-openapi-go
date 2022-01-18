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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface
type GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface struct {
	Fifo *bool `json:"fifo,omitempty"`
	Http *map[string]interface{} `json:"http,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1InterfaceWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1InterfaceWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface{}
	return &this
}

// GetFifo returns the Fifo field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) GetFifo() bool {
	if o == nil || o.Fifo == nil {
		var ret bool
		return ret
	}
	return *o.Fifo
}

// GetFifoOk returns a tuple with the Fifo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) GetFifoOk() (*bool, bool) {
	if o == nil || o.Fifo == nil {
		return nil, false
	}
	return o.Fifo, true
}

// HasFifo returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) HasFifo() bool {
	if o != nil && o.Fifo != nil {
		return true
	}

	return false
}

// SetFifo gets a reference to the given bool and assigns it to the Fifo field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) SetFifo(v bool) {
	o.Fifo = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) GetHttp() map[string]interface{} {
	if o == nil || o.Http == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) GetHttpOk() (*map[string]interface{}, bool) {
	if o == nil || o.Http == nil {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) HasHttp() bool {
	if o != nil && o.Http != nil {
		return true
	}

	return false
}

// SetHttp gets a reference to the given map[string]interface{} and assigns it to the Http field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) SetHttp(v map[string]interface{}) {
	o.Http = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fifo != nil {
		toSerialize["fifo"] = o.Fifo
	}
	if o.Http != nil {
		toSerialize["http"] = o.Http
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Interface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

