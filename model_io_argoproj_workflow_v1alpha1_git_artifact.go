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

// IoArgoprojWorkflowV1alpha1GitArtifact GitArtifact is the location of an git artifact
type IoArgoprojWorkflowV1alpha1GitArtifact struct {
	// Depth specifies clones/fetches should be shallow and include the given number of commits from the branch tip
	Depth *int32 `json:"depth,omitempty"`
	// DisableSubmodules disables submodules during git clone
	DisableSubmodules *bool `json:"disableSubmodules,omitempty"`
	// Fetch specifies a number of refs that should be fetched before checkout
	Fetch *[]string `json:"fetch,omitempty"`
	// InsecureIgnoreHostKey disables SSH strict host key checking during git clone
	InsecureIgnoreHostKey *bool `json:"insecureIgnoreHostKey,omitempty"`
	PasswordSecret *IoK8sApiCoreV1SecretKeySelector `json:"passwordSecret,omitempty"`
	// Repo is the git repository
	Repo string `json:"repo"`
	// Revision is the git commit, tag, branch to checkout
	Revision *string `json:"revision,omitempty"`
	SshPrivateKeySecret *IoK8sApiCoreV1SecretKeySelector `json:"sshPrivateKeySecret,omitempty"`
	UsernameSecret *IoK8sApiCoreV1SecretKeySelector `json:"usernameSecret,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1GitArtifact instantiates a new IoArgoprojWorkflowV1alpha1GitArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1GitArtifact(repo string) *IoArgoprojWorkflowV1alpha1GitArtifact {
	this := IoArgoprojWorkflowV1alpha1GitArtifact{}
	this.Repo = repo
	return &this
}

// NewIoArgoprojWorkflowV1alpha1GitArtifactWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1GitArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1GitArtifactWithDefaults() *IoArgoprojWorkflowV1alpha1GitArtifact {
	this := IoArgoprojWorkflowV1alpha1GitArtifact{}
	return &this
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetDepth() int32 {
	if o == nil || o.Depth == nil {
		var ret int32
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetDepthOk() (*int32, bool) {
	if o == nil || o.Depth == nil {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasDepth() bool {
	if o != nil && o.Depth != nil {
		return true
	}

	return false
}

// SetDepth gets a reference to the given int32 and assigns it to the Depth field.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetDepth(v int32) {
	o.Depth = &v
}

// GetDisableSubmodules returns the DisableSubmodules field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetDisableSubmodules() bool {
	if o == nil || o.DisableSubmodules == nil {
		var ret bool
		return ret
	}
	return *o.DisableSubmodules
}

// GetDisableSubmodulesOk returns a tuple with the DisableSubmodules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetDisableSubmodulesOk() (*bool, bool) {
	if o == nil || o.DisableSubmodules == nil {
		return nil, false
	}
	return o.DisableSubmodules, true
}

// HasDisableSubmodules returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasDisableSubmodules() bool {
	if o != nil && o.DisableSubmodules != nil {
		return true
	}

	return false
}

// SetDisableSubmodules gets a reference to the given bool and assigns it to the DisableSubmodules field.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetDisableSubmodules(v bool) {
	o.DisableSubmodules = &v
}

// GetFetch returns the Fetch field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetFetch() []string {
	if o == nil || o.Fetch == nil {
		var ret []string
		return ret
	}
	return *o.Fetch
}

// GetFetchOk returns a tuple with the Fetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetFetchOk() (*[]string, bool) {
	if o == nil || o.Fetch == nil {
		return nil, false
	}
	return o.Fetch, true
}

// HasFetch returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasFetch() bool {
	if o != nil && o.Fetch != nil {
		return true
	}

	return false
}

// SetFetch gets a reference to the given []string and assigns it to the Fetch field.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetFetch(v []string) {
	o.Fetch = &v
}

// GetInsecureIgnoreHostKey returns the InsecureIgnoreHostKey field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetInsecureIgnoreHostKey() bool {
	if o == nil || o.InsecureIgnoreHostKey == nil {
		var ret bool
		return ret
	}
	return *o.InsecureIgnoreHostKey
}

// GetInsecureIgnoreHostKeyOk returns a tuple with the InsecureIgnoreHostKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetInsecureIgnoreHostKeyOk() (*bool, bool) {
	if o == nil || o.InsecureIgnoreHostKey == nil {
		return nil, false
	}
	return o.InsecureIgnoreHostKey, true
}

// HasInsecureIgnoreHostKey returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasInsecureIgnoreHostKey() bool {
	if o != nil && o.InsecureIgnoreHostKey != nil {
		return true
	}

	return false
}

// SetInsecureIgnoreHostKey gets a reference to the given bool and assigns it to the InsecureIgnoreHostKey field.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetInsecureIgnoreHostKey(v bool) {
	o.InsecureIgnoreHostKey = &v
}

// GetPasswordSecret returns the PasswordSecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetPasswordSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.PasswordSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.PasswordSecret
}

// GetPasswordSecretOk returns a tuple with the PasswordSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetPasswordSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.PasswordSecret == nil {
		return nil, false
	}
	return o.PasswordSecret, true
}

// HasPasswordSecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasPasswordSecret() bool {
	if o != nil && o.PasswordSecret != nil {
		return true
	}

	return false
}

// SetPasswordSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the PasswordSecret field.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetPasswordSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.PasswordSecret = &v
}

// GetRepo returns the Repo field value
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetRepo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repo
}

// GetRepoOk returns a tuple with the Repo field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetRepoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Repo, true
}

// SetRepo sets field value
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetRepo(v string) {
	o.Repo = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetRevision(v string) {
	o.Revision = &v
}

// GetSshPrivateKeySecret returns the SshPrivateKeySecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetSshPrivateKeySecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.SshPrivateKeySecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.SshPrivateKeySecret
}

// GetSshPrivateKeySecretOk returns a tuple with the SshPrivateKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetSshPrivateKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.SshPrivateKeySecret == nil {
		return nil, false
	}
	return o.SshPrivateKeySecret, true
}

// HasSshPrivateKeySecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasSshPrivateKeySecret() bool {
	if o != nil && o.SshPrivateKeySecret != nil {
		return true
	}

	return false
}

// SetSshPrivateKeySecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the SshPrivateKeySecret field.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetSshPrivateKeySecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.SshPrivateKeySecret = &v
}

// GetUsernameSecret returns the UsernameSecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetUsernameSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.UsernameSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.UsernameSecret
}

// GetUsernameSecretOk returns a tuple with the UsernameSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetUsernameSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.UsernameSecret == nil {
		return nil, false
	}
	return o.UsernameSecret, true
}

// HasUsernameSecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasUsernameSecret() bool {
	if o != nil && o.UsernameSecret != nil {
		return true
	}

	return false
}

// SetUsernameSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the UsernameSecret field.
func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetUsernameSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.UsernameSecret = &v
}

func (o IoArgoprojWorkflowV1alpha1GitArtifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Depth != nil {
		toSerialize["depth"] = o.Depth
	}
	if o.DisableSubmodules != nil {
		toSerialize["disableSubmodules"] = o.DisableSubmodules
	}
	if o.Fetch != nil {
		toSerialize["fetch"] = o.Fetch
	}
	if o.InsecureIgnoreHostKey != nil {
		toSerialize["insecureIgnoreHostKey"] = o.InsecureIgnoreHostKey
	}
	if o.PasswordSecret != nil {
		toSerialize["passwordSecret"] = o.PasswordSecret
	}
	if true {
		toSerialize["repo"] = o.Repo
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.SshPrivateKeySecret != nil {
		toSerialize["sshPrivateKeySecret"] = o.SshPrivateKeySecret
	}
	if o.UsernameSecret != nil {
		toSerialize["usernameSecret"] = o.UsernameSecret
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1GitArtifact struct {
	value *IoArgoprojWorkflowV1alpha1GitArtifact
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1GitArtifact) Get() *IoArgoprojWorkflowV1alpha1GitArtifact {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1GitArtifact) Set(val *IoArgoprojWorkflowV1alpha1GitArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1GitArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1GitArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1GitArtifact(val *IoArgoprojWorkflowV1alpha1GitArtifact) *NullableIoArgoprojWorkflowV1alpha1GitArtifact {
	return &NullableIoArgoprojWorkflowV1alpha1GitArtifact{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1GitArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1GitArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


