# IoArgoprojWorkflowV1alpha1GCSArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Bucket is the name of the bucket | [optional] 
**Key** | **string** | Key is the path in the bucket where the artifact resides | 
**ServiceAccountKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1GCSArtifact

`func NewIoArgoprojWorkflowV1alpha1GCSArtifact(key string, ) *IoArgoprojWorkflowV1alpha1GCSArtifact`

NewIoArgoprojWorkflowV1alpha1GCSArtifact instantiates a new IoArgoprojWorkflowV1alpha1GCSArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1GCSArtifactWithDefaults

`func NewIoArgoprojWorkflowV1alpha1GCSArtifactWithDefaults() *IoArgoprojWorkflowV1alpha1GCSArtifact`

NewIoArgoprojWorkflowV1alpha1GCSArtifactWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1GCSArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetKey

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) SetKey(v string)`

SetKey sets Key field to given value.


### GetServiceAccountKeySecret

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) GetServiceAccountKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetServiceAccountKeySecret returns the ServiceAccountKeySecret field if non-nil, zero value otherwise.

### GetServiceAccountKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) GetServiceAccountKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetServiceAccountKeySecretOk returns a tuple with the ServiceAccountKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKeySecret

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) SetServiceAccountKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetServiceAccountKeySecret sets ServiceAccountKeySecret field to given value.

### HasServiceAccountKeySecret

`func (o *IoArgoprojWorkflowV1alpha1GCSArtifact) HasServiceAccountKeySecret() bool`

HasServiceAccountKeySecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


