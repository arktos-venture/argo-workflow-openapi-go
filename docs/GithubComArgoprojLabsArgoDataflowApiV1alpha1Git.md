# GithubComArgoprojLabsArgoDataflowApiV1alpha1Git

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to [**[]IoK8sApiCoreV1EnvVar**](IoK8sApiCoreV1EnvVar.md) |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**PasswordSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Path** | Pointer to **string** | +kubebuilder:default&#x3D;. | [optional] 
**SshPrivateKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UsernameSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Git

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Git() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Git instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Git object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1GitWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1GitWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1GitWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Git object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCommand

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnv

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetEnv() []IoK8sApiCoreV1EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetEnvOk() (*[]IoK8sApiCoreV1EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetEnv(v []IoK8sApiCoreV1EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetImage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetPasswordSecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetPasswordSecret() IoK8sApiCoreV1SecretKeySelector`

GetPasswordSecret returns the PasswordSecret field if non-nil, zero value otherwise.

### GetPasswordSecretOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetPasswordSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetPasswordSecretOk returns a tuple with the PasswordSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetPasswordSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetPasswordSecret sets PasswordSecret field to given value.

### HasPasswordSecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasPasswordSecret() bool`

HasPasswordSecret returns a boolean if a field has been set.

### GetPath

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSshPrivateKeySecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetSshPrivateKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetSshPrivateKeySecret returns the SshPrivateKeySecret field if non-nil, zero value otherwise.

### GetSshPrivateKeySecretOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetSshPrivateKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSshPrivateKeySecretOk returns a tuple with the SshPrivateKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeySecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetSshPrivateKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetSshPrivateKeySecret sets SshPrivateKeySecret field to given value.

### HasSshPrivateKeySecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasSshPrivateKeySecret() bool`

HasSshPrivateKeySecret returns a boolean if a field has been set.

### GetUrl

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsernameSecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetUsernameSecret() IoK8sApiCoreV1SecretKeySelector`

GetUsernameSecret returns the UsernameSecret field if non-nil, zero value otherwise.

### GetUsernameSecretOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) GetUsernameSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetUsernameSecretOk returns a tuple with the UsernameSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameSecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) SetUsernameSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetUsernameSecret sets UsernameSecret field to given value.

### HasUsernameSecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Git) HasUsernameSecret() bool`

HasUsernameSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


