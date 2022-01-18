# IoArgoprojWorkflowV1alpha1S3ArtifactRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Bucket** | Pointer to **string** | Bucket is the name of the bucket | [optional] 
**CreateBucketIfNotPresent** | Pointer to [**IoArgoprojWorkflowV1alpha1CreateS3BucketOptions**](IoArgoprojWorkflowV1alpha1CreateS3BucketOptions.md) |  | [optional] 
**EncryptionOptions** | Pointer to [**IoArgoprojWorkflowV1alpha1S3EncryptionOptions**](IoArgoprojWorkflowV1alpha1S3EncryptionOptions.md) |  | [optional] 
**Endpoint** | Pointer to **string** | Endpoint is the hostname of the bucket endpoint | [optional] 
**Insecure** | Pointer to **bool** | Insecure will connect to the service with TLS | [optional] 
**KeyFormat** | Pointer to **string** | KeyFormat is defines the format of how to store keys. Can reference workflow variables | [optional] 
**KeyPrefix** | Pointer to **string** | KeyPrefix is prefix used as part of the bucket key in which the controller will store artifacts. DEPRECATED. Use KeyFormat instead | [optional] 
**Region** | Pointer to **string** | Region contains the optional bucket region | [optional] 
**RoleARN** | Pointer to **string** | RoleARN is the Amazon Resource Name (ARN) of the role to assume. | [optional] 
**SecretKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**UseSDKCreds** | Pointer to **bool** | UseSDKCreds tells the driver to figure out credentials based on sdk defaults. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1S3ArtifactRepository

`func NewIoArgoprojWorkflowV1alpha1S3ArtifactRepository() *IoArgoprojWorkflowV1alpha1S3ArtifactRepository`

NewIoArgoprojWorkflowV1alpha1S3ArtifactRepository instantiates a new IoArgoprojWorkflowV1alpha1S3ArtifactRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1S3ArtifactRepositoryWithDefaults

`func NewIoArgoprojWorkflowV1alpha1S3ArtifactRepositoryWithDefaults() *IoArgoprojWorkflowV1alpha1S3ArtifactRepository`

NewIoArgoprojWorkflowV1alpha1S3ArtifactRepositoryWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1S3ArtifactRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeySecret

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetAccessKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetAccessKeySecret returns the AccessKeySecret field if non-nil, zero value otherwise.

### GetAccessKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetAccessKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAccessKeySecretOk returns a tuple with the AccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecret

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetAccessKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetAccessKeySecret sets AccessKeySecret field to given value.

### HasAccessKeySecret

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasAccessKeySecret() bool`

HasAccessKeySecret returns a boolean if a field has been set.

### GetBucket

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCreateBucketIfNotPresent

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetCreateBucketIfNotPresent() IoArgoprojWorkflowV1alpha1CreateS3BucketOptions`

GetCreateBucketIfNotPresent returns the CreateBucketIfNotPresent field if non-nil, zero value otherwise.

### GetCreateBucketIfNotPresentOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetCreateBucketIfNotPresentOk() (*IoArgoprojWorkflowV1alpha1CreateS3BucketOptions, bool)`

GetCreateBucketIfNotPresentOk returns a tuple with the CreateBucketIfNotPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBucketIfNotPresent

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetCreateBucketIfNotPresent(v IoArgoprojWorkflowV1alpha1CreateS3BucketOptions)`

SetCreateBucketIfNotPresent sets CreateBucketIfNotPresent field to given value.

### HasCreateBucketIfNotPresent

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasCreateBucketIfNotPresent() bool`

HasCreateBucketIfNotPresent returns a boolean if a field has been set.

### GetEncryptionOptions

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetEncryptionOptions() IoArgoprojWorkflowV1alpha1S3EncryptionOptions`

GetEncryptionOptions returns the EncryptionOptions field if non-nil, zero value otherwise.

### GetEncryptionOptionsOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetEncryptionOptionsOk() (*IoArgoprojWorkflowV1alpha1S3EncryptionOptions, bool)`

GetEncryptionOptionsOk returns a tuple with the EncryptionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionOptions

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetEncryptionOptions(v IoArgoprojWorkflowV1alpha1S3EncryptionOptions)`

SetEncryptionOptions sets EncryptionOptions field to given value.

### HasEncryptionOptions

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasEncryptionOptions() bool`

HasEncryptionOptions returns a boolean if a field has been set.

### GetEndpoint

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetInsecure

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetKeyFormat

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetKeyFormat() string`

GetKeyFormat returns the KeyFormat field if non-nil, zero value otherwise.

### GetKeyFormatOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetKeyFormatOk() (*string, bool)`

GetKeyFormatOk returns a tuple with the KeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFormat

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetKeyFormat(v string)`

SetKeyFormat sets KeyFormat field to given value.

### HasKeyFormat

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasKeyFormat() bool`

HasKeyFormat returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.

### GetRegion

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleARN

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetRoleARN() string`

GetRoleARN returns the RoleARN field if non-nil, zero value otherwise.

### GetRoleARNOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetRoleARNOk() (*string, bool)`

GetRoleARNOk returns a tuple with the RoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleARN

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetRoleARN(v string)`

SetRoleARN sets RoleARN field to given value.

### HasRoleARN

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasRoleARN() bool`

HasRoleARN returns a boolean if a field has been set.

### GetSecretKeySecret

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetSecretKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetSecretKeySecret returns the SecretKeySecret field if non-nil, zero value otherwise.

### GetSecretKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetSecretKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSecretKeySecretOk returns a tuple with the SecretKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeySecret

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetSecretKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetSecretKeySecret sets SecretKeySecret field to given value.

### HasSecretKeySecret

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasSecretKeySecret() bool`

HasSecretKeySecret returns a boolean if a field has been set.

### GetUseSDKCreds

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetUseSDKCreds() bool`

GetUseSDKCreds returns the UseSDKCreds field if non-nil, zero value otherwise.

### GetUseSDKCredsOk

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetUseSDKCredsOk() (*bool, bool)`

GetUseSDKCredsOk returns a tuple with the UseSDKCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSDKCreds

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetUseSDKCreds(v bool)`

SetUseSDKCreds sets UseSDKCreds field to given value.

### HasUseSDKCreds

`func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasUseSDKCreds() bool`

HasUseSDKCreds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


