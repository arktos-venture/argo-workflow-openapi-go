# IoArgoprojWorkflowV1alpha1GCSArtifactRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the name of the bucket | [optional] 
**KeyFormat** | Pointer to **string** | KeyFormat is defines the format of how to store keys. Can reference workflow variables | [optional] 
**ServiceAccountKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1GCSArtifactRepository

`func NewIoArgoprojWorkflowV1alpha1GCSArtifactRepository() *IoArgoprojWorkflowV1alpha1GCSArtifactRepository`

NewIoArgoprojWorkflowV1alpha1GCSArtifactRepository instantiates a new IoArgoprojWorkflowV1alpha1GCSArtifactRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1GCSArtifactRepositoryWithDefaults

`func NewIoArgoprojWorkflowV1alpha1GCSArtifactRepositoryWithDefaults() *IoArgoprojWorkflowV1alpha1GCSArtifactRepository`

NewIoArgoprojWorkflowV1alpha1GCSArtifactRepositoryWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1GCSArtifactRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKeyFormat

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) GetKeyFormat() string`

GetKeyFormat returns the KeyFormat field if non-nil, zero value otherwise.

### GetKeyFormatOk

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) GetKeyFormatOk() (*string, bool)`

GetKeyFormatOk returns a tuple with the KeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFormat

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) SetKeyFormat(v string)`

SetKeyFormat sets KeyFormat field to given value.

### HasKeyFormat

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) HasKeyFormat() bool`

HasKeyFormat returns a boolean if a field has been set.

### GetServiceAccountKeySecret

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) GetServiceAccountKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetServiceAccountKeySecret returns the ServiceAccountKeySecret field if non-nil, zero value otherwise.

### GetServiceAccountKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) GetServiceAccountKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetServiceAccountKeySecretOk returns a tuple with the ServiceAccountKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKeySecret

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) SetServiceAccountKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetServiceAccountKeySecret sets ServiceAccountKeySecret field to given value.

### HasServiceAccountKeySecret

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifactRepository) HasServiceAccountKeySecret() bool`

HasServiceAccountKeySecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


