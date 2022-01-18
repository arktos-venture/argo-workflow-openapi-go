# IoArgoprojWorkflowV1alpha1ArtifactoryArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Url** | **string** | URL of the artifact | 
**UsernameSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifact

`func NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifact(url string, ) *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact`

NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifact instantiates a new IoArgoprojWorkflowV1alpha1ArtifactoryArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactWithDefaults() *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact`

NewIoArgoprojWorkflowV1alpha1ArtifactoryArtifactWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArtifactoryArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) GetPasswordSecret() IoK8sApiCoreV1SecretKeySelector`

GetPasswordSecret returns the PasswordSecret field if non-nil, zero value otherwise.

### GetPasswordSecretOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) GetPasswordSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetPasswordSecretOk returns a tuple with the PasswordSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) SetPasswordSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetPasswordSecret sets PasswordSecret field to given value.

### HasPasswordSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) HasPasswordSecret() bool`

HasPasswordSecret returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsernameSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) GetUsernameSecret() IoK8sApiCoreV1SecretKeySelector`

GetUsernameSecret returns the UsernameSecret field if non-nil, zero value otherwise.

### GetUsernameSecretOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) GetUsernameSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetUsernameSecretOk returns a tuple with the UsernameSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) SetUsernameSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetUsernameSecret sets UsernameSecret field to given value.

### HasUsernameSecret

`func (o *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) HasUsernameSecret() bool`

HasUsernameSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


