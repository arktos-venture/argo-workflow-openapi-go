# IoArgoprojEventsV1alpha1GitArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** |  | [optional] 
**CloneDirectory** | Pointer to **string** | Directory to clone the repository. We clone complete directory because GitArtifact is not limited to any specific Git service providers. Hence we don&#39;t use any specific git provider client. | [optional] 
**Creds** | Pointer to [**IoArgoprojEventsV1alpha1GitCreds**](IoArgoprojEventsV1alpha1GitCreds.md) |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 
**Remote** | Pointer to [**IoArgoprojEventsV1alpha1GitRemoteConfig**](IoArgoprojEventsV1alpha1GitRemoteConfig.md) |  | [optional] 
**SshKeyPath** | Pointer to **string** |  | [optional] 
**SshKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1GitArtifact

`func NewIoArgoprojEventsV1alpha1GitArtifact() *IoArgoprojEventsV1alpha1GitArtifact`

NewIoArgoprojEventsV1alpha1GitArtifact instantiates a new IoArgoprojEventsV1alpha1GitArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1GitArtifactWithDefaults

`func NewIoArgoprojEventsV1alpha1GitArtifactWithDefaults() *IoArgoprojEventsV1alpha1GitArtifact`

NewIoArgoprojEventsV1alpha1GitArtifactWithDefaults instantiates a new IoArgoprojEventsV1alpha1GitArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCloneDirectory

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetCloneDirectory() string`

GetCloneDirectory returns the CloneDirectory field if non-nil, zero value otherwise.

### GetCloneDirectoryOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetCloneDirectoryOk() (*string, bool)`

GetCloneDirectoryOk returns a tuple with the CloneDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneDirectory

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetCloneDirectory(v string)`

SetCloneDirectory sets CloneDirectory field to given value.

### HasCloneDirectory

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasCloneDirectory() bool`

HasCloneDirectory returns a boolean if a field has been set.

### GetCreds

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetCreds() IoArgoprojEventsV1alpha1GitCreds`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetCredsOk() (*IoArgoprojEventsV1alpha1GitCreds, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetCreds(v IoArgoprojEventsV1alpha1GitCreds)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### GetFilePath

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetRef

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRemote

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetRemote() IoArgoprojEventsV1alpha1GitRemoteConfig`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetRemoteOk() (*IoArgoprojEventsV1alpha1GitRemoteConfig, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetRemote(v IoArgoprojEventsV1alpha1GitRemoteConfig)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetSshKeyPath

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetSshKeyPath() string`

GetSshKeyPath returns the SshKeyPath field if non-nil, zero value otherwise.

### GetSshKeyPathOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetSshKeyPathOk() (*string, bool)`

GetSshKeyPathOk returns a tuple with the SshKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPath

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetSshKeyPath(v string)`

SetSshKeyPath sets SshKeyPath field to given value.

### HasSshKeyPath

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasSshKeyPath() bool`

HasSshKeyPath returns a boolean if a field has been set.

### GetSshKeySecret

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetSshKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetSshKeySecret returns the SshKeySecret field if non-nil, zero value otherwise.

### GetSshKeySecretOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetSshKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSshKeySecretOk returns a tuple with the SshKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeySecret

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetSshKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetSshKeySecret sets SshKeySecret field to given value.

### HasSshKeySecret

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasSshKeySecret() bool`

HasSshKeySecret returns a boolean if a field has been set.

### GetTag

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1GitArtifact) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1GitArtifact) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1GitArtifact) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


