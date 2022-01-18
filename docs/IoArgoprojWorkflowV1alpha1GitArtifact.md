# IoArgoprojWorkflowV1alpha1GitArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depth** | Pointer to **int32** | Depth specifies clones/fetches should be shallow and include the given number of commits from the branch tip | [optional] 
**DisableSubmodules** | Pointer to **bool** | DisableSubmodules disables submodules during git clone | [optional] 
**Fetch** | Pointer to **[]string** | Fetch specifies a number of refs that should be fetched before checkout | [optional] 
**InsecureIgnoreHostKey** | Pointer to **bool** | InsecureIgnoreHostKey disables SSH strict host key checking during git clone | [optional] 
**PasswordSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Repo** | **string** | Repo is the git repository | 
**Revision** | Pointer to **string** | Revision is the git commit, tag, branch to checkout | [optional] 
**SshPrivateKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**UsernameSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1GitArtifact

`func NewIoArgoprojWorkflowV1alpha1GitArtifact(repo string, ) *IoArgoprojWorkflowV1alpha1GitArtifact`

NewIoArgoprojWorkflowV1alpha1GitArtifact instantiates a new IoArgoprojWorkflowV1alpha1GitArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1GitArtifactWithDefaults

`func NewIoArgoprojWorkflowV1alpha1GitArtifactWithDefaults() *IoArgoprojWorkflowV1alpha1GitArtifact`

NewIoArgoprojWorkflowV1alpha1GitArtifactWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1GitArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepth

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetDisableSubmodules

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetDisableSubmodules() bool`

GetDisableSubmodules returns the DisableSubmodules field if non-nil, zero value otherwise.

### GetDisableSubmodulesOk

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetDisableSubmodulesOk() (*bool, bool)`

GetDisableSubmodulesOk returns a tuple with the DisableSubmodules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSubmodules

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetDisableSubmodules(v bool)`

SetDisableSubmodules sets DisableSubmodules field to given value.

### HasDisableSubmodules

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasDisableSubmodules() bool`

HasDisableSubmodules returns a boolean if a field has been set.

### GetFetch

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetFetch() []string`

GetFetch returns the Fetch field if non-nil, zero value otherwise.

### GetFetchOk

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetFetchOk() (*[]string, bool)`

GetFetchOk returns a tuple with the Fetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetch

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetFetch(v []string)`

SetFetch sets Fetch field to given value.

### HasFetch

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasFetch() bool`

HasFetch returns a boolean if a field has been set.

### GetInsecureIgnoreHostKey

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetInsecureIgnoreHostKey() bool`

GetInsecureIgnoreHostKey returns the InsecureIgnoreHostKey field if non-nil, zero value otherwise.

### GetInsecureIgnoreHostKeyOk

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetInsecureIgnoreHostKeyOk() (*bool, bool)`

GetInsecureIgnoreHostKeyOk returns a tuple with the InsecureIgnoreHostKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureIgnoreHostKey

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetInsecureIgnoreHostKey(v bool)`

SetInsecureIgnoreHostKey sets InsecureIgnoreHostKey field to given value.

### HasInsecureIgnoreHostKey

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasInsecureIgnoreHostKey() bool`

HasInsecureIgnoreHostKey returns a boolean if a field has been set.

### GetPasswordSecret

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetPasswordSecret() IoK8sApiCoreV1SecretKeySelector`

GetPasswordSecret returns the PasswordSecret field if non-nil, zero value otherwise.

### GetPasswordSecretOk

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetPasswordSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetPasswordSecretOk returns a tuple with the PasswordSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSecret

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetPasswordSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetPasswordSecret sets PasswordSecret field to given value.

### HasPasswordSecret

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasPasswordSecret() bool`

HasPasswordSecret returns a boolean if a field has been set.

### GetRepo

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetRepo(v string)`

SetRepo sets Repo field to given value.


### GetRevision

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSshPrivateKeySecret

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetSshPrivateKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetSshPrivateKeySecret returns the SshPrivateKeySecret field if non-nil, zero value otherwise.

### GetSshPrivateKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetSshPrivateKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSshPrivateKeySecretOk returns a tuple with the SshPrivateKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeySecret

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetSshPrivateKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetSshPrivateKeySecret sets SshPrivateKeySecret field to given value.

### HasSshPrivateKeySecret

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasSshPrivateKeySecret() bool`

HasSshPrivateKeySecret returns a boolean if a field has been set.

### GetUsernameSecret

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetUsernameSecret() IoK8sApiCoreV1SecretKeySelector`

GetUsernameSecret returns the UsernameSecret field if non-nil, zero value otherwise.

### GetUsernameSecretOk

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) GetUsernameSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetUsernameSecretOk returns a tuple with the UsernameSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameSecret

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) SetUsernameSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetUsernameSecret sets UsernameSecret field to given value.

### HasUsernameSecret

`func (o *IoArgoprojWorkflowV1alpha1GitArtifact) HasUsernameSecret() bool`

HasUsernameSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


