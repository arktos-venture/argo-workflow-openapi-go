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

// IoArgoprojEventsV1alpha1URLArtifact URLArtifact contains information about an artifact at an http endpoint.
type IoArgoprojEventsV1alpha1URLArtifact struct {
	Path *string `json:"path,omitempty"`
	VerifyCert *bool `json:"verifyCert,omitempty"`
}

// NewIoArgoprojEventsV1alpha1URLArtifact instantiates a new IoArgoprojEventsV1alpha1URLArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1URLArtifact() *IoArgoprojEventsV1alpha1URLArtifact {
	this := IoArgoprojEventsV1alpha1URLArtifact{}
	return &this
}

// NewIoArgoprojEventsV1alpha1URLArtifactWithDefaults instantiates a new IoArgoprojEventsV1alpha1URLArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1URLArtifactWithDefaults() *IoArgoprojEventsV1alpha1URLArtifact {
	this := IoArgoprojEventsV1alpha1URLArtifact{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1URLArtifact) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1URLArtifact) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1URLArtifact) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *IoArgoprojEventsV1alpha1URLArtifact) SetPath(v string) {
	o.Path = &v
}

// GetVerifyCert returns the VerifyCert field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1URLArtifact) GetVerifyCert() bool {
	if o == nil || o.VerifyCert == nil {
		var ret bool
		return ret
	}
	return *o.VerifyCert
}

// GetVerifyCertOk returns a tuple with the VerifyCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1URLArtifact) GetVerifyCertOk() (*bool, bool) {
	if o == nil || o.VerifyCert == nil {
		return nil, false
	}
	return o.VerifyCert, true
}

// HasVerifyCert returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1URLArtifact) HasVerifyCert() bool {
	if o != nil && o.VerifyCert != nil {
		return true
	}

	return false
}

// SetVerifyCert gets a reference to the given bool and assigns it to the VerifyCert field.
func (o *IoArgoprojEventsV1alpha1URLArtifact) SetVerifyCert(v bool) {
	o.VerifyCert = &v
}

func (o IoArgoprojEventsV1alpha1URLArtifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.VerifyCert != nil {
		toSerialize["verifyCert"] = o.VerifyCert
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1URLArtifact struct {
	value *IoArgoprojEventsV1alpha1URLArtifact
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1URLArtifact) Get() *IoArgoprojEventsV1alpha1URLArtifact {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1URLArtifact) Set(val *IoArgoprojEventsV1alpha1URLArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1URLArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1URLArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1URLArtifact(val *IoArgoprojEventsV1alpha1URLArtifact) *NullableIoArgoprojEventsV1alpha1URLArtifact {
	return &NullableIoArgoprojEventsV1alpha1URLArtifact{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1URLArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1URLArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


