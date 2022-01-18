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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage
type GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage struct {
	Name *string `json:"name,omitempty"`
	SubPath *string `json:"subPath,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StorageWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StorageWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) SetName(v string) {
	o.Name = &v
}

// GetSubPath returns the SubPath field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) GetSubPath() string {
	if o == nil || o.SubPath == nil {
		var ret string
		return ret
	}
	return *o.SubPath
}

// GetSubPathOk returns a tuple with the SubPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) GetSubPathOk() (*string, bool) {
	if o == nil || o.SubPath == nil {
		return nil, false
	}
	return o.SubPath, true
}

// HasSubPath returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) HasSubPath() bool {
	if o != nil && o.SubPath != nil {
		return true
	}

	return false
}

// SetSubPath gets a reference to the given string and assigns it to the SubPath field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) SetSubPath(v string) {
	o.SubPath = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SubPath != nil {
		toSerialize["subPath"] = o.SubPath
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Storage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


