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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1Git struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1Git
type GithubComArgoprojLabsArgoDataflowApiV1alpha1Git struct {
	Branch *string `json:"branch,omitempty"`
	Command *[]string `json:"command,omitempty"`
	Env *[]IoK8sApiCoreV1EnvVar `json:"env,omitempty"`
	Image *string `json:"image,omitempty"`
	PasswordSecret *IoK8sApiCoreV1SecretKeySelector `json:"passwordSecret,omitempty"`
	// +kubebuilder:default=.
	Path *string `json:"path,omitempty"`
	SshPrivateKeySecret *IoK8sApiCoreV1SecretKeySelector `json:"sshPrivateKeySecret,omitempty"`
	Url *string `json:"url,omitempty"`
	UsernameSecret *IoK8sApiCoreV1SecretKeySelector `json:"usernameSecret,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Git instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Git object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Git() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Git{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1GitWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Git object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1GitWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Git{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetBranch(v string) {
	o.Branch = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetCommand() []string {
	if o == nil || o.Command == nil {
		var ret []string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetCommandOk() (*[]string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetCommand(v []string) {
	o.Command = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetEnv() []IoK8sApiCoreV1EnvVar {
	if o == nil || o.Env == nil {
		var ret []IoK8sApiCoreV1EnvVar
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetEnvOk() (*[]IoK8sApiCoreV1EnvVar, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []IoK8sApiCoreV1EnvVar and assigns it to the Env field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetEnv(v []IoK8sApiCoreV1EnvVar) {
	o.Env = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetImage(v string) {
	o.Image = &v
}

// GetPasswordSecret returns the PasswordSecret field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetPasswordSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.PasswordSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.PasswordSecret
}

// GetPasswordSecretOk returns a tuple with the PasswordSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetPasswordSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.PasswordSecret == nil {
		return nil, false
	}
	return o.PasswordSecret, true
}

// HasPasswordSecret returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasPasswordSecret() bool {
	if o != nil && o.PasswordSecret != nil {
		return true
	}

	return false
}

// SetPasswordSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the PasswordSecret field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetPasswordSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.PasswordSecret = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetPath(v string) {
	o.Path = &v
}

// GetSshPrivateKeySecret returns the SshPrivateKeySecret field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetSshPrivateKeySecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.SshPrivateKeySecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.SshPrivateKeySecret
}

// GetSshPrivateKeySecretOk returns a tuple with the SshPrivateKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetSshPrivateKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.SshPrivateKeySecret == nil {
		return nil, false
	}
	return o.SshPrivateKeySecret, true
}

// HasSshPrivateKeySecret returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasSshPrivateKeySecret() bool {
	if o != nil && o.SshPrivateKeySecret != nil {
		return true
	}

	return false
}

// SetSshPrivateKeySecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the SshPrivateKeySecret field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetSshPrivateKeySecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.SshPrivateKeySecret = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetUrl(v string) {
	o.Url = &v
}

// GetUsernameSecret returns the UsernameSecret field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetUsernameSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.UsernameSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.UsernameSecret
}

// GetUsernameSecretOk returns a tuple with the UsernameSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetUsernameSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.UsernameSecret == nil {
		return nil, false
	}
	return o.UsernameSecret, true
}

// HasUsernameSecret returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasUsernameSecret() bool {
	if o != nil && o.UsernameSecret != nil {
		return true
	}

	return false
}

// SetUsernameSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the UsernameSecret field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetUsernameSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.UsernameSecret = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.PasswordSecret != nil {
		toSerialize["passwordSecret"] = o.PasswordSecret
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.SshPrivateKeySecret != nil {
		toSerialize["sshPrivateKeySecret"] = o.SshPrivateKeySecret
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.UsernameSecret != nil {
		toSerialize["usernameSecret"] = o.UsernameSecret
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Git) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


