# IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**RepoURL** | Pointer to **string** | RepoURL is the url for artifactory repo. | [optional] 
**UsernameSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository

`func NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository() *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository`

NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository instantiates a new IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepositoryWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepositoryWithDefaults() *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository`

NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepositoryWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetPasswordSecret() IoK8sApiCoreV1SecretKeySelector`

GetPasswordSecret returns the PasswordSecret field if non-nil, zero value otherwise.

### GetPasswordSecretOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetPasswordSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetPasswordSecretOk returns a tuple with the PasswordSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) SetPasswordSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetPasswordSecret sets PasswordSecret field to given value.

### HasPasswordSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) HasPasswordSecret() bool`

HasPasswordSecret returns a boolean if a field has been set.

### GetRepoURL

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetRepoURL() string`

GetRepoURL returns the RepoURL field if non-nil, zero value otherwise.

### GetRepoURLOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetRepoURLOk() (*string, bool)`

GetRepoURLOk returns a tuple with the RepoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoURL

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) SetRepoURL(v string)`

SetRepoURL sets RepoURL field to given value.

### HasRepoURL

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) HasRepoURL() bool`

HasRepoURL returns a boolean if a field has been set.

### GetUsernameSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetUsernameSecret() IoK8sApiCoreV1SecretKeySelector`

GetUsernameSecret returns the UsernameSecret field if non-nil, zero value otherwise.

### GetUsernameSecretOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) GetUsernameSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetUsernameSecretOk returns a tuple with the UsernameSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) SetUsernameSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetUsernameSecret sets UsernameSecret field to given value.

### HasUsernameSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository) HasUsernameSecret() bool`

HasUsernameSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


