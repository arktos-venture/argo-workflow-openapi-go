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

// IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository ArtifactoryArtifactRepository defines the controller configuration for an artifactory artifact repository
type IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository struct {
	PasswordSecret *IoK8sApiCoreV1SecretKeySelector `json:"passwordSecret,omitempty"`
	// RepoURL is the url for artifactory repo.
	RepoURL *string `json:"repoURL,omitempty"`
	UsernameSecret *IoK8sApiCoreV1SecretKeySelector `json:"usernameSecret,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository instantiates a new IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository() *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository {
	this := IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepositoryWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepositoryWithDefaults() *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository {
	this := IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository{}
	return &this
}

// GetPasswordSecret returns the PasswordSecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetPasswordSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.PasswordSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.PasswordSecret
}

// GetPasswordSecretOk returns a tuple with the PasswordSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetPasswordSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.PasswordSecret == nil {
		return nil, false
	}
	return o.PasswordSecret, true
}

// HasPasswordSecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) HasPasswordSecret() bool {
	if o != nil && o.PasswordSecret != nil {
		return true
	}

	return false
}

// SetPasswordSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the PasswordSecret field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) SetPasswordSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.PasswordSecret = &v
}

// GetRepoURL returns the RepoURL field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetRepoURL() string {
	if o == nil || o.RepoURL == nil {
		var ret string
		return ret
	}
	return *o.RepoURL
}

// GetRepoURLOk returns a tuple with the RepoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetRepoURLOk() (*string, bool) {
	if o == nil || o.RepoURL == nil {
		return nil, false
	}
	return o.RepoURL, true
}

// HasRepoURL returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) HasRepoURL() bool {
	if o != nil && o.RepoURL != nil {
		return true
	}

	return false
}

// SetRepoURL gets a reference to the given string and assigns it to the RepoURL field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) SetRepoURL(v string) {
	o.RepoURL = &v
}

// GetUsernameSecret returns the UsernameSecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetUsernameSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.UsernameSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.UsernameSecret
}

// GetUsernameSecretOk returns a tuple with the UsernameSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetUsernameSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.UsernameSecret == nil {
		return nil, false
	}
	return o.UsernameSecret, true
}

// HasUsernameSecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) HasUsernameSecret() bool {
	if o != nil && o.UsernameSecret != nil {
		return true
	}

	return false
}

// SetUsernameSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the UsernameSecret field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) SetUsernameSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.UsernameSecret = &v
}

func (o IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PasswordSecret != nil {
		toSerialize["passwordSecret"] = o.PasswordSecret
	}
	if o.RepoURL != nil {
		toSerialize["repoURL"] = o.RepoURL
	}
	if o.UsernameSecret != nil {
		toSerialize["usernameSecret"] = o.UsernameSecret
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository struct {
	value *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) Get() *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) Set(val *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository(val *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) *NullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository {
	return &NullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


