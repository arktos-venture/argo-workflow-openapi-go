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

// IoArgoprojEventsV1alpha1ArtifactLocation struct for IoArgoprojEventsV1alpha1ArtifactLocation
type IoArgoprojEventsV1alpha1ArtifactLocation struct {
	Configmap *IoK8sApiCoreV1ConfigMapKeySelector `json:"configmap,omitempty"`
	File *IoArgoprojEventsV1alpha1FileArtifact `json:"file,omitempty"`
	Git *IoArgoprojEventsV1alpha1GitArtifact `json:"git,omitempty"`
	Inline *string `json:"inline,omitempty"`
	Resource *IoArgoprojEventsV1alpha1Resource `json:"resource,omitempty"`
	S3 *IoArgoprojEventsV1alpha1S3Artifact `json:"s3,omitempty"`
	Url *IoArgoprojEventsV1alpha1URLArtifact `json:"url,omitempty"`
}

// NewIoArgoprojEventsV1alpha1ArtifactLocation instantiates a new IoArgoprojEventsV1alpha1ArtifactLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1ArtifactLocation() *IoArgoprojEventsV1alpha1ArtifactLocation {
	this := IoArgoprojEventsV1alpha1ArtifactLocation{}
	return &this
}

// NewIoArgoprojEventsV1alpha1ArtifactLocationWithDefaults instantiates a new IoArgoprojEventsV1alpha1ArtifactLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1ArtifactLocationWithDefaults() *IoArgoprojEventsV1alpha1ArtifactLocation {
	this := IoArgoprojEventsV1alpha1ArtifactLocation{}
	return &this
}

// GetConfigmap returns the Configmap field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetConfigmap() IoK8sApiCoreV1ConfigMapKeySelector {
	if o == nil || o.Configmap == nil {
		var ret IoK8sApiCoreV1ConfigMapKeySelector
		return ret
	}
	return *o.Configmap
}

// GetConfigmapOk returns a tuple with the Configmap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetConfigmapOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool) {
	if o == nil || o.Configmap == nil {
		return nil, false
	}
	return o.Configmap, true
}

// HasConfigmap returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasConfigmap() bool {
	if o != nil && o.Configmap != nil {
		return true
	}

	return false
}

// SetConfigmap gets a reference to the given IoK8sApiCoreV1ConfigMapKeySelector and assigns it to the Configmap field.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetConfigmap(v IoK8sApiCoreV1ConfigMapKeySelector) {
	o.Configmap = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetFile() IoArgoprojEventsV1alpha1FileArtifact {
	if o == nil || o.File == nil {
		var ret IoArgoprojEventsV1alpha1FileArtifact
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetFileOk() (*IoArgoprojEventsV1alpha1FileArtifact, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given IoArgoprojEventsV1alpha1FileArtifact and assigns it to the File field.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetFile(v IoArgoprojEventsV1alpha1FileArtifact) {
	o.File = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetGit() IoArgoprojEventsV1alpha1GitArtifact {
	if o == nil || o.Git == nil {
		var ret IoArgoprojEventsV1alpha1GitArtifact
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetGitOk() (*IoArgoprojEventsV1alpha1GitArtifact, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasGit() bool {
	if o != nil && o.Git != nil {
		return true
	}

	return false
}

// SetGit gets a reference to the given IoArgoprojEventsV1alpha1GitArtifact and assigns it to the Git field.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetGit(v IoArgoprojEventsV1alpha1GitArtifact) {
	o.Git = &v
}

// GetInline returns the Inline field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetInline() string {
	if o == nil || o.Inline == nil {
		var ret string
		return ret
	}
	return *o.Inline
}

// GetInlineOk returns a tuple with the Inline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetInlineOk() (*string, bool) {
	if o == nil || o.Inline == nil {
		return nil, false
	}
	return o.Inline, true
}

// HasInline returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasInline() bool {
	if o != nil && o.Inline != nil {
		return true
	}

	return false
}

// SetInline gets a reference to the given string and assigns it to the Inline field.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetInline(v string) {
	o.Inline = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetResource() IoArgoprojEventsV1alpha1Resource {
	if o == nil || o.Resource == nil {
		var ret IoArgoprojEventsV1alpha1Resource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetResourceOk() (*IoArgoprojEventsV1alpha1Resource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given IoArgoprojEventsV1alpha1Resource and assigns it to the Resource field.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetResource(v IoArgoprojEventsV1alpha1Resource) {
	o.Resource = &v
}

// GetS3 returns the S3 field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetS3() IoArgoprojEventsV1alpha1S3Artifact {
	if o == nil || o.S3 == nil {
		var ret IoArgoprojEventsV1alpha1S3Artifact
		return ret
	}
	return *o.S3
}

// GetS3Ok returns a tuple with the S3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetS3Ok() (*IoArgoprojEventsV1alpha1S3Artifact, bool) {
	if o == nil || o.S3 == nil {
		return nil, false
	}
	return o.S3, true
}

// HasS3 returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasS3() bool {
	if o != nil && o.S3 != nil {
		return true
	}

	return false
}

// SetS3 gets a reference to the given IoArgoprojEventsV1alpha1S3Artifact and assigns it to the S3 field.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetS3(v IoArgoprojEventsV1alpha1S3Artifact) {
	o.S3 = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetUrl() IoArgoprojEventsV1alpha1URLArtifact {
	if o == nil || o.Url == nil {
		var ret IoArgoprojEventsV1alpha1URLArtifact
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetUrlOk() (*IoArgoprojEventsV1alpha1URLArtifact, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given IoArgoprojEventsV1alpha1URLArtifact and assigns it to the Url field.
func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetUrl(v IoArgoprojEventsV1alpha1URLArtifact) {
	o.Url = &v
}

func (o IoArgoprojEventsV1alpha1ArtifactLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configmap != nil {
		toSerialize["configmap"] = o.Configmap
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	if o.Inline != nil {
		toSerialize["inline"] = o.Inline
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.S3 != nil {
		toSerialize["s3"] = o.S3
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1ArtifactLocation struct {
	value *IoArgoprojEventsV1alpha1ArtifactLocation
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1ArtifactLocation) Get() *IoArgoprojEventsV1alpha1ArtifactLocation {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1ArtifactLocation) Set(val *IoArgoprojEventsV1alpha1ArtifactLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1ArtifactLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1ArtifactLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1ArtifactLocation(val *IoArgoprojEventsV1alpha1ArtifactLocation) *NullableIoArgoprojEventsV1alpha1ArtifactLocation {
	return &NullableIoArgoprojEventsV1alpha1ArtifactLocation{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1ArtifactLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1ArtifactLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


